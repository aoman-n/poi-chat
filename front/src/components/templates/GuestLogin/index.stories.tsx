import React from 'react'
import { Story, Meta } from '@storybook/react'
import GuestLoginTemplate, { GuestLoginProps } from '.'

export default {
  title: 'templates/GuestLogin',
  component: GuestLoginTemplate,
} as Meta

const Template: Story<GuestLoginProps> = (args) => (
  <GuestLoginTemplate {...args} />
)

export const Default = Template.bind({})
Default.args = {
  noop: 'ゲストログイン',
}
