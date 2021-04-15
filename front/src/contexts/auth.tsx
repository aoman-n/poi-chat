import React, { createContext, useContext, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { AuthQuery } from '@/graphql'

type CurrentUser = AuthQuery['me']

type AuthContextValue = {
  // undefined : まだログイン確認が完了していない状態とする
  // null      : ログイン確認をした結果、ログインしていなかった状態とする
  currentUser: undefined | null | CurrentUser
  setCurrentUser: (user: CurrentUser | null) => void
}

const defaultContextValue: AuthContextValue = {
  currentUser: undefined,
  setCurrentUser: () => {},
}

export const AuthContext = createContext<AuthContextValue>(defaultContextValue)
export const useAuthContext = () => useContext(AuthContext)
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

  const values: AuthContextValue = {
    currentUser,
    setCurrentUser,
  }

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>
}
