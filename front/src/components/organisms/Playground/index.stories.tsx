import React from 'react'
import { Story, Meta } from '@storybook/react'
import { ROOM_SIZE } from '@/constants'
import { mockMessages, mockUserManager } from '@/mocks'
import Playground, { PlaygroundProps } from './presentation'

export default {
  title: 'organisms/Playground',
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
  handleSubmitMessage: (e: React.FormEvent<HTMLFormElement>) =>
    e.preventDefault(),
  userManager: mockUserManager,
}
