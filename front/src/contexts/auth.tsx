import React, { createContext, useCallback, useContext, useState } from 'react'
import { AuthQuery } from '@/graphql'

type CurrentUser = AuthQuery['me']
type OnlineUser = AuthQuery['onlineUsers'][0]

type AuthContextValue = {
  // undefined : まだログイン確認が完了していない状態とする
  // null      : ログイン確認をした結果、ログインしていなかった状態とする
  currentUser: undefined | null | CurrentUser
  setCurrentUser: (user: CurrentUser | null) => void
  onlineUsers: OnlineUser[]
  setOnlineUsers: (users: OnlineUser[]) => void
  addOnlineUser: (user: OnlineUser) => void
  removeOnlineUser: (id: string) => void
}

const defaultContextValue: AuthContextValue = {
  currentUser: undefined,
  setCurrentUser: () => {},
  onlineUsers: [],
  setOnlineUsers: () => {},
  addOnlineUser: () => {},
  removeOnlineUser: () => {},
}

export const AuthContext = createContext<AuthContextValue>(defaultContextValue)
export const useAuthContext = () => {
  const { currentUser, ...props } = useContext(AuthContext)
  const isAuthChecking = currentUser === undefined
  const isLoggedIn = !!currentUser

  return {
    currentUser,
    isAuthChecking,
    isLoggedIn,
    ...props,
  }
}
export const useCurrentUser = () => {
  const { currentUser } = useContext(AuthContext)
  const isAuthChecking = currentUser === undefined

  return {
    currentUser,
    isAuthChecking,
  }
}

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [currentUser, setCurrentUser] = useState<
    undefined | null | CurrentUser
  >(undefined)
  const [onlineUsers, setOnlineUsers] = useState<OnlineUser[]>([])

  const handleAddOnlineUser = useCallback((user: OnlineUser) => {
    setOnlineUsers((prev) => ({
      ...prev,
      user,
    }))
  }, [])

  const handleRemoveOnlineUser = useCallback((id: string) => {
    setOnlineUsers((prev) => {
      return prev.filter((u) => u.id !== id)
    })
  }, [])

  const values: AuthContextValue = {
    currentUser,
    setCurrentUser,
    onlineUsers,
    setOnlineUsers,
    addOnlineUser: handleAddOnlineUser,
    removeOnlineUser: handleRemoveOnlineUser,
  }

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>
}
