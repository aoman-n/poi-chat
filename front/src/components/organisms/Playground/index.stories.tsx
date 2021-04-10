import React from 'react'
import { Story, Meta } from '@storybook/react'
import { mockMessages } from '@/mocks'
import Playground, { PlaygroundProps } from './presentation'

export default {
  title: 'organisms/Playground',
  component: Playground,
} as Meta

const Template: Story<PlaygroundProps> = (args) => (
  <div className="w-2/3">
    <Playground {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  messages: mockMessages,
  handleSubmitMessage: (e: React.FormEvent<HTMLFormElement>) =>
    e.preventDefault(),
}
