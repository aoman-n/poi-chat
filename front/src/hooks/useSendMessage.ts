import { useCallback } from 'react'
import { useSendMessageMutation } from '@/graphql'

export const useSendMessage = (roomId: string) => {
  const [sendMessage] = useSendMessageMutation()

  const handleSubmitMessage = useCallback(
    (values: { body: string }) => {
      if (!values.body) return

      sendMessage({
        variables: {
          roomId,
          body: values.body,
        },
      })
    },
    [roomId, sendMessage],
  )

  return { handleSubmitMessage }
}
