import { useCallback } from 'react'
import { useApolloClient } from '@apollo/client'

import { CommonQuery, CommonDocument } from '@/graphql'
import { GlobalUser } from '@/types/user'

export const useGlobalUsers = () => {
  const client = useApolloClient()

  const data = client.readQuery<CommonQuery>({
    query: CommonDocument,
  })

  const addOnlineUser = useCallback((user: GlobalUser) => {
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
  }, [])

  const removeOnlineUser = useCallback((id: string) => {
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
  }, [])

  return {
    onlineUsers: data?.globalUsers || [],
    addOnlineUser,
    removeOnlineUser,
  }
}
