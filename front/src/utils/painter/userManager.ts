import { User, BalloonPosition } from './user'
import { RoomFragment } from '@/graphql'

export class UserManager {
  #users: User[]

  constructor(roomUsers: RoomFragment['room']['users']) {
    this.#users = roomUsers.map(
      (u) =>
        new User({
          id: u.id,
          name: u.name,
          avatarUrl: u.avatarUrl,
          currentX: u.x,
          currentY: u.y,
          lastMessage: u.lastMessage?.body || '',
        }),
    )
  }

  update() {
    for (const user of this.#users) user.update()
  }

  draw(ctx: CanvasRenderingContext2D) {
    for (const user of this.#users) user.draw(ctx)
  }

  changePos(id: string, targetX: number, targetY: number) {
    const user = this.#users.find((u) => u.equalId(id))
    user?.changePos(targetX, targetY)
  }

  addUser(roomUser: RoomFragment['room']['users'][0]) {
    this.#users.push(
      new User({
        id: roomUser.id,
        name: roomUser.name,
        avatarUrl: roomUser.avatarUrl,
        currentX: roomUser.x,
        currentY: roomUser.y,
      }),
    )
  }

  deleteUser(userId: string) {
    this.#users = this.#users.filter((u) => !u.equalId(userId))
  }

  chanageBalloonPos(userId: string, pos: BalloonPosition) {
    const targetUser = this.findUserById(userId)
    targetUser?.changeBalloonPos(pos)
  }

  updateMessage(userId: string, message: string) {
    const targetUser = this.findUserById(userId)
    targetUser?.updateMessage(message)
  }

  findUserById(id: string) {
    return this.#users.find((u) => u.equalId(id))
  }
}
