import React, { useCallback } from 'react'
import { UserManager, User } from '@/painter/user'
import Playground from './presentation'
import {
  RoomDetailFragment,
  useUserEventSubscription,
  useJoinRoomSubscription,
  useMoveMutation,
} from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'
import { mockMessages } from '@/mocks'

type PlaygroundContainerProps = {
  roomId: string
  roomDetail: RoomDetailFragment['roomDetail']
  userManager: UserManager
}

const PlaygroundContainer: React.FC<PlaygroundContainerProps> = ({
  roomId,
  roomDetail,
  userManager,
}) => {
  const [moveMutation] = useMoveMutation()
  const { currentUser } = useCurrentUser()

  useJoinRoomSubscription({ variables: { roomId } })
  useUserEventSubscription({
    variables: { roomId },
    onSubscriptionData: ({ subscriptionData }) => {
      if (!subscriptionData.data) return

      const { subUserEvent } = subscriptionData.data

      switch (subUserEvent.__typename) {
        case 'JoinedUser':
          userManager.addUser(
            new User({
              id: subUserEvent.id,
              avatarUrl: subUserEvent.avatarUrl,
              currentX: subUserEvent.x,
              currentY: subUserEvent.y,
            }),
          )
          break
        case 'MovedUser':
          if (currentUser && subUserEvent.id === currentUser.id) break
          userManager.changePos(subUserEvent.id, subUserEvent.x, subUserEvent.y)
          break
        case 'ExitedUser':
          userManager.deleteUser(subUserEvent.id)
          break
      }
    },
  })

  const handleMovePos = useCallback(
    (x: number, y: number) => {
      if (!userManager) return
      if (!currentUser) return

      moveMutation({
        variables: {
          roomId,
          x,
          y,
        },
      })
      userManager.changePos(currentUser.id, x, y)
    },
    [userManager, currentUser, moveMutation, roomId],
  )

  if (!userManager) {
    return null
  }

  return (
    <Playground
      messages={mockMessages}
      handleSubmitMessage={(e: React.FormEvent<HTMLFormElement>) =>
        e.preventDefault()
      }
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
      }}
    />
  )
}

export default PlaygroundContainer
