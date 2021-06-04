import React from 'react'
import { UserManager } from '@/utils/painter'
import Playground from './presentation'
import { RoomFragment } from '@/graphql'
import { useSubscribeRoomUserEvent, useSendMessage, useMove } from '@/hooks'

type PlaygroundContainerProps = {
  roomId: string
  room: RoomFragment['room']
  userManager: UserManager
  handleMoreMessage: () => void
}

const PlaygroundContainer: React.FC<PlaygroundContainerProps> = ({
  roomId,
  room,
  userManager,
  handleMoreMessage,
}) => {
  useSubscribeRoomUserEvent(roomId, userManager)
  const { handleMovePos } = useMove(roomId, userManager)
  const { handleSubmitMessage } = useSendMessage(roomId)

  return (
    <Playground
      messages={room.messages.nodes}
      handleSubmitMessage={handleSubmitMessage}
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
      }}
      handleMoreMessage={handleMoreMessage}
    />
  )
}

export default PlaygroundContainer
