import { useCallback } from 'react'
import { useSendMessageMutation } from '@/graphql'

export const useSendMessage = (roomId: string, callback: () => void) => {
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
      callback()
    },
    [roomId, sendMessage, callback],
  )

  return { handleSubmitMessage }
}
