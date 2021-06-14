import { valueOf } from '@/utils/types'

type UserOpts = {
  id: string
  name: string
  avatarUrl: string
  currentX: number
  currentY: number
  balloonPos?: valueOf<typeof BALLOON_POSITIONS>
  lastMessage?: string
}

export const BALLOON_POSITIONS = {
  TOP_RIGHT: 'TOP_RIGHT',
  TOP_LEFT: 'TOP_LEFT',
  BOTTOM_RIGHT: 'BOTTOM_RIGHT',
  BOTTOM_LEFT: 'BOTTOM_LEFT',
} as const

export type BalloonPosition = valueOf<typeof BALLOON_POSITIONS>

export class User {
  #avatarSize = 54
  #lineHeight = 1.3
  #textMaxWidth = 160
  #balloonFontSize = 14
  #balloonFontColor = '#333333'
  #balloonPaddingX = 12
  #balloonPaddingY = 12
  #balloonBgColor = '#fff'
  #balloonStrokeColor = '#696969'
  #balloonStrokeWidth = 1
  #nameFontSize = 18
  #nameFontColor = '#1a1a1a'

  #id: string
  #name: string
  #currentX: number
  #currentY: number
  #targetX: number
  #targetY: number
  #isLoadedAvatar: boolean
  #avatarCvs: HTMLCanvasElement
  #balloonPos: BalloonPosition
  #lastMessage: string

  constructor(opts: UserOpts) {
    this.#id = opts.id
    this.#name = opts.name
    this.#currentX = opts.currentX
    this.#currentY = opts.currentY
    this.#targetX = opts.currentX
    this.#targetY = opts.currentY
    this.#balloonPos = opts.balloonPos || BALLOON_POSITIONS.TOP_RIGHT
    this.#lastMessage = opts.lastMessage || ''

    // アバター画像を円形に切り取るために別キャンバスを作成し画像をトリミング
    this.#avatarCvs = document.createElement('canvas')
    this.#avatarCvs.width = this.#avatarSize
    this.#avatarCvs.height = this.#avatarSize
    const avatarCtx = this.#avatarCvs.getContext(
      '2d',
    ) as CanvasRenderingContext2D
    avatarCtx.beginPath()
    avatarCtx.arc(
      this.#avatarSize / 2,
      this.#avatarSize / 2,
      this.#avatarSize / 2,
      0,
      Math.PI * 2,
      false,
    )
    avatarCtx.clip()
    // ---------------------------------------------

    this.#isLoadedAvatar = false
    const img = new Image()
    img.src = opts.avatarUrl
    img.onload = () => {
      avatarCtx.drawImage(img, 0, 0, this.#avatarSize, this.#avatarSize)
      this.#isLoadedAvatar = true
    }
  }

  // target位置を更新する
  changePos(targetX: number, targetY: number) {
    this.#targetX = targetX
    this.#targetY = targetY
  }

  // currentX,currentYをもとにavatarを描画する
  draw(ctx: CanvasRenderingContext2D) {
    if (!this.#isLoadedAvatar) return

    // アバター画像の描画
    ctx.drawImage(
      this.#avatarCvs,
      this.#currentX - this.#avatarSize / 2,
      this.#currentY - this.#avatarSize / 2,
      this.#avatarSize,
      this.#avatarSize,
    )

    // 名前の描画
    this.#drawName(ctx)

    // 吹き出しメッセージの描画
    this.#drawBalloonAndMessage(ctx)
  }

  // 現在地をtargetに向かって移動させ更新する
  update() {
    this.#currentX = this.#currentX + (this.#targetX - this.#currentX) / 8
    this.#currentY = this.#currentY + (this.#targetY - this.#currentY) / 8
  }

  equalId(id: string) {
    return id === this.#id
  }

  equal(user: User) {
    return user.#id === this.#id
  }

  changeBalloonPos(pos: BalloonPosition) {
    this.#balloonPos = pos
  }

  updateMessage(message: string) {
    this.#lastMessage = message
  }

  #drawName(ctx: CanvasRenderingContext2D) {
    ctx.font = this.#nameFontSize + 'px Arial, meiryo, sans-serif'
    ctx.fillStyle = this.#nameFontColor
    ctx.textAlign = 'center'
    ctx.textBaseline = 'top'
    ctx.fillText(
      this.#name,
      this.#currentX,
      this.#currentY + this.#avatarSize / 2 + 8,
    )
  }

  // lastMessageを吹き出しとして描画する
  #drawBalloonAndMessage(ctx: CanvasRenderingContext2D) {
    if (this.#lastMessage === '') return

    ctx.font = this.#balloonFontSize + 'px Arial, meiryo, sans-serif'
    const splittedMessages = this.#splitMultilineText(ctx, this.#lastMessage)

    // 吹き出しの描画
    const {
      balloonStartPosX,
      balloonStartPosY,
      balloonWidth,
      balloonHeight,
    } = this.#getBalloonConfig(ctx, splittedMessages)
    this.#fillAndStrokeBalloonRect(
      ctx,
      balloonStartPosX,
      balloonStartPosY,
      balloonWidth,
      balloonHeight,
    )

    // メッセージの描画
    ctx.beginPath()
    ctx.fillStyle = this.#balloonFontColor
    ctx.textAlign = 'left'
    ctx.textBaseline = 'top'
    for (const [index, line] of splittedMessages.entries()) {
      let addedLineY = this.#balloonPaddingY
      if (index >= 1) {
        addedLineY += this.#balloonFontSize * this.#lineHeight * index
      }
      ctx.fillText(
        line,
        balloonStartPosX + this.#balloonPaddingX,
        balloonStartPosY + addedLineY,
      )
    }
  }

  #splitMultilineText(ctx: CanvasRenderingContext2D, text: string) {
    const len = text.length
    const strArray: string[] = []
    let tmp = ''
    let i = 0

    if (len < 1) {
      //textの文字数が0だったら終わり
      return strArray
    }

    for (i = 0; i < len; i++) {
      const c = text.charAt(i) //textから１文字抽出
      if (c == '\n') {
        /* 改行コードの場合はそれまでの文字列を配列にセット */
        strArray.push(tmp)
        tmp = ''

        continue
      }

      /* ctxの現在のフォントスタイルで描画したときの長さを取得 */
      if (ctx.measureText(tmp + c).width <= this.#textMaxWidth) {
        /* 指定幅を超えるまでは文字列を繋げていく */
        tmp += c
      } else {
        /* 超えたら、それまでの文字列を配列にセット */
        strArray.push(tmp)
        tmp = c
      }
    }

    /* 繋げたままの分があれば回収 */
    if (tmp.length > 0) strArray.push(tmp)

    return strArray
  }

  #getBalloonConfig(ctx: CanvasRenderingContext2D, messages: string[]) {
    // テキストの高さ、幅
    const [lineWidth, linesHeight] = ((): [number, number] => {
      if (messages.length === 0) {
        return [0, 0]
      } else if (messages.length === 1) {
        return [ctx.measureText(messages[0]).width, this.#balloonFontSize]
      } else {
        const linesHeight =
          (messages.length - 1) * (this.#balloonFontSize * this.#lineHeight) +
          this.#balloonFontSize
        return [this.#textMaxWidth, linesHeight]
      }
    })()

    // 吹き出しのサイズ
    const balloonWidth = lineWidth + this.#balloonPaddingX * 2
    const balloonHeight = linesHeight + this.#balloonPaddingY * 2

    // 吹き出し左上開始位置
    const [balloonStartPosX, balloonStartPosY] = ((): [number, number] => {
      switch (this.#balloonPos) {
        case 'TOP_LEFT':
          return [
            this.#currentX - balloonWidth - this.#avatarSize / 2,
            this.#currentY -
              linesHeight -
              this.#balloonPaddingY * 2 -
              this.#avatarSize / 2,
          ]
        case 'TOP_RIGHT':
          return [
            this.#currentX + this.#avatarSize / 2,
            this.#currentY -
              linesHeight -
              this.#balloonPaddingY * 2 -
              this.#avatarSize / 2,
          ]
        case 'BOTTOM_LEFT':
          return [
            this.#currentX - balloonWidth - this.#avatarSize / 2,
            this.#currentY + this.#avatarSize / 2,
          ]
        case 'BOTTOM_RIGHT':
          return [
            this.#currentX + this.#avatarSize / 2,
            this.#currentY + this.#avatarSize / 2,
          ]
      }
    })()

    return {
      balloonStartPosX,
      balloonStartPosY,
      balloonWidth,
      balloonHeight,
    }
  }

  #fillAndStrokeBalloonRect(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    w: number,
    h: number,
  ) {
    const r = 10
    const bl = 16
    const br = 36
    const bh = 16

    // 描画情報
    ctx.beginPath()
    ctx.fillStyle = this.#balloonBgColor
    ctx.strokeStyle = this.#balloonStrokeColor
    ctx.lineWidth = this.#balloonStrokeWidth

    // TODO: Refactor
    switch (this.#balloonPos) {
      case 'BOTTOM_LEFT':
        this.#createBottomLeftBalloonRoundRectPath(
          ctx,
          x,
          y,
          w,
          h,
          r,
          bl,
          br,
          bh,
        )
        break
      case 'BOTTOM_RIGHT':
        this.#createBottomRightBalloonRoundRectPath(
          ctx,
          x,
          y,
          w,
          h,
          r,
          bl,
          br,
          bh,
        )
        break
      case 'TOP_LEFT':
        this.#createTopLeftBalloonRoundRectPath(ctx, x, y, w, h, r, bl, br, bh)
        break
      case 'TOP_RIGHT':
        this.#createTopRightBalloonRoundRectPath(ctx, x, y, w, h, r, bl, br, bh)
        break
    }

    ctx.fill()
    ctx.stroke()
  }

  #createTopLeftBalloonRoundRectPath(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    w: number,
    h: number,
    r: number,
    bl: number,
    br: number,
    bh: number,
  ) {
    ctx.beginPath()
    ctx.moveTo(x + r, y)
    ctx.lineTo(x + w - r, y)
    ctx.arc(x + w - r, y + r, r, Math.PI * (3 / 2), 0, false)
    ctx.lineTo(x + w, y + h - r)
    ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * (1 / 2), false)

    // 吹き出し矢印開始
    ctx.lineTo(x + w - r - bl, y + h)
    ctx.lineTo(x + w, y + h + bh)
    ctx.lineTo(x + w - r - br, y + h)
    ctx.lineTo(x + r, y + h)
    // 吹き出し矢印終了

    ctx.arc(x + r, y + h - r, r, Math.PI * (1 / 2), Math.PI, false)
    ctx.lineTo(x, y + r)
    ctx.arc(x + r, y + r, r, Math.PI, Math.PI * (3 / 2), false)
    ctx.closePath()
  }

  #createBottomLeftBalloonRoundRectPath(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    w: number,
    h: number,
    r: number,
    bl: number,
    br: number,
    bh: number,
  ) {
    ctx.beginPath()
    ctx.moveTo(x + r, y)

    // 吹き出し矢印開始
    ctx.lineTo(x + w - r - br, y)
    ctx.lineTo(x + w, y - bh)
    ctx.lineTo(x + w - r - bl, y)
    ctx.lineTo(x + w - r, y)
    // 吹き出し矢印終了

    ctx.arc(x + w - r, y + r, r, Math.PI * (3 / 2), 0, false)
    ctx.lineTo(x + w, y + h - r)
    ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * (1 / 2), false)
    ctx.lineTo(x + r, y + h)

    ctx.arc(x + r, y + h - r, r, Math.PI * (1 / 2), Math.PI, false)
    ctx.lineTo(x, y + r)
    ctx.arc(x + r, y + r, r, Math.PI, Math.PI * (3 / 2), false)
    ctx.closePath()
  }

  #createBottomRightBalloonRoundRectPath(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    w: number,
    h: number,
    r: number,
    bl: number,
    br: number,
    bh: number,
  ) {
    ctx.beginPath()
    ctx.moveTo(x + r, y)

    // 吹き出し矢印開始
    ctx.lineTo(x + bl, y)
    ctx.lineTo(x, y - bh)
    ctx.lineTo(x + br, y)
    ctx.lineTo(x + w - r, y)
    // 吹き出し矢印終了

    ctx.arc(x + w - r, y + r, r, Math.PI * (3 / 2), 0, false)
    ctx.lineTo(x + w, y + h - r)
    ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * (1 / 2), false)
    ctx.lineTo(x + r, y + h)

    ctx.arc(x + r, y + h - r, r, Math.PI * (1 / 2), Math.PI, false)
    ctx.lineTo(x, y + r)
    ctx.arc(x + r, y + r, r, Math.PI, Math.PI * (3 / 2), false)
    ctx.closePath()
  }

  #createTopRightBalloonRoundRectPath(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    w: number,
    h: number,
    r: number,
    bl: number,
    br: number,
    bh: number,
  ) {
    ctx.beginPath()
    ctx.moveTo(x + r, y)
    ctx.lineTo(x + w - r, y)
    ctx.arc(x + w - r, y + r, r, Math.PI * (3 / 2), 0, false)
    ctx.lineTo(x + w, y + h - r)
    ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * (1 / 2), false)

    // 吹き出し矢印開始
    ctx.lineTo(x + br, y + h)
    ctx.lineTo(x, y + h + bh)
    ctx.lineTo(x + bl, y + h)
    ctx.lineTo(x + r, y + h)
    // 吹き出し矢印終了

    ctx.arc(x + r, y + h - r, r, Math.PI * (1 / 2), Math.PI, false)
    ctx.lineTo(x, y + r)
    ctx.arc(x + r, y + r, r, Math.PI, Math.PI * (3 / 2), false)
    ctx.closePath()
  }
}
