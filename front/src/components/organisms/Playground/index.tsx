import React, { useEffect, useState, useCallback } from 'react'

import { UserManager, User } from '@/painter/user'
import Playground from './presentation'

import { mockUsers, mockMessages } from '@/mocks'

const PlaygroundContainer: React.FC = () => {
  const [userManager, setUserManager] = useState<UserManager | null>(null)

  useEffect(() => {
    setUserManager(new UserManager(mockUsers.map((u) => new User(u))))
  }, [])

  const handleMovePos = useCallback(
    (x: number, y: number) => {
      if (!userManager) return

      userManager.changePos('2', x, y)
    },
    [userManager],
  )

  if (!userManager) {
    return null
  }

  return (
    <Playground
      messages={mockMessages}
      handleSubmitMessage={(e: React.FormEvent<HTMLFormElement>) =>
        e.preventDefault()
      }
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
      }}
    />
  )
}

export default PlaygroundContainer
