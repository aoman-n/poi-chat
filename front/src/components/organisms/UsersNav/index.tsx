import React from 'react'
import { useAuthContext } from '@/contexts/auth'
import { useUsersContext } from '@/contexts/users'
import UsersNav from './presentation'
import { useActedGlobalUserEventSubscription } from '@/graphql'

const UsersNavContainer: React.FC = () => {
  const { currentUser, isAuthChecking } = useAuthContext()
  const { onlineUsers, addOnlineUser, removeOnlineUser } = useUsersContext()

  useActedGlobalUserEventSubscription({
    onSubscriptionComplete: () => {
      console.log('start subscribe global user events')
    },
    onSubscriptionData: ({ client, subscriptionData }) => {
      if (!subscriptionData.data) return
      if (!subscriptionData.data.actedGlobalUserEvent) return

      const { actedGlobalUserEvent } = subscriptionData.data

      switch (actedGlobalUserEvent.__typename) {
        case 'OnlinedPayload':
          addOnlineUser(actedGlobalUserEvent.globalUser)
          break
        case 'OfflinedPayload':
          removeOnlineUser(actedGlobalUserEvent.userId)
      }
    },
  })

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

  if (isAuthChecking) {
    return <div>スケルトン表示</div>
  }

  return <UsersNav {...props} />
}

export default UsersNavContainer
