import { ROOM_SIZE, DEFAULT_ROOM_BG_COLOR } from '@/constants'

type RoomScreenOpts = {
  bgImage?: string
  bgColor?: string
}

class RoomScreen {
  private _bgImageEl: HTMLImageElement
  private _isBgLoaded: boolean
  private _bgColor: string

  constructor(opts?: RoomScreenOpts) {
    this._bgImageEl = new Image()
    this._isBgLoaded = false

    if (opts && opts.bgImage) {
      this._bgImageEl.src = opts.bgImage
      this._bgImageEl.onload = () => {
        this._isBgLoaded = true
      }
    }

    this._bgColor = (opts && opts.bgColor) || DEFAULT_ROOM_BG_COLOR
  }

  update() {
    console.log('room update')
  }

  draw(ctx: CanvasRenderingContext2D) {
    if (this._isBgLoaded) {
      ctx.drawImage(
        this._bgImageEl,
        0,
        0,
        ROOM_SIZE.WIDTH,
        (this._bgImageEl.height * ROOM_SIZE.WIDTH) / this._bgImageEl.width,
      )
    } else {
      ctx.fillStyle = this._bgColor
      ctx.fillRect(0, 0, ROOM_SIZE.WIDTH, ROOM_SIZE.HEIGHT)
    }
  }
}

export default RoomScreen
