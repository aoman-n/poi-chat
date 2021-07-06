import React from 'react'
import { Story, Meta } from '@storybook/react'
import { ROOM_SCREEN_SIZE, ROOM_BG_IMAGES } from '@/constants'
import { mockMessages, mockUsers } from '@/mocks'
import Playground, { PlaygroundProps } from './presenter'
import { UserManager } from '@/utils/painter'

const mockUserManager = new UserManager(mockUsers)

export default {
  title: 'RoomPage/Playground',
  component: Playground,
} as Meta<PlaygroundProps>

const Template: Story<PlaygroundProps> = (args) => (
  <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
    <Playground {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  messages: mockMessages,
  hasMoreMessage: false,
  handleMoreMessage: () => {},
  handleSubmitMessage: () => {},
  rooomScreenProps: {
    userManager: mockUserManager,
    handleMovePos: () => {},
    bgColor: '',
    bgUrl: ROOM_BG_IMAGES[0].url,
  },
  moreLoading: false,
  handleChangeBalloonPos: () => {},
  handleRemoveBalloon: () => {},
  balloonState: {
    hasBalloon: true,
    position: 'TOP_RIGHT',
  },
}
