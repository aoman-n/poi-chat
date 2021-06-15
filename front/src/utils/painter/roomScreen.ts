import { ROOM_SIZE, DEFAULT_ROOM_BG_COLOR } from '@/constants'

type RoomScreenOpts = {
  bgImage?: string
  bgColor?: string
}

export class RoomScreenPainter {
  #bgImageEl: HTMLImageElement
  #isBgLoaded: boolean
  #bgColor: string

  constructor(opts?: RoomScreenOpts) {
    this.#bgImageEl = new Image()
    this.#isBgLoaded = false

    if (opts && opts.bgImage) {
      this.#bgImageEl.src = opts.bgImage
      this.#bgImageEl.onload = () => {
        this.#isBgLoaded = true
      }
    }

    this.#bgColor = (opts && opts.bgColor) || DEFAULT_ROOM_BG_COLOR
  }

  update() {
    console.log('room update')
  }

  draw(ctx: CanvasRenderingContext2D) {
    if (this.#isBgLoaded) {
      ctx.drawImage(
        this.#bgImageEl,
        0,
        0,
        ROOM_SIZE.WIDTH,
        (this.#bgImageEl.height * ROOM_SIZE.WIDTH) / this.#bgImageEl.width,
      )
    } else {
      ctx.fillStyle = this.#bgColor
      ctx.fillRect(0, 0, ROOM_SIZE.WIDTH, ROOM_SIZE.HEIGHT)
    }
  }
}
