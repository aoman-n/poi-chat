import React from 'react'
import { NextPage } from 'next'
import { filter } from 'graphql-anywhere'
import { useRequireLogin, useUserManager } from '@/hooks'
import { AppGetServerSideProps } from '@/types'
import Playground from '@/components/organisms/Playground'
import {
  useRoomPageQuery,
  RoomFragment,
  RoomFragmentDoc,
  OnlyMessagesDocument,
} from '@/graphql'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  useRequireLogin()

  const { data, fetchMore } = useRoomPageQuery({
    variables: { roomId },
  })
  const { userManager } = useUserManager(data?.room.users)

  const room =
    (data && filter<RoomFragment>(RoomFragmentDoc, data).room) || null

  if (!room || !userManager) return <div>スケルトン表示</div>

  return (
    <div>
      <div>
        {room.messages.pageInfo.hasPreviousPage && (
          <button
            onClick={() => {
              fetchMore({
                query: OnlyMessagesDocument,
                variables: {
                  roomId,
                  before: room.messages.pageInfo.startCursor,
                },
              })
            }}
          >
            more
          </button>
        )}
      </div>
      <div>
        <Playground room={room} roomId={roomId} userManager={userManager} />
      </div>
    </div>
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
