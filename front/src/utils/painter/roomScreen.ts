import { ROOM_SCREEN_SIZE, DEFAULT_ROOM_BG_COLOR } from '@/constants'

type RoomScreenOpts = {
  bgUrl?: string
  bgColor?: string
}

export class RoomScreenPainter {
  #bgImageEl: HTMLImageElement | null
  #isBgLoaded: boolean
  #bgColor: string

  constructor(opts?: RoomScreenOpts) {
    this.#isBgLoaded = false

    if (opts && opts.bgUrl) {
      this.#bgImageEl = new Image()
      this.#bgImageEl.src = opts.bgUrl
      this.#bgImageEl.onload = () => {
        this.#isBgLoaded = true
      }
      this.#bgImageEl.onerror = () => {
        this.#bgImageEl = null
      }
    } else {
      this.#bgImageEl = null
    }

    this.#bgColor = (opts && opts.bgColor) || DEFAULT_ROOM_BG_COLOR
  }

  update() {
    console.log('room update')
  }

  draw(ctx: CanvasRenderingContext2D) {
    if (this.#bgImageEl) {
      if (!this.#isBgLoaded) return
      ctx.drawImage(
        this.#bgImageEl,
        0,
        0,
        ROOM_SCREEN_SIZE.WIDTH,
        (this.#bgImageEl.height * ROOM_SCREEN_SIZE.WIDTH) /
          this.#bgImageEl.width,
      )
    } else {
      ctx.fillStyle = this.#bgColor
      ctx.fillRect(0, 0, ROOM_SCREEN_SIZE.WIDTH, ROOM_SCREEN_SIZE.HEIGHT)
    }
  }

  get isInitialized() {
    if (!this.#bgImageEl) {
      return true
    }

    return this.#isBgLoaded
  }
}
