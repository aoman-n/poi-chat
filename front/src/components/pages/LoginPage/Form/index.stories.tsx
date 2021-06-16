import React from 'react'
import { Story, Meta } from '@storybook/react'
import LoginForm, { LoginFormProps } from '.'

export default {
  title: 'LoginPage/Form',
  component: LoginForm,
} as Meta

const Template: Story<LoginFormProps> = (args) => (
  <div className="w-96">
    <LoginForm {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {}
