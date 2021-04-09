import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockRooms } from '@/mocks'
import RoomList, { RoomListProps } from '.'

export default {
  title: 'organisms/RoomList',
  component: RoomList,
} as Meta

const Template: Story<RoomListProps> = (args) => (
  <div style={{ width: '600px' }}>
    <RoomList {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  rooms: mockRooms,
}
