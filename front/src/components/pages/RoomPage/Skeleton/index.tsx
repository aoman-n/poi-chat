import React from 'react'
import Frame from '@/components/pages/RoomPage/Frame'
import ScreenSkeleton from './ScreenSkeleton'
import SettingsSkeleton from './SettingsSkeleton'
import CommentsSkeleton from './CommentsSkeleton'

const Skeleton: React.VFC = () => {
  return (
    <Frame
      screen={<ScreenSkeleton />}
      settings={<SettingsSkeleton />}
      comments={<CommentsSkeleton />}
    />
  )
}

export default Skeleton
