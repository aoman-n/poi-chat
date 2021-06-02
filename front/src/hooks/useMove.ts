import { useCallback } from 'react'
import { UserManager } from '@/utils/painter'
import { useMoveMutation } from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'

export const useMove = (roomId: string, userManager: UserManager) => {
  const [moveMutation] = useMoveMutation()
  const { currentUser } = useCurrentUser()

  const handleMovePos = useCallback(
    (x: number, y: number) => {
      if (!userManager) return
      if (!currentUser) return

      moveMutation({
        variables: {
          roomId,
          x,
          y,
        },
      })
      userManager.changePos(currentUser.id, x, y)
    },
    [userManager, currentUser, moveMutation, roomId],
  )

  return { handleMovePos }
}
