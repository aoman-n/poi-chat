import React, { useState, useCallback } from 'react'
import { filter } from 'graphql-anywhere'
import { useRequireLogin, useUserManager } from '@/hooks'
import {
  useRoomPageQuery,
  RoomFragment,
  RoomFragmentDoc,
  MoreRoomMessagesDocument,
} from '@/graphql'
import { getErrorMsg } from './errors'
import Skeleton from './Skeleton'
import Playground from './Playground'

const RoomPage: React.VFC<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()

  const [moreLoading, setMoreLoading] = useState(false)
  const { data, fetchMore, error } = useRoomPageQuery({
    variables: { roomId },
    notifyOnNetworkStatusChange: true,
  })
  const { userManager } = useUserManager(data?.room.users)

  const handleMoreMessage = useCallback(async () => {
    setMoreLoading(true)
    await fetchMore({
      query: MoreRoomMessagesDocument,
      variables: {
        roomId,
        before: data?.room.messages.pageInfo.startCursor,
      },
    })
    setMoreLoading(false)
  }, [fetchMore, data, roomId])

  const room =
    (data && filter<RoomFragment>(RoomFragmentDoc, data).room) || null

  if (error) return <div>{getErrorMsg(error)}</div>
  if (!room || !userManager) return <Skeleton />

  return (
    <Playground
      room={room}
      roomId={roomId}
      userManager={userManager}
      handleMoreMessage={handleMoreMessage}
      moreLoading={moreLoading}
    />
  )
}

export default RoomPage
