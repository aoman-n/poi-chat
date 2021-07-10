import React from 'react'
import { Story, Meta } from '@storybook/react'
import { ROOM_SCREEN_SIZE } from '@/constants'
import Skeleton from '.'
import ScreenSkeletonComponent from './ScreenSkeleton'
import SettingsSkeletonComponent from './SettingsSkeleton'
import MessagesSkeletonComponent from './MessagesSkeleton'
import MessageFormSkeletonComponent from './MessageFormSkeleton'

export default {
  title: 'RoomPage/Skeleton',
  component: Skeleton,
} as Meta

const Template: Story = () => {
  return (
    <div style={{ width: `1000px` }}>
      <Skeleton />
    </div>
  )
}
export const All = Template.bind({})

const Template2: Story = () => {
  return (
    <div
      style={{
        width: `${ROOM_SCREEN_SIZE.WIDTH}px`,
        height: `${ROOM_SCREEN_SIZE.HEIGHT}px`,
      }}
    >
      <ScreenSkeletonComponent />
    </div>
  )
}
export const ScreenSkeleton = Template2.bind({})

const Template3: Story = () => {
  return (
    <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
      <SettingsSkeletonComponent />
    </div>
  )
}
export const SettingsSkeleton = Template3.bind({})

const Template4: Story = () => {
  return (
    <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
      <MessagesSkeletonComponent />
    </div>
  )
}
export const CommentsSkeleton = Template4.bind({})

const Template5: Story = () => {
  return (
    <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
      <MessageFormSkeletonComponent />
    </div>
  )
}
export const MessageFormSkeleton = Template5.bind({})
