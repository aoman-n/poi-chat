import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockOnlineUsers } from '@/mocks'
import OnlineUserList, { OnlineUserListProps } from '.'

export default {
  title: 'organisms/OnlineUserList',
  component: OnlineUserList,
} as Meta

const Template: Story<OnlineUserListProps> = (args) => (
  <div className="w-72">
    <OnlineUserList {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  users: mockOnlineUsers,
}
