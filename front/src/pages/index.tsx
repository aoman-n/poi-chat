import React from 'react'
import { NextPage } from 'next'
import { filter } from 'graphql-anywhere'

import { AppGetStaticProps } from '@/types'
import {
  useIndexPageQuery,
  RoomListFragment,
  RoomListFragmentDoc,
} from '@/graphql'
import RoomList from '@/components/organisms/RoomList'

const IndexRoomsPage: NextPage = () => {
  const { data } = useIndexPageQuery({ fetchPolicy: 'network-only' })

  const rooms =
    data && filter<RoomListFragment>(RoomListFragmentDoc, data).rooms.nodes

  if (!rooms) return <div>スケルトン表示</div>

  return (
    <div>
      <RoomList rooms={rooms} />
    </div>
  )
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
