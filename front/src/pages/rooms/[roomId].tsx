import React from 'react'
import { NextPage } from 'next'

import { AppGetServerSideProps } from '@/types'
import MainTemplate from '@/components/templates/Main'
import Header from '@/components/organisms/Header'
import OnlineUserList from '@/components/organisms/OnlineUserList'
import Profile from '@/components/organisms/Profile'
import Playground from '@/components/organisms/Playground'

import { mockOnlineUsers } from '@/mocks'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  console.log({ roomId })

  return (
    <MainTemplate
      HeaderComponent={<Header isLoggedIn={true} />}
      MainComponent={<Playground />}
      MyProfileComponent={
        <Profile
          profile={{
            name: 'sample name',
            avatarUrl:
              'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
          }}
        />
      }
      OnlineUserListComponent={<OnlineUserList users={mockOnlineUsers} />}
    />
  )
}

export const getServerSideProps: AppGetServerSideProps<{
  roomId: string | string[] | undefined
}> = async (ctx) => {
  return {
    props: {
      roomId: ctx.params?.roomId,
      title: 'ルーム',
      layout: 'Main',
    },
  }
}

export default RoomPage
