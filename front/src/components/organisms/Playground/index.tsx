import React, { useEffect, useState } from 'react'

import { UserManager, User } from '@/painter/user'
import Playground from './presentation'

import { mockUsers, mockMessages } from '@/mocks'

const PlaygroundContainer: React.FC = () => {
  const [userManager, setUserManager] = useState<UserManager | null>(null)

  useEffect(() => {
    console.log('useEffect!')
    setUserManager(new UserManager(mockUsers.map((u) => new User(u))))
  }, [])

  if (!userManager) {
    return null
  }

  return (
    <Playground
      messages={mockMessages}
      handleSubmitMessage={(e: React.FormEvent<HTMLFormElement>) =>
        e.preventDefault()
      }
      userManager={userManager}
    />
  )
}

export default PlaygroundContainer
