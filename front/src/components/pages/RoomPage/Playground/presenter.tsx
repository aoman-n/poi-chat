import React, { useCallback, useState } from 'react'
import { useScrollBottom, useReverseFetchMore } from '@/hooks'
import RoomScreen, {
  RoomScreenProps,
} from '@/components/pages/RoomPage/RoomScreen'
import Button from '@/components/parts/Button'
import { BalloonPosition } from '@/constants'
import { RoomFragment } from '@/graphql'

export type BalloonState = {
  hasBalloon: boolean
  position: BalloonPosition | null
}

export type PlaygroundProps = {
  messages: RoomFragment['room']['messages']['nodes']
  hasMoreMessage: boolean
  handleSubmitMessage: (values: { body: string }) => void
  rooomScreenProps: RoomScreenProps
  handleMoreMessage: () => void
  moreLoading: boolean
  handleChangeBalloonPos: (pos: BalloonPosition) => void
  handleRemoveBalloon: () => void
  balloonState: BalloonState
}

const Playground: React.FC<PlaygroundProps> = ({
  messages,
  hasMoreMessage,
  handleSubmitMessage,
  rooomScreenProps,
  handleMoreMessage,
  moreLoading,
  handleChangeBalloonPos,
  handleRemoveBalloon,
  balloonState,
}) => {
  const { scrollBottomRef } = useScrollBottom(messages)
  const { scrollTopRef, prevFirstItem, firstItemRef } = useReverseFetchMore(
    messages,
    handleMoreMessage,
    !moreLoading && hasMoreMessage,
  )
  const [inputMesage, setInputMessage] = useState('')

  const wrappedHandleSubmitMessage = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault()
      handleSubmitMessage({ body: inputMesage })
      setInputMessage('')
    },
    [handleSubmitMessage, inputMesage],
  )

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputMessage(e.target.value)
  }, [])

  return (
    <div>
      {/* アバタースクリーン。一旦決め打ちサイズで */}
      <RoomScreen {...rooomScreenProps} />

      {/* セッティングパネル */}
      <div className="pt-4 flex">
        <div>
          <h4 className="mb-2">吹き出し位置変更</h4>
          <div className="space-x-2 pb-2">
            <Button
              onClick={() => handleChangeBalloonPos('TOP_LEFT')}
              disabled={disabledBalloonPosButton(balloonState, 'TOP_LEFT')}
            >
              ↖左上
            </Button>
            <Button
              onClick={() => handleChangeBalloonPos('TOP_RIGHT')}
              disabled={disabledBalloonPosButton(balloonState, 'TOP_RIGHT')}
            >
              右上↗
            </Button>
          </div>
          <div className="space-x-2">
            <Button
              onClick={() => handleChangeBalloonPos('BOTTOM_LEFT')}
              disabled={disabledBalloonPosButton(balloonState, 'BOTTOM_LEFT')}
            >
              ↙左下
            </Button>
            <Button
              onClick={() => handleChangeBalloonPos('BOTTOM_RIGHT')}
              disabled={disabledBalloonPosButton(balloonState, 'BOTTOM_RIGHT')}
            >
              右下↘
            </Button>
          </div>
        </div>
        <div className="ml-auto">
          <Button
            onClick={handleRemoveBalloon}
            color="red"
            disabled={!balloonState.hasBalloon}
          >
            吹き出しを消す
          </Button>
        </div>
      </div>

      {/* コメント欄 */}
      <div className={['mt-6'].join(' ')}>
        <h4 className={['mb-1', 'text-gray-900'].join(' ')}>コメント欄</h4>
        <ul
          className={[
            'py-4',
            'px-4',
            'border',
            'border-gray-300',
            'bg-white',
            'h-52',
            'overflow-y-auto',
            'text-sm',
          ].join(' ')}
        >
          <div ref={scrollTopRef} />
          {moreLoading && <div>Now Loading...</div>}
          {messages.map((message) => {
            if (prevFirstItem && prevFirstItem.id === message.id) {
              return (
                <li key={message.id} ref={firstItemRef}>
                  <Message message={message} />
                </li>
              )
            } else {
              return (
                <li key={message.id} className="mt-2">
                  <Message message={message} />
                </li>
              )
            }
          })}
          <div ref={scrollBottomRef} />
        </ul>
      </div>
      <form
        className={['mt-6', 'text-gray-900', 'flex'].join(' ')}
        onSubmit={wrappedHandleSubmitMessage}
      >
        <input
          id="username"
          type="text"
          className="px-3 py-2 bg-white focus:outline-none text-sm flex-grow"
          placeholder="コメントを入力"
          autoComplete="off"
          onChange={handleChange}
          value={inputMesage}
        />
        <button className="px-4 bg-green-600 text-white font-medium opacity-80 hover:opacity-100 duration-100 focus:outline-none">
          送信
        </button>
      </form>
    </div>
  )
}

const disabledBalloonPosButton = (
  balloonState: BalloonState,
  buttonType: BalloonPosition,
) => {
  return !balloonState.hasBalloon || balloonState.position === buttonType
}

type MessageProps = {
  message: RoomFragment['room']['messages']['nodes'][0]
}

const Message: React.FC<MessageProps> = ({ message }) => {
  return (
    <div className="m-0">
      <span className="text-gray-400 font-medium pr-1.5">
        {message.userName}:
      </span>
      <span>{message.body}</span>
    </div>
  )
}

export default Playground
