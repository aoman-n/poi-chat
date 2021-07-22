import { useOnlineUsers } from '@/hooks'
import { useActedUserEventSubscription } from '@/graphql'

export const useSubscribeGlobalUserEvent = () => {
  const { onlineUsers, addOnlineUser, removeOnlineUser } = useOnlineUsers()

  useActedUserEventSubscription({
    onSubscriptionData: async ({ subscriptionData }) => {
      if (!subscriptionData.data) return
      if (!subscriptionData.data.actedUserEvent) return

      const { actedUserEvent } = subscriptionData.data

      switch (actedUserEvent.__typename) {
        case 'OnlinedPayload':
          addOnlineUser(actedUserEvent.user)
          break
        case 'OfflinedPayload':
          removeOnlineUser(actedUserEvent.user.id)
      }
    },
  })

  return { onlineUsers }
}
