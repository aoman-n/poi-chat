import React from 'react'
import { mockOnlineUsers } from '@/mocks'
import UsersNav from './presentation'

const UsersNavContainer: React.FC = () => {
  const props = {
    profile: {
      name: 'sample name',
      avatarUrl:
        'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
    },
    onlineUserList: {
      users: mockOnlineUsers,
    },
  }

  return <UsersNav {...props} />
}

export default UsersNavContainer
