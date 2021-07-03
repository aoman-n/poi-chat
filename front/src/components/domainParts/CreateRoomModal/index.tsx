import React, { useCallback, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useForm, FormProvider } from 'react-hook-form'
import { useCreateRoomMutation } from '@/graphql'
import { getRoomIdParam } from '@/utils/ids'
import { ROOM_BG_IMAGES } from '@/constants'
import { getCreateRoomErrorMsg } from './errors'
import Component, { FormData } from './presentation'

const defaultFormValues: FormData = {
  name: '',
}

export type CreateRoomModalProps = {
  open: boolean
  handleClose: () => void
}

const CreateRoomModal: React.FC<CreateRoomModalProps> = (props) => {
  const router = useRouter()
  const methods = useForm({ defaultValues: defaultFormValues })
  const [createRoom, { loading, error }] = useCreateRoomMutation({
    onCompleted: (data) => {
      methods.reset()
      router.push(`/rooms/${getRoomIdParam(data.createRoom.room.id)}`)
    },
  })

  const handleSubmit = useCallback(
    (data: FormData & { bgUrl: string }) => {
      console.log({ data })
      createRoom({ variables: { name: data.name, bgUrl: data.bgUrl } })
    },
    [createRoom],
  )

  useEffect(() => {
    methods.reset()
  }, [])

  return (
    <FormProvider {...methods}>
      <Component
        {...props}
        bgImages={ROOM_BG_IMAGES}
        handleOnSubmit={handleSubmit}
        loading={loading}
        errorMsgs={getCreateRoomErrorMsg(error)}
      />
    </FormProvider>
  )
}

export default CreateRoomModal
