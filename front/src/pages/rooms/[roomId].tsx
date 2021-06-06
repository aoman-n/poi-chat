import React, { useState, useCallback } from 'react'
import { NextPage } from 'next'
import { filter } from 'graphql-anywhere'
import { useRequireLogin, useUserManager } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import Playground from '@/components/organisms/Playground'
import {
  useRoomPageQuery,
  RoomFragment,
  RoomFragmentDoc,
  MoreRoomMessagesDocument,
} from '@/graphql'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()

  const [moreLoading, setMoreLoading] = useState(false)
  const { data, fetchMore, loading, networkStatus } = useRoomPageQuery({
    variables: { roomId },
    notifyOnNetworkStatusChange: true,
  })
  const { userManager } = useUserManager(data?.room.users)

  console.log({ loading, networkStatus, moreLoading })

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

export const getServerSideProps: AppGetServerSideProps<{
  roomId: string | string[] | undefined
}> = async (ctx) => {
  return {
    props: {
      roomId: 'Room:' + ctx.params?.roomId,
      title: 'ルーム',
      layout: 'Main',
    },
  }
}

export default RoomPage
