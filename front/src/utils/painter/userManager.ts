import { User } from './user'
import { RoomFragment } from '@/graphql'

export class UserManager {
  private _users: User[]

  constructor(roomUsers: RoomFragment['room']['users']) {
    this._users = roomUsers.map(
      (u) =>
        new User({
          id: u.id,
          avatarUrl: u.avatarUrl,
          currentX: u.x,
          currentY: u.y,
        }),
    )
  }

  update() {
    for (const user of this._users) user.update()
  }

  draw(ctx: CanvasRenderingContext2D) {
    for (const user of this._users) user.draw(ctx)
  }

  changePos(id: string, targetX: number, targetY: number) {
    const user = this._users.find((u) => u.equalId(id))
    if (!user) return

    user.changePos(targetX, targetY)
  }

  addUser(roomUser: RoomFragment['room']['users'][0]) {
    this._users.push(
      new User({
        id: roomUser.id,
        avatarUrl: roomUser.avatarUrl,
        currentX: roomUser.x,
        currentY: roomUser.y,
      }),
    )
  }

  deleteUser(id: string) {
    this._users = this._users.filter((u) => !u.equalId(id))
  }
}
