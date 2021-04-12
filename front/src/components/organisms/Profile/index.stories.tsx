import React from 'react'
import { Story, Meta } from '@storybook/react'
import Profile, { ProfileProps } from '.'

export default {
  title: 'organisms/Profile',
  component: Profile,
} as Meta

const Template: Story<ProfileProps> = (args) => (
  <div className="w-64">
    <Profile {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  name: 'sample name',
  avatarUrl:
    'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
}
