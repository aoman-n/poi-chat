import React, { useCallback, useState } from 'react'
import { UserManager, User } from '@/painter/user'
import Playground, { Message } from './presentation'
import {
  RoomDetailFragment,
  useUserEventSubscription,
  useJoinRoomSubscription,
  useMoveMutation,
  useMessageSubscription,
  MessageFieldsFragment,
  useSendMessageMutation,
} from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'

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
  const [sendMessage] = useSendMessageMutation()
  const fetchedMessages = roomDetail.messages.nodes
  const [messages, setMessages] = useState<Message[]>(
    fetchedMessages
      .map((m) => {
        return {
          id: m?.id || '',
          userId: m?.userId || '',
          userName: m?.userName || '',
          userAvatarUrl: m?.userAvatarUrl || '',
          body: m?.body || '',
          createdAt: m?.createdAt || '',
        }
      })
      .reverse(),
  )
  const [moveMutation] = useMoveMutation()
  const { currentUser } = useCurrentUser()

  const handleAddMessage = useCallback(
    (message: MessageFieldsFragment) => {
      setMessages((prev) => [
        ...prev,
        {
          id: message.id,
          userId: message.userId,
          userName: message.userName,
          userAvatarUrl: message.userAvatarUrl,
          body: message.body,
          createdAt: message.createdAt,
        },
      ])
    },
    [setMessages],
  )

  useMessageSubscription({
    variables: { roomId },
    onSubscriptionData: ({ subscriptionData }) => {
      if (!subscriptionData.data) return
      handleAddMessage(subscriptionData.data.subMessage)
    },
  })

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

  const handleSubmitMessage = useCallback(
    (values: { body: string }) => {
      if (!values.body) return

      sendMessage({
        variables: {
          roomId,
          body: values.body,
        },
      })
    },
    [roomId, sendMessage],
  )

  if (!userManager) {
    return null
  }

  return (
    <Playground
      messages={messages}
      handleSubmitMessage={handleSubmitMessage}
      rooomScreenProps={{
        userManager: userManager,
        handleMovePos: handleMovePos,
      }}
    />
  )
}

export default PlaygroundContainer
