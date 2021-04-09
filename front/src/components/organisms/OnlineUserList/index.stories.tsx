import React from 'react'
import { Story, Meta } from '@storybook/react'
import OnlineUserList, { OnlineUserListProps, User } from '.'

export default {
  title: 'organisms/OnlineUserList',
  component: OnlineUserList,
} as Meta

const mockUsers: User[] = [
  {
    id: '1',
    name: 'サンプルユーザー1',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '2',
    name: 'サンプルユーザー2',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '3',
    name: 'サンプルユーザー3',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '3',
    name: 'サンプルユーザー4',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
]

const Template: Story<OnlineUserListProps> = (args) => (
  <div className="w-72">
    <OnlineUserList {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  users: mockUsers,
}
