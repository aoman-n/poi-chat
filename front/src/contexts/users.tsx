import { useCallback } from 'react'
import { apolloClient } from '@/lib/apolloClient'
import { AuthQuery, AuthDocument } from '@/graphql'

type OnlineUser = AuthQuery['onlineUsers'][0]

export const useUsersContext = () => {
  const data = apolloClient.readQuery<AuthQuery>({
    query: AuthDocument,
  })

  const addOnlineUser = useCallback((user: OnlineUser) => {
    const data = apolloClient.readQuery<AuthQuery>({
      query: AuthDocument,
    })

    if (!data) return

    apolloClient.writeQuery<AuthQuery>({
      query: AuthDocument,
      data: {
        ...data,
        onlineUsers: data.onlineUsers.concat(user),
      },
    })
  }, [])

  const removeOnlineUser = useCallback((id: string) => {
    const data = apolloClient.readQuery<AuthQuery>({
      query: AuthDocument,
    })

    if (!data) return

    apolloClient.writeQuery<AuthQuery>({
      query: AuthDocument,
      data: {
        ...data,
        onlineUsers: data.onlineUsers.filter((u) => u.id !== id),
      },
    })
  }, [])

  return {
    onlineUsers: data?.onlineUsers || [],
    addOnlineUser,
    removeOnlineUser,
  }
}
