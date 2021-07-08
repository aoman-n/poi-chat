import React from 'react'
import { useAuthContext } from '@/contexts/auth'
import withAuthcheckRequired from '@/components/domainParts/withAuthcheckRequired'
import { useSubscribeGlobalUserEvent } from '../Header/hooks'
import UsersNav from './presenter'

const UsersNavContainer: React.FC = () => {
  const { currentUser, isLoggedIn } = useAuthContext()
  const { onlineUsers } = useSubscribeGlobalUserEvent()

  const props = {
    isLoggedIn,
    profile: currentUser
      ? {
          name: currentUser.name,
          avatarUrl: currentUser.avatarUrl,
        }
      : null,
    onlineUserList: {
      users: onlineUsers,
    },
  }

  return <UsersNav {...props} />
}

export default withAuthcheckRequired(UsersNavContainer, () => (
  <div>display skelton</div>
))
