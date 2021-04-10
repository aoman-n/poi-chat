import React, { useEffect } from 'react'
import RoomScreen from '@/painter/roomScreen'
import { UserManager } from '@/painter/user'
import { ROOM_SIZE } from '@/constants'

const mainLoop = (ctx: CanvasRenderingContext2D, userManager: UserManager) => {
  const roomScreen = new RoomScreen()

  setInterval(() => {
    roomScreen.draw(ctx)
    userManager.update()
    userManager.draw(ctx)
  }, 1000 / 30)
}

export type RoomScreenComponentProps = {
  userManager: UserManager
}

const RoomScreenComponent: React.FC<RoomScreenComponentProps> = ({
  userManager,
}) => {
  useEffect(() => {
    const canvas = document.getElementById('canvas') as HTMLCanvasElement
    const ctx = canvas.getContext('2d') as CanvasRenderingContext2D
    mainLoop(ctx, userManager)
  }, [userManager])

  return (
    <canvas id="canvas" width={ROOM_SIZE.WIDTH} height={ROOM_SIZE.HEIGHT} />
  )
}

export default RoomScreenComponent
