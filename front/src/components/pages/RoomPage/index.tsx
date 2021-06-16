import React, { useState, useCallback } from 'react'
import { filter } from 'graphql-anywhere'
import { useRequireLogin, useUserManager } from '@/hooks'
import Playground from '@/components/pages/RoomPage/Playground'
import {
  useRoomPageQuery,
  RoomFragment,
  RoomFragmentDoc,
  MoreRoomMessagesDocument,
} from '@/graphql'

const RoomPage: React.VFC<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()

  const [moreLoading, setMoreLoading] = useState(false)
  const { data, fetchMore } = useRoomPageQuery({
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

  if (!room || !userManager) return <div>スケルトン表示</div>

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
