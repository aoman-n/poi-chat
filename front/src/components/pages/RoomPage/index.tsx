import React from 'react'
import { filter } from 'graphql-anywhere'
import {
  useSubscribeRoomUserEvent,
  useSendMessage,
  useMovePos,
  useBalloon,
} from '@/hooks'
import { RoomPageQuery } from '@/graphql'
import { UserManager } from '@/utils/painter'
import { getErrorMsg } from './errors'
import Skeleton from './Skeleton'
import Component from './presenter'
import { ApolloError } from '@apollo/client'

export type RoomPageProps = {
  data: RoomPageQuery
  roomId: string
  userManager: UserManager
  handleMoreMessage: () => void
  moreLoading: boolean
  error: ApolloError | undefined
}

const RoomPage: React.VFC<RoomPageProps> = ({
  data,
  roomId,
  userManager,
  handleMoreMessage,
  moreLoading,
  error,
}) => {
  useSubscribeRoomUserEvent(roomId, userManager)
  const { handleMovePos } = useMovePos(roomId, userManager)
  const {
    handleChangeBalloonPos,
    handleRemoveBalloon,
    handleChangeInitialBalloonState,
    balloonState,
  } = useBalloon(userManager, roomId)

  const { handleSubmitMessage } = useSendMessage(
    roomId,
    handleChangeInitialBalloonState,
  )

  const room = data.room

  if (error) return <div>{getErrorMsg(error)}</div>
  if (!room || !userManager) return <Skeleton />

  return (
    <Component
      playgroundProps={{
        userManager: userManager,
        handleMovePos,
        bgColor: room.bgColor,
        bgUrl: room.bgUrl,
      }}
      settingsProps={{
        handleChangeBalloonPos,
        handleRemoveBalloon,
        balloonState,
      }}
      messageListProps={{
        handleMoreMessage: handleMoreMessage,
        moreLoading: moreLoading,
        messages: room.messages.nodes,
        hasMoreMessage: room.messages.pageInfo.hasPreviousPage,
      }}
      messageFormProps={{ handleSubmitMessage }}
    />
  )
}

export default RoomPage
