import React from 'react'
import Frame from './Frame'
import Playground, { PlaygroundProps } from './Playground'
import Settings, { SettingsProps } from './Settings'
import MessageList, { MessageListProps } from './MessageList'
import MessageForm, { MessageFormProps } from './MessageForm'

export type RoomPageProps = {
  playgroundProps: PlaygroundProps
  settingsProps: SettingsProps
  messageListProps: MessageListProps
  messageFormProps: MessageFormProps
}

const RoomPagePresenter: React.VFC<RoomPageProps> = ({
  playgroundProps,
  settingsProps,
  messageListProps,
  messageFormProps,
}) => {
  return (
    <div style={{ width: '1180px' }}>
      <Frame
        screen={<Playground {...playgroundProps} />}
        settings={<Settings {...settingsProps} />}
        messages={<MessageList {...messageListProps} />}
        messageFrom={<MessageForm {...messageFormProps} />}
      />
    </div>
  )
}

export default RoomPagePresenter
