import React from 'react'
import { Story, Meta } from '@storybook/react'
import Skeleton from '.'

export default {
  title: 'RoomPage/Skeleton',
  component: Skeleton,
} as Meta

const Template: Story = (args) => <Skeleton {...args} />

export const Default = Template.bind({})
Default.args = {}
