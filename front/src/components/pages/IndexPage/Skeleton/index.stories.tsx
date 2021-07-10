import React from 'react'
import { Story, Meta } from '@storybook/react'
import Skeleton from '.'
import RoomListSkeletonComponent from './RoomListSkeleton'

export default {
  title: 'IndexPage/Skeleton',
  component: Skeleton,
} as Meta

const Template: Story = () => {
  return <Skeleton />
}
export const All = Template.bind({})

const ContentHeaderSkeletonTmp: Story = () => {
  return (
    <div style={{ width: '600px' }}>
      <RoomListSkeletonComponent />
    </div>
  )
}
export const ContentHeaderSkeleton = ContentHeaderSkeletonTmp.bind({})
