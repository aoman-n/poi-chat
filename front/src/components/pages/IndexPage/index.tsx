import React, { useState, useCallback } from 'react'
import { filter } from 'graphql-anywhere'
import {
  useIndexPageQuery,
  RoomListFragment,
  RoomListFragmentDoc,
} from '@/graphql'
import Component, { IndexPagePresenterProps } from './presenter'

const IndexPageContainer: React.VFC = () => {
  const [openModal, setOpenModal] = useState(false)
  const { data } = useIndexPageQuery({ fetchPolicy: 'network-only' })

  const handleOpenModal = useCallback(() => {
    setOpenModal(true)
  }, [])

  const handleCloseModal = useCallback(() => {
    setOpenModal(false)
  }, [])

  const rooms =
    data && filter<RoomListFragment>(RoomListFragmentDoc, data).rooms.nodes

  if (!rooms) return <div>スケルトン表示</div>

  const passProps: IndexPagePresenterProps = {
    navigationProps: { handleOpenModal },
    roomListProps: { rooms },
    createRoomModalProps: {
      open: openModal,
      handleClose: handleCloseModal,
    },
  }

  return <Component {...passProps} />
}

export default IndexPageContainer
