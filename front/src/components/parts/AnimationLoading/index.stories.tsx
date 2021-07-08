import React from 'react'
import { Story, Meta } from '@storybook/react'
import AnimationLoading from '.'

export default {
  title: 'parts/AnimationLoading',
  component: AnimationLoading,
} as Meta

const Template: Story = (args) => <AnimationLoading {...args} />

export const Default = Template.bind({})
Default.args = {}
