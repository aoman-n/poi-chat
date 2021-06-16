import React from 'react'
import { Story, Meta } from '@storybook/react'
import { ROOM_SIZE } from '@/constants'
import { mockMessages, mockUsers } from '@/mocks'
import Playground, { PlaygroundProps } from './presentation'
import { UserManager } from '@/utils/painter'

const mockUserManager = new UserManager(mockUsers)

export default {
  title: 'RoomPage/Playground',
  component: Playground,
} as Meta

const Template: Story<PlaygroundProps> = (args) => (
  <div style={{ width: `${ROOM_SIZE.WIDTH}px` }}>
    <Playground {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  messages: mockMessages,
  hasMoreMessage: false,
  rooomScreenProps: {
    userManager: mockUserManager,
    handleMovePos: () => {},
  },
  handleMoreMessage: () => {},
  moreLoading: false,
  handleRemoveBalloon: () => {},
}
