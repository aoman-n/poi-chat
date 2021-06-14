import React from 'react'
import { UserManager, BalloonPosition } from '@/utils/painter'
import Playground from './presentation'
import { RoomFragment } from '@/graphql'
import { useSubscribeRoomUserEvent, useSendMessage, useMove } from '@/hooks'
import { useCurrentUser } from '@/contexts/auth'

type PlaygroundContainerProps = {
  roomId: string
  room: RoomFragment['room']
  userManager: UserManager
  handleMoreMessage: () => void
  moreLoading: boolean
}

const useBalloonPos = (userManager: UserManager) => {
  const { currentUser } = useCurrentUser()

  const handleChangeBalloonPos = (pos: BalloonPosition) => {
    if (currentUser) {
      console.log({ currentUser, pos })
      // TODO: globalUserとroomUserのidを同じにする
      // roomStatus/onlineStatusで管理する
      // 一旦はidを変換
      const ids = currentUser.id.split(':')

      userManager.chanageBalloonPos('RoomUser:' + ids[1], pos)
    }
  }

  return { handleChangeBalloonPos }
}

const PlaygroundContainer: React.FC<PlaygroundContainerProps> = ({
  roomId,
  room,
  userManager,
  handleMoreMessage,
  moreLoading,
}) => {
  useSubscribeRoomUserEvent(roomId, userManager)
  const { handleMovePos } = useMove(roomId, userManager)
  const { handleSubmitMessage } = useSendMessage(roomId)
  const { handleChangeBalloonPos } = useBalloonPos(userManager)

  return (
    <Playground
      messages={room.messages.nodes}
      hasMoreMessage={room.messages.pageInfo.hasPreviousPage}
      handleSubmitMessage={handleSubmitMessage}
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
      }}
      handleMoreMessage={handleMoreMessage}
      moreLoading={moreLoading}
      handleChangeBalloonPos={handleChangeBalloonPos}
    />
  )
}

export default PlaygroundContainer
