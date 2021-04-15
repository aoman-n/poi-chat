import React from 'react'
import { AuthProvider } from './auth'

export const TotalProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return <AuthProvider>{children}</AuthProvider>
}
