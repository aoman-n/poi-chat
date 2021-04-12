import React from 'react'
import { NextPage } from 'next'

import { AppGetServerSideProps } from '@/types'
import Playground from '@/components/organisms/Playground'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  console.log({ roomId })

  return <Playground />
}

export const getServerSideProps: AppGetServerSideProps<{
  roomId: string | string[] | undefined
}> = async (ctx) => {
  return {
    props: {
      roomId: ctx.params?.roomId,
      title: 'ルーム',
      layout: 'Main',
    },
  }
}

export default RoomPage
