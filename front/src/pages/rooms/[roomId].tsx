import { NextPage, GetServerSideProps } from 'next'
import Head from 'next/head'

import MainTemplate from '@/components/templates/Main'
import Header from '@/components/organisms/Header'
import OnlineUserList from '@/components/organisms/OnlineUserList'
import Profile from '@/components/organisms/Profile'
import Playground from '@/components/organisms/Playground'

import { mockOnlineUsers, mockMessages } from '@/mocks'
import React from 'react'

const RoomPage: NextPage<{ roomId: string }> = ({ roomId }) => {
  console.log({ roomId })

  return (
    <>
      <Head>
        <title>ルーム一覧</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainTemplate
        HeaderComponent={<Header />}
        MainComponent={
          <Playground
            messages={mockMessages}
            handleSubmitMessage={(e: React.FormEvent<HTMLFormElement>) =>
              e.preventDefault()
            }
          />
        }
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

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  return {
    props: { roomId: ctx.params?.roomId },
  }
}

export default RoomPage
