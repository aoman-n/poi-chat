import React, { useEffect } from 'react'
import { RoomScreenPainter, UserManager } from '@/utils/painter'
import { ROOM_SCREEN_SIZE } from '@/constants'

const mainLoop = (
  ctx: CanvasRenderingContext2D,
  userManager: UserManager,
  roomScreenPainter: RoomScreenPainter,
) => {
  setInterval(() => {
    roomScreenPainter.draw(ctx)
    userManager.update()
    if (roomScreenPainter.isInitialized) {
      userManager.draw(ctx)
    }
  }, 1000 / 30)
}

export type RoomScreenProps = {
  userManager: UserManager
  handleMovePos: (x: number, y: number) => void
  bgColor: string
  bgUrl: string
}

const RoomScreen: React.FC<RoomScreenProps> = ({
  userManager,
  handleMovePos,
  bgColor,
  bgUrl,
}) => {
  /* eslint react-hooks/exhaustive-deps: 0 */
  useEffect(() => {
    // TODO: refを使う
    const canvas = document.getElementById('canvas') as HTMLCanvasElement
    const ctx = canvas.getContext('2d') as CanvasRenderingContext2D
    const roomScreenPainter = new RoomScreenPainter({ bgColor, bgUrl })
    mainLoop(ctx, userManager, roomScreenPainter)

    const clickHandler = (e: MouseEvent) => {
      const x = e.offsetX
      const y = e.offsetY
      handleMovePos(x, y)
    }

    // TODO: イベントをまびく
    canvas.addEventListener('click', clickHandler)

    return () => {
      canvas.removeEventListener('click', clickHandler)
    }
  }, [])

  return (
    <canvas
      id="canvas"
      width={ROOM_SCREEN_SIZE.WIDTH}
      height={ROOM_SCREEN_SIZE.HEIGHT}
    />
  )
}

export default RoomScreen
