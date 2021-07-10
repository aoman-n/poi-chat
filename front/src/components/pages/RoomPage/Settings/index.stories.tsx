import React from 'react'
import { Story, Meta } from '@storybook/react'
import Settings, { SettingsProps } from '.'

export default {
  title: 'RoomPage/Settings',
  component: Settings,
} as Meta<SettingsProps>

const Template: Story<SettingsProps> = (args) => (
  <div style={{ width: '700px' }}>
    <Settings {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  handleChangeBalloonPos: () => {},
  handleRemoveBalloon: () => {},
  balloonState: {
    hasBalloon: true,
    position: 'TOP_RIGHT',
  },
}
