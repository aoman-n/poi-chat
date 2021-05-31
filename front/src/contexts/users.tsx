import { useCallback } from 'react'
import { apolloClient } from '@/lib/apolloClient'
import { CommonQuery, CommonDocument } from '@/graphql'

type OnlineUser = CommonQuery['globalUsers'][0]

export const useUsersContext = () => {
  const data = apolloClient.readQuery<CommonQuery>({
    query: CommonDocument,
  })

  const addOnlineUser = useCallback((user: OnlineUser) => {
    const data = apolloClient.readQuery<CommonQuery>({
      query: CommonDocument,
    })

    if (!data) return

    apolloClient.writeQuery<CommonQuery>({
      query: CommonDocument,
      data: {
        ...data,
        globalUsers: data.globalUsers.concat(user),
      },
    })
  }, [])

  const removeOnlineUser = useCallback((id: string) => {
    const data = apolloClient.readQuery<CommonQuery>({
      query: CommonDocument,
    })

    if (!data) return

    apolloClient.writeQuery<CommonQuery>({
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
