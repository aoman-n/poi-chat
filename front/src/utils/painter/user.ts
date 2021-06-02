export type UserInfo = {
  id: string
  avatarUrl: string
  currentX: number
  currentY: number
}

type UserOpts = {
  id: string
  avatarUrl: string
  currentX: number
  currentY: number
}

export class User {
  private _sizeW = 50
  private _sizeH = 50
  private _id: string
  private _currentX: number
  private _currentY: number
  private _targetX: number
  private _targetY: number
  private _avatarEl: HTMLImageElement
  private _isAvatarLoaded: boolean

  constructor(opts: UserOpts) {
    this._id = opts.id
    this._currentX = opts.currentX
    this._currentY = opts.currentY
    this._targetX = opts.currentX
    this._targetY = opts.currentY
    this._isAvatarLoaded = false
    this._avatarEl = new Image()
    this._avatarEl.src = opts.avatarUrl
    this._avatarEl.onload = () => {
      this._isAvatarLoaded = true
    }
  }

  changePos(targetX: number, targetY: number) {
    this._targetX = targetX
    this._targetY = targetY
  }

  draw(ctx: CanvasRenderingContext2D) {
    if (!this._isAvatarLoaded) return

    ctx.drawImage(
      this._avatarEl,
      this._currentX - this._sizeW / 2,
      this._currentY - this._sizeH / 2,
      this._sizeW,
      this._sizeH,
    )
  }

  update() {
    this._currentX = this._currentX + (this._targetX - this._currentX) / 8
    this._currentY = this._currentY + (this._targetY - this._currentY) / 8
  }

  equalId(id: string) {
    return id === this._id
  }

  equal(user: User) {
    return user._id === this._id
  }
}
