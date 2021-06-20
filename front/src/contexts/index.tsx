import React from 'react'
import { SnackbarProvider } from 'notistack'
import { AuthProvider } from './auth'

export const TotalProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return (
    <AuthProvider>
      <SnackbarProvider maxSnack={3}>{children}</SnackbarProvider>
    </AuthProvider>
  )
}
