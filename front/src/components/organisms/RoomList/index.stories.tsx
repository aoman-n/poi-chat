import React from 'react'
import { Story, Meta } from '@storybook/react'
import RoomList, { RoomListProps, Room } from '.'

export default {
  title: 'organisms/RoomList',
  component: RoomList,
} as Meta

const mockRooms: Room[] = [
  {
    id: '1',
    name: 'サンプルチャットルーム1',
    userCount: 8,
  },
  {
    id: '2',
    name: 'サンプルチャットルーム2',
    userCount: 10,
  },
  {
    id: '3',
    name: 'サンプルチャットルーム3',
    userCount: 3,
  },
  {
    id: '4',
    name: 'サンプルチャットルーム4',
    userCount: 60,
  },
]

const Template: Story<RoomListProps> = (args) => (
  <div style={{ width: '600px' }}>
    <RoomList {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  rooms: mockRooms,
}
