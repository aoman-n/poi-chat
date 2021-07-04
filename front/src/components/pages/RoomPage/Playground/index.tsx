import React, { useCallback } from 'react'
import { UserManager } from '@/utils/painter'
import Playground from './presenter'
import {
  RoomFragment,
  useRemoveBalloonMutation,
  useChangeBalloonPositionMutation,
  BalloonPosition,
} from '@/graphql'
import { useSubscribeRoomUserEvent, useSendMessage, useMovePos } from '@/hooks'
import { useCurrentUser } from '@/contexts/auth'

type PlaygroundContainerProps = {
  roomId: string
  room: RoomFragment['room']
  userManager: UserManager
  handleMoreMessage: () => void
  moreLoading: boolean
}

const useChangeBalloonPos = (userManager: UserManager, roomId: string) => {
  const { currentUser } = useCurrentUser()
  const [changeBalloonPos] = useChangeBalloonPositionMutation()

  const handleChangeBalloonPos = (balloonPosition: BalloonPosition) => {
    if (currentUser) {
      // TODO: 自身の情報はここで更新するようにする
      // TODO: globalUserとroomUserのidを同じにする
      // roomStatus/onlineStatusで管理する
      // 一旦はidを変換
      // const ids = currentUser.id.split(':')
      // userManager.chanageBalloonPos('RoomUser:' + ids[1], pos)

      changeBalloonPos({
        variables: {
          roomId,
          balloonPosition,
        },
      })
    }
  }

  return { handleChangeBalloonPos }
}

const useRemoveBalloon = (roomId: string) => {
  const [removeBalloon] = useRemoveBalloonMutation()

  const handleRemoveBalloon = useCallback(() => {
    removeBalloon({
      variables: { roomId },
    })
  }, [removeBalloon, roomId])

  return { handleRemoveBalloon }
}

const PlaygroundContainer: React.FC<PlaygroundContainerProps> = ({
  roomId,
  room,
  userManager,
  handleMoreMessage,
  moreLoading,
}) => {
  useSubscribeRoomUserEvent(roomId, userManager)
  const { handleMovePos } = useMovePos(roomId, userManager)
  const { handleSubmitMessage } = useSendMessage(roomId)
  const { handleChangeBalloonPos } = useChangeBalloonPos(userManager, roomId)
  const { handleRemoveBalloon } = useRemoveBalloon(roomId)

  return (
    <Playground
      messages={room.messages.nodes}
      hasMoreMessage={room.messages.pageInfo.hasPreviousPage}
      handleSubmitMessage={handleSubmitMessage}
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
        bgColor: room.bgColor,
        bgUrl: room.bgUrl,
      }}
      handleMoreMessage={handleMoreMessage}
      moreLoading={moreLoading}
      handleChangeBalloonPos={handleChangeBalloonPos}
      handleRemoveBalloon={handleRemoveBalloon}
    />
  )
}

export default PlaygroundContainer
