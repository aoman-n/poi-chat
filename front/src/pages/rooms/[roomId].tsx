import React, { useState, useCallback } from 'react'
import { NextPage } from 'next'
import { useRoomPageQuery, MoreRoomMessagesDocument } from '@/graphql'
import { useRequireLogin, useUserManager } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import { destroyAccessPathOnServer } from '@/utils/cookies'
import RoomPageComponent from '@/components/pages/RoomPage'
import Skeleton from '@/components/pages/RoomPage/Skeleton'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
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

  if (!userManager || !data) return <Skeleton />

  return (
    <RoomPageComponent
      data={data}
      roomId={roomId}
      userManager={userManager}
      handleMoreMessage={handleMoreMessage}
      moreLoading={moreLoading}
      error={error}
    />
  )
}

export const getServerSideProps: AppGetServerSideProps<{
  roomId: string | string[] | undefined
}> = async (ctx) => {
  destroyAccessPathOnServer(ctx)

  return {
    props: {
      roomId: 'Room:' + ctx.params?.roomId,
      title: 'ルーム',
      layout: 'Main',
    },
  }
}

export default RoomPage
