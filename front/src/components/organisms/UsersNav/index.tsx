import React from 'react'
import { useAuthContext } from '@/contexts/auth'
import { useUsersContext } from '@/contexts/users'
import UsersNav from './presentation'
import {
  useKeeoOnlineSubscription,
  useChangedUserStatusSubscription,
} from '@/graphql'

const UsersNavContainer: React.FC = () => {
  const { currentUser, isAuthChecking } = useAuthContext()
  const { onlineUsers, addOnlineUser, removeOnlineUser } = useUsersContext()
  useKeeoOnlineSubscription({
    onSubscriptionComplete: () => {
      console.log('complete keep online connected')
    },
  })
  useChangedUserStatusSubscription({
    onSubscriptionData: ({ client, subscriptionData }) => {
      if (!subscriptionData.data) return
      if (!subscriptionData.data.changedUserStatus) return

      const { changedUserStatus } = subscriptionData.data

      if (changedUserStatus.__typename === 'OnlineUserStatus') {
        addOnlineUser(changedUserStatus)
      } else if (changedUserStatus.__typename === 'OfflineUserStatus') {
        removeOnlineUser(changedUserStatus.id)
      }
    },
  })

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
