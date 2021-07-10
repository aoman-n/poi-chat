import React from 'react'
import Frame from '@/components/pages/RoomPage/Frame'
import ScreenSkeleton from './ScreenSkeleton'
import SettingsSkeleton from './SettingsSkeleton'
import MessagesSkeleton from './MessagesSkeleton'
import MessageFormSkeleton from './MessageFormSkeleton'

const Skeleton: React.VFC = () => {
  return (
    <Frame
      screen={<ScreenSkeleton />}
      settings={<SettingsSkeleton />}
      messages={<MessagesSkeleton />}
      messageFrom={<MessageFormSkeleton />}
    />
  )
}

export default Skeleton
