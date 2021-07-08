import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockOnlineUsers } from '@/mocks'
import Header, { HeaderProps } from './presenter'

export default {
  title: 'domainParts/Header',
  component: Header,
} as Meta<HeaderProps>

const Template: Story<HeaderProps> = (args) => (
  <div className="h-16 bg-white">
    <Header {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  profile: {
    name: 'sample name',
    avatarUrl:
      'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png',
  },
  onlineUserList: {
    users: mockOnlineUsers,
  },
}
