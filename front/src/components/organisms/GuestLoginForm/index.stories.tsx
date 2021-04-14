import React from 'react'
import { Story, Meta } from '@storybook/react'
import GuestLoginForm, { GuestLoginFormProps } from '.'

export default {
  title: 'organisms/GuestLoginForm',
  component: GuestLoginForm,
} as Meta

const Template: Story<GuestLoginFormProps> = (args) => (
  <GuestLoginForm {...args} />
)

export const Default = Template.bind({})
Default.args = {}
