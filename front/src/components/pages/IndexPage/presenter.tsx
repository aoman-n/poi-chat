import React from 'react'
import RoomList, { RoomListProps } from '@/components/pages/IndexPage/RoomList'
import CreateRoomModal, {
  CreateRoomModalProps,
} from '@/components/domainParts/CreateRoomModal'
import Button from '@/components/parts/Button'

export type IndexPagePresenterProps = {
  navigationProps: {
    handleOpenModal: () => void
  }
  roomListProps: RoomListProps
  createRoomModalProps: CreateRoomModalProps
}

const IndexPagePresenter: React.VFC<IndexPagePresenterProps> = ({
  navigationProps,
  roomListProps,
  createRoomModalProps,
}) => {
  return (
    <div>
      <div className="flex justify-end">
        <Button onClick={navigationProps.handleOpenModal}>ルーム作成</Button>
      </div>
      <RoomList {...roomListProps} />
      <CreateRoomModal {...createRoomModalProps} />
    </div>
  )
}

export default IndexPagePresenter
