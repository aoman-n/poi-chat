import { useCallback } from 'react'
import { produce } from 'immer'
import { useApolloClient } from '@apollo/client'
import { CommonQuery, CommonDocument, useCommonQuery, User } from '@/graphql'

export const useOnlineUsers = () => {
  const client = useApolloClient()

  const { data } = useCommonQuery({
    fetchPolicy: 'cache-only',
  })

  const addOnlineUser = useCallback(
    (user: User) => {
      const data = client.readQuery<CommonQuery>({
        query: CommonDocument,
      })

      if (!data) return

      const newData = produce(data, (draft) => {
        draft.onlineUsers.push(user)
      })

      client.writeQuery<CommonQuery>({
        query: CommonDocument,
        data: newData,
      })
    },
    [client],
  )

  const removeOnlineUser = useCallback(
    (id: string) => {
      const data = client.readQuery<CommonQuery>({
        query: CommonDocument,
      })

      if (!data) return

      client.writeQuery<CommonQuery>({
        query: CommonDocument,
        data: {
          ...data,
          onlineUsers: data.onlineUsers.filter((u) => u.id !== id),
        },
      })
    },
    [client],
  )

  return {
    onlineUsers: data?.onlineUsers || [],
    addOnlineUser,
    removeOnlineUser,
  }
}
