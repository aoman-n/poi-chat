import React from 'react'
import ContentHeader, { ContentHeaderProps } from './ContentHeader'
import RoomList, { RoomListProps } from './RoomList'
import Frame from './Frame'
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
    <>
      <Frame
        contentHeader={<ContentHeader {...contentHeaderProps} />}
        roomList={<RoomList {...roomListProps} />}
      />
      <CreateRoomModal {...createRoomModalProps} />
    </>
  )
}

export default IndexPagePresenter
