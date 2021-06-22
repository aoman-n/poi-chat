import React from 'react'
import { Story, Meta } from '@storybook/react'
import CraeteRooomModal, { CreateRoomModalProps } from './presentation'

export default {
  title: 'domainParts/CreateRoomModal',
  component: CraeteRooomModal,
} as Meta

const Template: Story<CreateRoomModalProps> = (args) => (
  <div>
    <CraeteRooomModal {...args} />
  </div>
)

export const Default = Template.bind({})
Default.args = {
  open: true,
  handleClose: () => {},
}
