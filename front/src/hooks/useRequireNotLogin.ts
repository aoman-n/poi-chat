import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useCurrentUser } from '@/contexts/auth'

export function useRequireNotLogin() {
  const { isAuthChecking, currentUser } = useCurrentUser()
  console.log({ isAuthChecking, currentUser })

  const router = useRouter()

  useEffect(() => {
    if (isAuthChecking) return
    if (currentUser) router.push('/') // ログインしていたのでリダイレクト
  }, [isAuthChecking, currentUser, router])
}
