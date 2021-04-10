import React from 'react'
import useScrollBottom from '@/hooks/useScrollBottom'
import styles from './index.module.scss'

const defaultBgColor = '#5f9ea0'

export type Message = {
  id: string
  userId: string
  userName: string
  userAvatarUrl: string
  body: string
  createdAt: string
}

export type PlaygroundProps = {
  bgColor?: string
  messages: Message[]
  handleSubmitMessage: (e: React.FormEvent<HTMLFormElement>) => void
}

const Playground: React.FC<PlaygroundProps> = ({
  bgColor = defaultBgColor,
  messages,
  handleSubmitMessage,
}) => {
  const { scrollAreaRef, endItemRef } = useScrollBottom(messages)

  return (
    <div>
      <div
        className={[styles.screenFrame].join(' ')}
        style={{ backgroundColor: bgColor }}
      >
        playground
      </div>
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
            'overflow-y-scroll',
            'text-sm',
            'space-y-2',
          ].join(' ')}
        >
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
        onSubmit={handleSubmitMessage}
      >
        <input
          id="username"
          type="text"
          className="px-3 py-2 bg-white focus:outline-none text-sm flex-grow"
          placeholder="コメントを入力"
        />
        <button className="px-4 bg-green-600 text-white font-medium opacity-80 hover:opacity-100 duration-100 focus:outline-none">
          送信
        </button>
      </form>
    </div>
  )
}

export default Playground
