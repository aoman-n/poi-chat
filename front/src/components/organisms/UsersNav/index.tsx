import React from 'react'
import { useAuthContext } from '@/contexts/auth'
import UsersNav from './presentation'

const UsersNavContainer: React.FC = () => {
  const { currentUser, isAuthChecking, onlineUsers } = useAuthContext()

  const props = {
    profile: currentUser
      ? {
          name: currentUser.displayName,
          avatarUrl: currentUser.avatarUrl,
        }
      : null,
    onlineUserList: {
      users: onlineUsers,
    },
  }

  if (isAuthChecking) {
    return <div>スケルトン表示</div>
  }

  return <UsersNav {...props} />
}

export default UsersNavContainer
