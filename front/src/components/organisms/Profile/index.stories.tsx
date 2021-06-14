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
    'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png',
  // 'https://pbs.twimg.com/profile_images/1443747894/a_400x400.jpg',
}
