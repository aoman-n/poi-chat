import React, { useState } from 'react'
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
  bgImages: [
    {
      id: '1',
      url: 'roomBg1.jpg',
    },
    {
      id: '2',
      url: 'roomBg2.png',
    },
    {
      id: '3',
      url: 'roomBg3.jpg',
    },
  ],
  handleSubmit: () => {},
}

const Template2: Story<CreateRoomModalProps> = () => {
  const [open, setOpen] = useState(false)

  const images = [
    {
      id: '1',
      url: 'roomBg1.jpg',
    },
    {
      id: '2',
      url: 'roomBg2.png',
    },
    {
      id: '3',
      url: 'roomBg3.jpg',
    },
  ]

  const handleOpen = () => {
    setOpen(true)
  }

  const handleClose = () => {
    setOpen(false)
  }

  return (
    <div>
      <button onClick={handleOpen}>open</button>
      <CraeteRooomModal
        bgImages={images}
        open={open}
        handleClose={handleClose}
        handleSubmit={() => {}}
      />
    </div>
  )
}

export const OpenToClose = Template2.bind({})
