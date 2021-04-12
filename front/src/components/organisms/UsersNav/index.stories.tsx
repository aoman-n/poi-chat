import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockOnlineUsers } from '@/mocks'
import UsersNav, { UsersNavProps } from './presentation'

export default {
  title: 'organisms/UsersNav',
  component: UsersNav,
} as Meta

const Template: Story<UsersNavProps> = (args) => (
  <div className="w-64">
    <UsersNav {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  profile: {
    name: 'sample name',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  onlineUserList: {
    users: mockOnlineUsers,
  },
}
