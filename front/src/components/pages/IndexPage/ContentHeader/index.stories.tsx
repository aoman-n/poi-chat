import React from 'react'
import { Story, Meta } from '@storybook/react'
import ContentHeader, { ContentHeaderProps } from '.'

export default {
  title: 'IndexPage/ContentHeader',
  component: ContentHeader,
} as Meta

const Template: Story<ContentHeaderProps> = (args) => (
  <div style={{ width: '600px' }}>
    <ContentHeader {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  isLoggedIn: true,
  handleOpenModal: () => {},
}
