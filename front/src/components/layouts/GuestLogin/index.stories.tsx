import React from 'react'
import { Story, Meta } from '@storybook/react'
import GuestLoginTemplate, { GuestLoginProps } from '.'

export default {
  title: 'layouts/GuestLogin',
  component: GuestLoginTemplate,
} as Meta

const Template: Story<GuestLoginProps> = (args) => (
  <GuestLoginTemplate {...args}>
    <div className="bg-red-300 h-60" />
  </GuestLoginTemplate>
)

export const Default = Template.bind({})
Default.args = {}
