import React from 'react'
import { Story, Meta } from '@storybook/react'
import EntranceTemplate, { EntranceProps } from '.'

export default {
  title: 'layouts/Entrance',
  component: EntranceTemplate,
} as Meta

const Template: Story<EntranceProps> = (args) => (
  <EntranceTemplate {...args}>
    <div className="bg-red-300 h-60" />
  </EntranceTemplate>
)

export const Default = Template.bind({})
Default.args = {}
