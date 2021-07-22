import React from 'react'
import { useSubscribeGlobalUserEvent } from './hooks'
import { useAuthContext } from '@/contexts/auth'
import withAuthcheckRequired from '@/components/domainParts/withAuthcheckRequired'
import Header from './presenter'

const HeaderContainer: React.FC = () => {
  const { currentUser } = useAuthContext()
  const { onlineUsers } = useSubscribeGlobalUserEvent()

  const props = {
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

  return <Header {...props} />
}

export default withAuthcheckRequired(HeaderContainer, () => (
  <div>display skelton</div>
))
