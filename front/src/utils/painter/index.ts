// import { EventEmitter } from 'events'

// export type UserData = {
//   id: string
//   x: number
//   y: number
// }

// 中で変更を検知できればよい？

// ユーザーの位置情報をサーバーからサブスクライブ
// subcribe((data: UserData) => {
//   userManager.updates(data)
// })

// export class User extends EventEmitter {
//   private id: string
//   private name: string
//   private x: number
//   private y: number

//   constructor(id: string, name: string, opts?: { x: number; y: number }) {
//     super()
//     this.id = id
//     this.name = name
//     this.x = opts?.x || 100
//     this.y = opts?.y || 100
//   }

//   move(data: UserData) {
//     this.emit('move', data)
//   }

//   onMove(callback: (data: UserData) => void) {
//     this.on('move', callback)
//   }
// }

// export default User
