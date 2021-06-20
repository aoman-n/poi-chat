import { produce } from 'immer'
import { useSnackbar } from 'notistack'
import { UserManager, BALLOON_POSITIONS } from '@/utils/painter'
import {
  RoomPageDocument,
  RoomPageQuery,
  RoomPageQueryVariables,
  useActedRoomUserEventSubscription,
  BalloonPosition as GraphBalloonPosition,
} from '@/graphql'

export const useSubscribeRoomUserEvent = (
  roomId: string,
  userManager: UserManager,
) => {
  const { enqueueSnackbar } = useSnackbar()

  useActedRoomUserEventSubscription({
    variables: { roomId },
    onSubscriptionData: ({ subscriptionData, client }) => {
      console.log({ roomUserEvent: subscriptionData })

      if (!subscriptionData.data || !subscriptionData.data.actedRoomUserEvent)
        return
      const { actedRoomUserEvent } = subscriptionData.data

      switch (actedRoomUserEvent.__typename) {
        case 'JoinedPayload': {
          const { roomUser } = actedRoomUserEvent
          enqueueSnackbar(`${roomUser.name} さんが入室しました`)
          userManager.addUser(roomUser)
          break
        }
        case 'ExitedPayload': {
          const exitedUser = userManager.findUserById(actedRoomUserEvent.userId)
          enqueueSnackbar(`${exitedUser?.name || 'XXX'} さんが退出しました`)
          userManager.deleteUser(actedRoomUserEvent.userId)
          break
        }
        case 'MovedPayload': {
          const { roomUser } = actedRoomUserEvent
          userManager.changePos(roomUser.id, roomUser.x, roomUser.y)
          break
        }
        case 'SentMassagePayload': {
          const {
            roomUser: { lastMessage, id },
          } = actedRoomUserEvent
          if (!lastMessage) return

          userManager.updateMessage(id, lastMessage.body)

          const pageQueryData = client.readQuery<
            RoomPageQuery,
            RoomPageQueryVariables
          >({
            query: RoomPageDocument,
            variables: {
              roomId,
            },
          })

          if (!pageQueryData) return

          const newPageQueryData = produce(pageQueryData, (draft) => {
            draft.room.messages.nodes.push(lastMessage)
          })

          client.writeQuery<RoomPageQuery>({
            query: RoomPageDocument,
            data: newPageQueryData,
          })
          break
        }
        case 'RemovedLastMessagePayload': {
          const { roomUser } = actedRoomUserEvent
          userManager.updateMessage(roomUser.id, '')
          break
        }
        case 'ChangedBalloonPositionPayload': {
          const { roomUser } = actedRoomUserEvent
          userManager.chanageBalloonPos(
            roomUser.id,
            convertBalloonPos(roomUser.balloonPosition),
          )
          break
        }
      }
    },
  })
}

const convertBalloonPos = (graphPosision: GraphBalloonPosition) => {
  switch (graphPosision) {
    case GraphBalloonPosition.TopRight:
      return BALLOON_POSITIONS.TOP_RIGHT
    case GraphBalloonPosition.TopLeft:
      return BALLOON_POSITIONS.TOP_LEFT
    case GraphBalloonPosition.BottomRight:
      return BALLOON_POSITIONS.BOTTOM_RIGHT
    case GraphBalloonPosition.BottomLeft:
      return BALLOON_POSITIONS.BOTTOM_LEFT
    default:
      return BALLOON_POSITIONS.TOP_RIGHT
  }
}
