import React from 'react'
import { Story, Meta } from '@storybook/react'
import CropIcon, { CropIconProps } from '.'

export default {
  title: 'organisms/CropIcon',
  component: CropIcon,
  argTypes: {
    handleUpdateImage: { action: 'updateImage!' },
  },
} as Meta

const Template: Story<CropIconProps> = (args) => <CropIcon {...args} />

export const Default = Template.bind({})
Default.args = {
  imageUrl: 'http://placekitten.com/500/800',
  height: 400,
  width: 400,
}
