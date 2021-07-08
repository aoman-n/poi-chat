import { useGlobalUsers } from '@/hooks'
import { useActedGlobalUserEventSubscription } from '@/graphql'

export const useSubscribeGlobalUserEvent = () => {
  const { onlineUsers, addOnlineUser, removeOnlineUser } = useGlobalUsers()

  useActedGlobalUserEventSubscription({
    onSubscriptionData: async ({ subscriptionData }) => {
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

  return { onlineUsers }
}
