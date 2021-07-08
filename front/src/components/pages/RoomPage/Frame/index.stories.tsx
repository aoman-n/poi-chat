import React from 'react'
import { Story, Meta } from '@storybook/react'
import { ROOM_SCREEN_SIZE } from '@/constants'
import Frame from '.'

export default {
  title: 'RoomPage/Frame',
  component: Frame,
} as Meta

const Template: Story = () => {
  return (
    <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
      <Frame
        screen={<div className="bg-blue-300">screenArea</div>}
        settings={<div className="bg-yellow-300 h-20">settingsArea</div>}
        comments={<div className="bg-green-300 h-32">commentsArea</div>}
      />
    </div>
  )
}

export const Default = Template.bind({})
