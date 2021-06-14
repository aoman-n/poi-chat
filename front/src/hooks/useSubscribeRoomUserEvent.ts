import { produce } from 'immer'
import { UserManager } from '@/utils/painter'
import {
  RoomPageDocument,
  RoomPageQuery,
  RoomPageQueryVariables,
  useActedRoomUserEventSubscription,
} from '@/graphql'

export const useSubscribeRoomUserEvent = (
  roomId: string,
  userManager: UserManager,
) => {
  useActedRoomUserEventSubscription({
    variables: { roomId },
    onSubscriptionComplete: () => {
      console.log('start subscribe actedRoomUserEvent')
    },
    onSubscriptionData: ({ subscriptionData, client }) => {
      console.log({ roomUserEvent: subscriptionData })

      if (!subscriptionData.data) return

      const { actedRoomUserEvent } = subscriptionData.data

      switch (actedRoomUserEvent.__typename) {
        case 'JoinedPayload': {
          const { roomUser } = actedRoomUserEvent
          userManager.addUser(roomUser)
          break
        }
        case 'ExitedPayload':
          userManager.deleteUser(actedRoomUserEvent.userId)
          break
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
      }
    },
  })
}
