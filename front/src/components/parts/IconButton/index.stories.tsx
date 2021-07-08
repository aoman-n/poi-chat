import React from 'react'
import { Story, Meta } from '@storybook/react'
import Icon from '@/components/parts/Icon'
import IconButton, { IconButtonProps } from '.'

export default {
  title: 'parts/IconButton',
  component: IconButton,
} as Meta<IconButtonProps>

const Template: Story<IconButtonProps> = () => (
  <IconButton>
    <Icon type="users" />
  </IconButton>
)

export const Default = Template.bind({})
Default.args = {}
