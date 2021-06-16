import React from 'react'
import { Story, Meta } from '@storybook/react'
import Header, { HeaderProps } from './presentation'

export default {
  title: 'domainParts/Header',
  component: Header,
} as Meta

const Template: Story<HeaderProps> = (args) => (
  <div className="h-16 bg-white">
    <Header {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  isLoggedIn: true,
}
