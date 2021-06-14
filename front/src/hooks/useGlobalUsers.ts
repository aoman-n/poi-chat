import { useCallback } from 'react'
import { useApolloClient } from '@apollo/client'

import { CommonQuery, CommonDocument, useCommonQuery } from '@/graphql'
import { GlobalUser } from '@/types/user'

export const useGlobalUsers = () => {
  const client = useApolloClient()

  const { data } = useCommonQuery({
    fetchPolicy: 'cache-only',
  })

  const addOnlineUser = useCallback(
    (user: GlobalUser) => {
      const data = client.readQuery<CommonQuery>({
        query: CommonDocument,
      })

      if (!data) return

      client.writeQuery<CommonQuery>({
        query: CommonDocument,
        data: {
          ...data,
          globalUsers: data.globalUsers.concat(user),
        },
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
          globalUsers: data.globalUsers.filter((u) => u.id !== id),
        },
      })
    },
    [client],
  )

  return {
    onlineUsers: data?.globalUsers || [],
    addOnlineUser,
    removeOnlineUser,
  }
}
