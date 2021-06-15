import React, { useCallback, useState, forwardRef } from 'react'
import { useScrollBottom, useReverseFetchMore } from '@/hooks'
import RoomScreen, { RoomScreenProps } from '@/components/organisms/RoomScreen'
import { RoomFragment, BalloonPosition } from '@/graphql'

export type PlaygroundProps = {
  messages: RoomFragment['room']['messages']['nodes']
  hasMoreMessage: boolean
  handleSubmitMessage: (values: { body: string }) => void
  rooomScreenProps: RoomScreenProps
  handleMoreMessage: () => void
  moreLoading: boolean
  handleChangeBalloonPos: (pos: BalloonPosition) => void
  handleRemoveBalloon: () => void
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
          <div className="space-x-4 pb-4">
            <button
              className="outline-none bg-transparent duration-75 hover:bg-gray-500 hover:text-white text-gray-800 py-1 w-20 border border-gray-500 hover:border-transparent rounded"
              onClick={() => handleChangeBalloonPos(BalloonPosition.TopLeft)}
            >
              ↖左上
            </button>
            <button
              className="focus:outline-none bg-transparent duration-75 hover:bg-gray-500 hover:text-white text-gray-800 py-1 w-20 border border-gray-500 hover:border-transparent rounded"
              onClick={() => handleChangeBalloonPos(BalloonPosition.TopRight)}
            >
              右上↗
            </button>
          </div>
          <div className="space-x-4">
            <button
              className="focus:outline-none bg-transparent duration-75 hover:bg-gray-500 hover:text-white text-gray-800 py-1 w-20 border border-gray-500 hover:border-transparent rounded"
              onClick={() => handleChangeBalloonPos(BalloonPosition.BottomLeft)}
            >
              ↙左下
            </button>
            <button
              className="focus:outline-none bg-transparent duration-75 hover:bg-gray-500 hover:text-white text-gray-800 py-1 w-20 border border-gray-500 hover:border-transparent rounded"
              onClick={() =>
                handleChangeBalloonPos(BalloonPosition.BottomRight)
              }
            >
              右下↘
            </button>
          </div>
        </div>
        <div className="ml-auto">
          <button
            className="focus:outline-none duration-75 bg-red-500 text-white py-2 px-6 rounded hover:bg-red-600"
            onClick={handleRemoveBalloon}
          >
            吹き出しを消す
          </button>
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
            'space-y-2',
          ].join(' ')}
        >
          <div ref={scrollTopRef} />
          {moreLoading && <div>Now Loading...</div>}
          {messages.map((message) => {
            if (prevFirstItem && prevFirstItem.id === message.id) {
              return (
                <Message
                  key={message.id}
                  message={message}
                  ref={firstItemRef}
                />
              )
            } else {
              return <Message key={message.id} message={message} />
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

type MessageProps = {
  message: RoomFragment['room']['messages']['nodes'][0]
}

const Message = forwardRef<HTMLLIElement, MessageProps>(({ message }, ref) => {
  return (
    <li className="m-0" ref={ref}>
      <span className="text-gray-400 font-medium pr-1.5">
        {message.userName}:
      </span>
      <span>{message.body}</span>
    </li>
  )
})

export default Playground
