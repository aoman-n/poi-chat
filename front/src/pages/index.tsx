import React from 'react'
import { NextPage } from 'next'

import { AppGetStaticProps } from '@/types'
import RoomList from '@/components/organisms/RoomList'

import { mockRooms } from '@/mocks'

const IndexRoomsPage: NextPage = () => {
  return <RoomList rooms={mockRooms} />
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ルーム一覧',
      layout: 'Main',
    },
  }
}

export default IndexRoomsPage
