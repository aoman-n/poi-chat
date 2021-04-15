import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useCurrentUser } from '@/contexts/auth'

export function useRequireNotLogin() {
  const { isAuthChecking, currentUser } = useCurrentUser()
  const router = useRouter()

  useEffect(() => {
    if (isAuthChecking) return
    if (currentUser) router.push('/') // ログインしていたのでリダイレクト
  }, [isAuthChecking, currentUser, router])
}
