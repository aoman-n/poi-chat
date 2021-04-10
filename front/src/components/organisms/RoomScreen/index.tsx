import React, { useEffect } from 'react'
import RoomScreenPainter from '@/painter/roomScreen'
import { UserManager } from '@/painter/user'
import { ROOM_SIZE } from '@/constants'

const mainLoop = (ctx: CanvasRenderingContext2D, userManager: UserManager) => {
  const roomScrenPainter = new RoomScreenPainter()

  setInterval(() => {
    roomScrenPainter.draw(ctx)
    userManager.update()
    userManager.draw(ctx)
  }, 1000 / 30)
}

export type RoomScreenProps = {
  userManager: UserManager
  handleMovePos?: (x: number, y: number) => void
}

const RoomScreen: React.FC<RoomScreenProps> = ({
  userManager,
  handleMovePos = (x: number, y: number) => {
    console.log({ x, y })
  },
}) => {
  /* eslint react-hooks/exhaustive-deps: 0 */
  useEffect(() => {
    const canvas = document.getElementById('canvas') as HTMLCanvasElement
    const ctx = canvas.getContext('2d') as CanvasRenderingContext2D
    mainLoop(ctx, userManager)

    const clickHandler = (e: MouseEvent) => {
      const x = e.offsetX
      const y = e.offsetY

      console.log({ x, y })
      handleMovePos(x, y)
    }

    // TODO: イベントをまびく
    canvas.addEventListener('click', clickHandler)

    return () => {
      canvas.removeEventListener('click', clickHandler)
    }
  }, [])

  return (
    <canvas id="canvas" width={ROOM_SIZE.WIDTH} height={ROOM_SIZE.HEIGHT} />
  )
}

export default RoomScreen
