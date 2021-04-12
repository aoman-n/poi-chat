import React from 'react'
import { NextPage } from 'next'

import { AppGetStaticProps } from '@/types'
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
    <MainTemplate
      HeaderComponent={<Header isLoggedIn={true} />}
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
  )
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ルーム一覧',
      layout: 'Main',
    },
  }
}

export default RoomsPageContainer
