import React, { useCallback, useState } from 'react'
import { UserManager, User } from '@/utils/painter/user'
import Playground, { Message } from './presentation'
import {
  RoomFragment,
  useActedRoomUserEventSubscription,
  useMoveMutation,
  MessageFieldsFragment,
  useSendMessageMutation,
} from '@/graphql'
import { useCurrentUser } from '@/contexts/auth'

type PlaygroundContainerProps = {
  roomId: string
  roomDetail: RoomFragment['room']
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
          id: m.id,
          userId: m.userId,
          userName: m.userName,
          userAvatarUrl: m.userAvatarUrl,
          body: m.body,
          createdAt: m.createdAt,
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

  useActedRoomUserEventSubscription({
    variables: { roomId },
    onSubscriptionData: ({ subscriptionData }) => {
      if (!subscriptionData.data) return

      const { actedRoomUserEvent } = subscriptionData.data

      switch (actedRoomUserEvent.__typename) {
        case 'JoinedPayload': {
          // TODO: 自身は弾く
          const { roomUser } = actedRoomUserEvent
          userManager.addUser(
            new User({
              id: roomUser.id,
              avatarUrl: roomUser.avatarUrl,
              currentX: roomUser.x,
              currentY: roomUser.y,
            }),
          )
          break
        }
        case 'ExitedPayload':
          userManager.deleteUser(actedRoomUserEvent.userId)
          break
        case 'MovedPayload': {
          // TODO: 自身は弾く？
          const { roomUser } = actedRoomUserEvent
          userManager.changePos(roomUser.id, roomUser.x, roomUser.y)
          break
        }
        case 'SentMassagePayload': {
          const { lastMessage } = actedRoomUserEvent.roomUser
          if (!lastMessage) return

          handleAddMessage(lastMessage)
          break
        }
      }

      console.log({ subscriptionData })
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
