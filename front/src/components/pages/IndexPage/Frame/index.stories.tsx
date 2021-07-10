import React from 'react'
import { Story, Meta } from '@storybook/react'
import Frame from '.'

export default {
  title: 'IndexPage/Frame',
  component: Frame,
} as Meta

const Template: Story = () => {
  return (
    <Frame
      contentHeader={<div className="bg-blue-300 h-24">screenArea</div>}
      roomList={<div className="bg-yellow-300 h-80">settingsArea</div>}
    />
  )
}

export const Default = Template.bind({})
