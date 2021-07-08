import { useCallback } from 'react'
import { UserManager } from '@/utils/painter'
import { useMoveMutation } from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'
import { useThrottleFn } from '@/hooks'

export const useMovePos = (roomId: string, userManager: UserManager) => {
  const [moveMutation] = useMoveMutation()
  const { currentUser } = useCurrentUser()

  const handleMovePos = useCallback(
    (e: MouseEvent) => {
      if (!userManager) return
      if (!currentUser) return

      const x = e.offsetX
      const y = e.offsetY

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

  return {
    handleMovePos: useThrottleFn(handleMovePos, 400),
  }
}
