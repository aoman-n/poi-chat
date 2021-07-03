import React from 'react'
import { useAuthContext } from '@/contexts/auth'
import { useGlobalUsers } from '@/hooks'
import UsersNav from './presenter'
import { useActedGlobalUserEventSubscription } from '@/graphql'
import withAuthcheckRequired from '@/components/domainParts/withAuthcheckRequired'

const UsersNavContainer: React.FC = () => {
  const { currentUser, isAuthChecking } = useAuthContext()
  const { onlineUsers, addOnlineUser, removeOnlineUser } = useGlobalUsers()

  useActedGlobalUserEventSubscription({
    onSubscriptionData: async ({ subscriptionData }) => {
      console.log({ globalEventData: subscriptionData })

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

export default withAuthcheckRequired(UsersNavContainer, () => (
  <div>checking!!!</div>
))
