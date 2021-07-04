import React from 'react'
import { Story, Meta } from '@storybook/react'
import CircleArrowButton, { CircleArrowButtonProps } from '.'

export default {
  title: 'parts/CircleArrowButton',
  component: CircleArrowButton,
} as Meta<CircleArrowButtonProps>

const Template: Story<CircleArrowButtonProps> = (args) => (
  <CircleArrowButton {...args} />
)

export const Left = Template.bind({})
Left.args = {
  arrowType: 'left',
  onClick: () => {},
}

export const Right = Template.bind({})
Right.args = {
  arrowType: 'right',
  onClick: () => {},
}
