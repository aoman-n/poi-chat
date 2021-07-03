import React from 'react'
import RoomList, { RoomListProps } from '@/components/pages/IndexPage/RoomList'
import CreateRoomModal, {
  CreateRoomModalProps,
} from '@/components/domainParts/CreateRoomModal'

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
        <button
          className="py-2 px-6 text-white tracking-wide bg-gray-800 duration-200 hover:opacity-90 focus:outline-none text-sm rounded-sm"
          onClick={navigationProps.handleOpenModal}
        >
          ルーム作成
        </button>
      </div>
      <RoomList {...roomListProps} />
      <CreateRoomModal {...createRoomModalProps} />
    </div>
  )
}

export default IndexPagePresenter
