import { useEffect, useRef, useState } from 'react'
import { RoomFragment } from '@/graphql'
import { UserManager } from '@/utils/painter'

export const useUserManager = (
  users: RoomFragment['room']['users'] | undefined,
) => {
  const isCreatedUserManager = useRef<boolean>(false)
  const [userManager, setUserManager] = useState<UserManager | null>(null)

  useEffect(() => {
    if (isCreatedUserManager.current) return

    if (users) {
      setUserManager(new UserManager(users))
      isCreatedUserManager.current = true
    }
  }, [users])

  return { userManager }
}
