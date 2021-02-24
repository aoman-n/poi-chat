type RoomScreenOpts = {
  bgImage: string
}

class RoomScreen {
  static WIDTH = 800
  static HEIGHT = 600
  private _bgImageEl: HTMLImageElement
  private _isBgLoaded: boolean

  constructor(opts: RoomScreenOpts) {
    this._isBgLoaded = false
    this._bgImageEl = new Image()
    this._bgImageEl.src = opts.bgImage
    this._bgImageEl.onload = () => {
      this._isBgLoaded = true
    }
  }

  update() {
    console.log('room update')
  }

  draw(ctx: CanvasRenderingContext2D) {
    if (!this._isBgLoaded) return

    ctx.drawImage(
      this._bgImageEl,
      0,
      0,
      RoomScreen.WIDTH,
      (this._bgImageEl.height * RoomScreen.WIDTH) / this._bgImageEl.width,
    )
  }
}

export default RoomScreen
