import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockMessages } from '@/mocks'
import MessageList, { MessageListProps } from '.'

export default {
  title: 'RoomPage/MesasgeList',
  component: MessageList,
} as Meta<MessageListProps>

const Template: Story<MessageListProps> = (args) => (
  <div style={{ width: '340px' }}>
    <MessageList {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  handleMoreMessage: () => {},
  moreLoading: false,
  messages: mockMessages,
  hasMoreMessage: false,
}
