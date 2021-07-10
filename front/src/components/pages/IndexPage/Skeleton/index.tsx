import React from 'react'
import Frame from '@/components/pages/IndexPage/Frame'
import ContentHeader from '@/components/pages/IndexPage/ContentHeader'
import RoomListSkeleton from './RoomListSkeleton'

const MessagesSkeleton: React.VFC = () => {
  return (
    <Frame
      contentHeader={
        <ContentHeader isLoggedIn={false} handleOpenModal={() => {}} />
      }
      roomList={<RoomListSkeleton />}
    />
  )
}

export default MessagesSkeleton
