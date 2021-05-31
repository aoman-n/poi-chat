import React, { useEffect, useState, useRef } from 'react'
import { NextPage } from 'next'
import { filter } from 'graphql-anywhere'

import { useRequireLogin } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import Playground from '@/components/organisms/Playground'
import { useRoomPageQuery, RoomFragment, RoomFragmentDoc } from '@/graphql'
import { UserManager, User } from '@/utils/painter/user'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()
  const isCreatedUserManager = useRef<boolean>(false)
  const [userManager, setUserManager] = useState<UserManager | null>(null)
  const { data } = useRoomPageQuery({ variables: { roomId } })

  const roomDetail =
    (data && filter<RoomFragment>(RoomFragmentDoc, data).room) || null

  useEffect(() => {
    if (isCreatedUserManager.current) return

    if (roomDetail) {
      const { users } = roomDetail
      const userInstances = users.map(
        (user) =>
          new User({
            id: user.id,
            avatarUrl: user.avatarUrl,
            currentX: user.x,
            currentY: user.y,
          }),
      )
      setUserManager(new UserManager(userInstances))
      isCreatedUserManager.current = true
    }
  }, [roomDetail])

  // roomDeailtがnull & userManager未作成 のときはスケルトンを表示
  if (!roomDetail || !userManager) return <div>スケルトン表示</div>

  return (
    <Playground
      roomDetail={roomDetail}
      roomId={roomId}
      userManager={userManager}
    />
  )
}

export const getServerSideProps: AppGetServerSideProps<{
  roomId: string | string[] | undefined
}> = async (ctx) => {
  return {
    props: {
      roomId: 'Room:' + ctx.params?.roomId,
      title: 'ルーム',
      layout: 'Main',
    },
  }
}

export default RoomPage
