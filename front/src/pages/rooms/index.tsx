import React from 'react'
import Head from 'next/head'
import { NextPage } from 'next'

import MainTemplate from '@/components/templates/Main'
import Header from '@/components/organisms/Header'
import OnlineUserList from '@/components/organisms/OnlineUserList'
import RoomList from '@/components/organisms/RoomList'
import Profile from '@/components/organisms/Profile'

import { mockOnlineUsers, mockRooms } from '@/mocks'

const RoomsPageContainer: NextPage = () => {
  return <RoomsPage />
}

const RoomsPage: React.FC = () => {
  return (
    <>
      <Head>
        <title>ルーム一覧</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainTemplate
        HeaderComponent={<Header />}
        MainComponent={<RoomList rooms={mockRooms} />}
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
    </>
  )
}

export default RoomsPageContainer
