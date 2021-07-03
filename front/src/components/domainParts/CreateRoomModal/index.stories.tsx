import React, { useState } from 'react'
import { useForm, FormProvider } from 'react-hook-form'
import { Story, Meta } from '@storybook/react'
import { ROOM_BG_IMAGES } from '@/constants'
import CraeteRooomModal, { CreateRoomModalProps } from './presenter'

export default {
  title: 'domainParts/CreateRoomModal',
  component: CraeteRooomModal,
} as Meta

const Template: Story<CreateRoomModalProps> = (args) => {
  const methods = useForm()

  return (
    <FormProvider {...methods}>
      <CraeteRooomModal {...args} />
    </FormProvider>
  )
}

export const Default = Template.bind({})
Default.args = {
  open: true,
  handleClose: () => {},
  bgImages: ROOM_BG_IMAGES,
  handleOnSubmit: () => {},
  loading: false,
  errorMsgs: [],
}

const Template2: Story<CreateRoomModalProps> = () => {
  const methods = useForm()
  const [open, setOpen] = useState(false)

  const handleOpen = () => {
    setOpen(true)
  }

  const handleClose = () => {
    setOpen(false)
  }

  return (
    <FormProvider {...methods}>
      <button onClick={handleOpen}>open</button>
      <CraeteRooomModal
        bgImages={ROOM_BG_IMAGES}
        open={open}
        handleClose={handleClose}
        handleOnSubmit={() => {}}
        loading={false}
        errorMsgs={[]}
      />
    </FormProvider>
  )
}

export const OpenToClose = Template2.bind({})
