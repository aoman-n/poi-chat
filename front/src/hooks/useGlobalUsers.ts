import { useCallback } from 'react'
import { useApolloClient, useReactiveVar } from '@apollo/client'
import { globalUsersVar } from '@/lib/users'
import { GlobalUser } from '@/types/user'

// TODO: cacheを利用する方法を考えたい
export const useGlobalUsers = () => {
  const users = useReactiveVar(globalUsersVar)

  const addOnlineUser = useCallback(
    (user: GlobalUser) => {
      globalUsersVar([...users, user])
    },
    [users],
  )

  const removeOnlineUser = useCallback(
    (id: string) => {
      globalUsersVar(users.filter((u) => u.id !== id))
    },
    [users],
  )

  return {
    onlineUsers: users,
    addOnlineUser,
    removeOnlineUser,
  }
}
