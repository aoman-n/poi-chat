import React from 'react'
import { UserManager } from '@/utils/painter'
import { RoomFragment } from '@/graphql'
import { useSubscribeRoomUserEvent, useSendMessage, useMovePos } from '@/hooks'
import { useBalloon } from './hooks'
import Playground from './presenter'

type PlaygroundContainerProps = {
  roomId: string
  room: RoomFragment['room']
  userManager: UserManager
  handleMoreMessage: () => void
  moreLoading: boolean
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
  const {
    handleChangeBalloonPos,
    handleRemoveBalloon,
    handleBalloonStateToShowStatus,
    balloonState,
  } = useBalloon(userManager, roomId)
  const { handleSubmitMessage } = useSendMessage(
    roomId,
    handleBalloonStateToShowStatus,
  )

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
      balloonState={balloonState}
    />
  )
}

export default PlaygroundContainer
