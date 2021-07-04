import React from 'react'
import ContentHeader, { ContentHeaderProps } from './ContentHeader'
import RoomList, { RoomListProps } from './RoomList'
import CreateRoomModal, {
  CreateRoomModalProps,
} from '@/components/domainParts/CreateRoomModal'

export type IndexPagePresenterProps = {
  contentHeaderProps: ContentHeaderProps
  roomListProps: RoomListProps
  createRoomModalProps: CreateRoomModalProps
}

const IndexPagePresenter: React.VFC<IndexPagePresenterProps> = ({
  contentHeaderProps,
  roomListProps,
  createRoomModalProps,
}) => {
  return (
    <div>
      <ContentHeader {...contentHeaderProps} />
      <RoomList {...roomListProps} />
      <CreateRoomModal {...createRoomModalProps} />
    </div>
  )
}

export default IndexPagePresenter
