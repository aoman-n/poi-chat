import { User } from './user'
import { BalloonPosition } from '@/constants'
import { RoomUserFieldsFragment } from '@/graphql'

export class UserManager {
  #users: User[]

  constructor(roomUsers: RoomUserFieldsFragment[]) {
    this.#users = roomUsers.map(
      (u) =>
        new User({
          id: u.id,
          name: u.user.name,
          avatarUrl: u.user.avatarUrl,
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

  addUser(roomUser: RoomUserFieldsFragment) {
    this.#users.push(
      new User({
        id: roomUser.id,
        name: roomUser.user.name,
        avatarUrl: roomUser.user.avatarUrl,
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
