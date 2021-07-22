import React, { useState, useCallback } from 'react'
import { useIndexPageQuery } from '@/graphql'
import { useAuthContext } from '@/contexts/auth'
import Skeleton from './Skeleton'
import Component from './presenter'

const IndexPageContainer: React.VFC = () => {
  const { isLoggedIn } = useAuthContext()
  const [openModal, setOpenModal] = useState(false)
  const { data } = useIndexPageQuery({
    fetchPolicy: 'network-only',
    variables: { first: 10 },
  })

  const handleOpenModal = useCallback(() => {
    setOpenModal(true)
  }, [])

  const handleCloseModal = useCallback(() => {
    setOpenModal(false)
  }, [])

  if (!data) return <Skeleton />

  return (
    <Component
      contentHeaderProps={{ isLoggedIn, handleOpenModal }}
      roomListProps={{ rooms: data.rooms.nodes }}
      createRoomModalProps={{ open: openModal, handleClose: handleCloseModal }}
    />
  )
}

export default IndexPageContainer
