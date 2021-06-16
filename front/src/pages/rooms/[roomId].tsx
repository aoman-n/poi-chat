import React from 'react'
import { NextPage } from 'next'
import { useRequireLogin } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import RoomPageComponent from '@/components/pages/RoomPage'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()
  return <RoomPageComponent roomId={roomId} />
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
