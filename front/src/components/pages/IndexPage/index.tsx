import React from 'react'
import { filter } from 'graphql-anywhere'
import {
  useIndexPageQuery,
  RoomListFragment,
  RoomListFragmentDoc,
} from '@/graphql'
import RoomList from '@/components/pages/IndexPage/RoomList'

const IndexPage: React.VFC = () => {
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

export default IndexPage
