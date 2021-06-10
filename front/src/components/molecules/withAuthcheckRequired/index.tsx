import React from 'react'
import { useAuthContext } from '@/contexts/auth'

const withAuthcheckRequired = <P extends Record<string, unknown>>(
  Component: React.ComponentType<P>,
  LoadingComponent: React.ComponentType,
): React.FC<P> => {
  return function WithAuthcheckRequired(props: P): JSX.Element {
    const { isAuthChecking } = useAuthContext()

    if (isAuthChecking) return <LoadingComponent />

    return <Component {...props} />
  }
}

export default withAuthcheckRequired
