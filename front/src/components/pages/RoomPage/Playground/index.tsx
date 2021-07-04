import React, { useState, useCallback } from 'react'
import { UserManager } from '@/utils/painter'
import {
  RoomFragment,
  useRemoveBalloonMutation,
  useChangeBalloonPositionMutation,
} from '@/graphql'
import { useSubscribeRoomUserEvent, useSendMessage, useMovePos } from '@/hooks'
import { useCurrentUser } from '@/contexts/auth'
import { BalloonPosition, convertToGraphBalloonPos } from '@/constants'
import Playground, { BalloonState } from './presenter'

type PlaygroundContainerProps = {
  roomId: string
  room: RoomFragment['room']
  userManager: UserManager
  handleMoreMessage: () => void
  moreLoading: boolean
}

const initialBalloonState: BalloonState = {
  hasBalloon: false,
  position: null,
}

const defaultBalloonState: BalloonState = {
  hasBalloon: true,
  position: 'TOP_RIGHT',
}

const useBalloon = (userManager: UserManager, roomId: string) => {
  const { currentUser } = useCurrentUser()
  const [balloonState, setBalloonState] = useState<BalloonState>(
    initialBalloonState,
  )
  const [changeBalloonPos] = useChangeBalloonPositionMutation()
  const [removeBalloon] = useRemoveBalloonMutation()

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
          balloonPosition: convertToGraphBalloonPos(balloonPosition),
        },
      })
      setBalloonState((prev) => ({
        ...prev,
        position: balloonPosition,
      }))
    }
  }

  const handleRemoveBalloon = useCallback(() => {
    removeBalloon({
      variables: { roomId },
    })
    setBalloonState(initialBalloonState)
  }, [removeBalloon, roomId])

  const handleSetDefaultBalloonState = useCallback(() => {
    setBalloonState(defaultBalloonState)
  }, [])

  return {
    handleChangeBalloonPos,
    handleRemoveBalloon,
    handleSetDefaultBalloonState,
    balloonState,
  }
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
    handleSetDefaultBalloonState,
    balloonState,
  } = useBalloon(userManager, roomId)
  const { handleSubmitMessage } = useSendMessage(
    roomId,
    handleSetDefaultBalloonState,
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
