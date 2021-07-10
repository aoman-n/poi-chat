import React from 'react'
import { Story, Meta } from '@storybook/react'
import Frame from '.'

export default {
  title: 'RoomPage/Frame',
  component: Frame,
} as Meta

const Template: Story = () => {
  return (
    <div style={{ width: `1000px` }}>
      <Frame
        screen={<div className="bg-blue-300 h-full">screenArea</div>}
        settings={<div className="bg-yellow-300 h-20">settingsArea</div>}
        messages={<div className="bg-green-300 h-80">messagesArea</div>}
        messageFrom={<div className="bg-pink-300 h-14">messageFormArea</div>}
      />
    </div>
  )
}

export const Default = Template.bind({})
