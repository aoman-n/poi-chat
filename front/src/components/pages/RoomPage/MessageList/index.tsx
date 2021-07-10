import React from 'react'
import { useScrollBottom, useReverseFetchMore } from '@/hooks'
import { RoomFragment } from '@/graphql'

export type MessageListProps = {
  handleMoreMessage: () => void
  moreLoading: boolean
  messages: RoomFragment['room']['messages']['nodes']
  hasMoreMessage: boolean
}

const MessageList: React.VFC<MessageListProps> = ({
  messages,
  handleMoreMessage,
  moreLoading,
  hasMoreMessage,
}) => {
  const { scrollBottomRef, parentRef } = useScrollBottom(messages)
  const { scrollTopRef, prevFirstItem, firstItemRef } = useReverseFetchMore(
    messages,
    handleMoreMessage,
    !moreLoading && hasMoreMessage,
  )

  return (
    <div>
      <div className={['mt-6'].join(' ')}>
        <h4 className={['mb-2', 'text-gray-900'].join(' ')}>コメント欄</h4>
        <ul
          ref={parentRef}
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
    </div>
  )
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

export default MessageList
