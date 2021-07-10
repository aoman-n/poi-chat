import { useState, useCallback } from 'react'
import { UserManager } from '@/utils/painter'
import {
  useRemoveBalloonMutation,
  useChangeBalloonPositionMutation,
} from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'
import { BalloonPosition, convertToGraphBalloonPos } from '@/constants'
import { BalloonState } from '@/components/pages/RoomPage/Settings'

const initialBalloonState: BalloonState = {
  hasBalloon: false,
  position: null,
}

export const useBalloon = (userManager: UserManager, roomId: string) => {
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
    setBalloonState((prev) => ({
      ...prev,
      hasBalloon: false,
    }))
  }, [removeBalloon, roomId])

  const handleChangeInitialBalloonState = useCallback(() => {
    setBalloonState({
      hasBalloon: true,
      position: 'TOP_RIGHT',
    })
  }, [])

  return {
    handleChangeBalloonPos,
    handleRemoveBalloon,
    handleChangeInitialBalloonState,
    balloonState,
  }
}
