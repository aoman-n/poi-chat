import React, { useCallback, useState } from 'react'
import { Waypoint } from 'react-waypoint'
import { useScrollBottom, usePrevScroll } from '@/hooks'
import RoomScreen, { RoomScreenProps } from '@/components/organisms/RoomScreen'
import { RoomFragment } from '@/graphql'

export type PlaygroundProps = {
  messages: RoomFragment['room']['messages']['nodes']
  hasMoreMessage: boolean
  handleSubmitMessage: (values: { body: string }) => void
  rooomScreenProps: RoomScreenProps
  handleMoreMessage: () => void
  moreLoading: boolean
}

const Playground: React.FC<PlaygroundProps> = ({
  messages,
  hasMoreMessage,
  handleSubmitMessage,
  rooomScreenProps,
  handleMoreMessage,
  moreLoading,
}) => {
  const [inputMesage, setInputMessage] = useState('')
  const { scrollAreaRef, endItemRef, isBottom } = useScrollBottom(messages)
  const { setOldScrollHeight } = usePrevScroll(
    scrollAreaRef,
    messages,
    !isBottom,
  )

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

  const wrappedHandleMoreMessage = useCallback(() => {
    if (scrollAreaRef.current) {
      setOldScrollHeight(scrollAreaRef.current.scrollHeight)
    }
    handleMoreMessage()
  }, [handleMoreMessage, setOldScrollHeight, scrollAreaRef])

  return (
    <div>
      <div>
        <button onClick={handleMoreMessage}>more</button>
      </div>
      {/* RoomScreenは一旦決め打ちサイズで */}
      <RoomScreen {...rooomScreenProps} />

      <div className={['mt-6'].join(' ')}>
        <h4 className={['mb-1', 'text-gray-900'].join(' ')}>コメント欄</h4>
        <ul
          ref={scrollAreaRef}
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
          {hasMoreMessage && !moreLoading && (
            <Waypoint onEnter={wrappedHandleMoreMessage} />
          )}
          {moreLoading && <div>Now Loading...</div>}
          {messages.map((message) => (
            <li key={message.id} className="m-0" ref={endItemRef}>
              <span className="text-gray-400 font-medium pr-1.5">
                {message.userName}:
              </span>
              <span>{message.body}</span>
            </li>
          ))}
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

export default Playground
