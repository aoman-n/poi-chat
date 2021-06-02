import React from 'react'
import { NextPage } from 'next'
import { filter } from 'graphql-anywhere'
import { useRequireLogin, useUserManager } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import Playground from '@/components/organisms/Playground'
import { useRoomPageQuery, RoomFragment, RoomFragmentDoc } from '@/graphql'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()

  const { data } = useRoomPageQuery({ variables: { roomId } })
  const { userManager } = useUserManager(data?.room.users)

  const room =
    (data && filter<RoomFragment>(RoomFragmentDoc, data).room) || null

  if (!room || !userManager) return <div>スケルトン表示</div>

  return <Playground room={room} roomId={roomId} userManager={userManager} />
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
