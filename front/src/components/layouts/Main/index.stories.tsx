import React from 'react'
import { Story, Meta } from '@storybook/react'
import MainTemplate, { MainTemplateProps } from '.'

export default {
  title: 'layouts/Main',
  component: MainTemplate,
} as Meta

const Template: Story<MainTemplateProps> = (args) => (
  <MainTemplate {...args}>
    <div className="bg-red-300 h-60" />
  </MainTemplate>
)

export const Default = Template.bind({})
Default.args = {}
