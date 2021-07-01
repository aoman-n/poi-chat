import React, { useCallback } from 'react'
import { useRouter } from 'next/router'
import { useForm, FormProvider } from 'react-hook-form'
import { useCreateRoomMutation } from '@/graphql'
import { getRoomIdParam } from '@/utils/ids'
import Component, { BgImage, FormData } from './presentation'

const images: BgImage[] = [
  {
    id: '1',
    url: 'https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg1.jpg',
  },
  {
    id: '2',
    url: 'https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg2.png',
  },
  {
    id: '3',
    url: 'https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg3.jpg',
  },
]

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

  return (
    <FormProvider {...methods}>
      <Component
        {...props}
        bgImages={images}
        handleOnSubmit={handleSubmit}
        loading={loading}
      />
    </FormProvider>
  )
}

export default CreateRoomModal
