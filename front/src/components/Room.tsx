import React, { useEffect } from 'react'
import RoomScreen from '../painter/roomScreen'
import { UserManager, User } from '../painter/user'
import { mockUsers } from '../mocks'

const SAMPLE_BG_IMAGE = 'https://pbs.twimg.com/media/EVUqmD3U4AABXgv.jpg'

const mainLoop = (ctx: CanvasRenderingContext2D, userManager: UserManager) => {
  const roomScreen = new RoomScreen({ bgImage: SAMPLE_BG_IMAGE })

  setInterval(() => {
    roomScreen.draw(ctx)
    userManager.update()
    userManager.draw(ctx)
  }, 1000 / 30)
}

const Room: React.FC = () => {
  useEffect(() => {
    // TODO: usersを取得
    const userManager = new UserManager(mockUsers.map((u) => new User(u)))
    // TODO: userの位置情報をSubscribe
    const canvas = document.getElementById('canvas') as HTMLCanvasElement
    const ctx = canvas.getContext('2d') as CanvasRenderingContext2D
    mainLoop(ctx, userManager)
  }, [])

  return (
    <div>
      <canvas id="canvas" width={RoomScreen.WIDTH} height={RoomScreen.HEIGHT} />
    </div>
  )
}

export default Room
