import React, { useEffect } from 'react'
import { AppPageProps } from 'next'
import Head from 'next/head'
import { ApolloProvider } from '@apollo/client'
import { apolloClient } from '@/lib/apolloClient'
import { TotalProvider } from '@/contexts'
import { useAuthContext } from '@/contexts/auth'
import Main from '@/components/templates/Main'
import Entrance from '@/components/templates/Entrance'
import GuestLogin from '@/components/templates/GuestLogin'
import '../styles/globals.css'

import { useAuthQuery } from '@/graphql'

const Noop: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <>{children}</>
)

const WithCurrentUser: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const { setCurrentUser, setOnlineUsers } = useAuthContext()
  const { data, error } = useAuthQuery()

  useEffect(() => {
    if (data) {
      setCurrentUser(data.me)
      setOnlineUsers(data.onlineUsers)
    }

    if (error) {
      setCurrentUser(null)
    }
  }, [data, error, setCurrentUser, setOnlineUsers])

  // TODO: エラーコンポーネントを表示する？
  // if (error) return <div>error</div>

  return <>{children}</>
}

function MyApp({ Component, pageProps }: AppPageProps) {
  const getLayout = () => {
    switch (pageProps.layout) {
      case 'Main':
        return Main
      case 'Entrance':
        return Entrance
      case 'GuestLogin':
        return GuestLogin
      default:
        return Noop
    }
  }

  const Layout = getLayout()

  return (
    <>
      <Head>
        <title>{pageProps.title} | poi-chat</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ApolloProvider client={apolloClient}>
        <TotalProvider>
          <WithCurrentUser>
            <Layout>
              <Component {...pageProps} />
            </Layout>
          </WithCurrentUser>
        </TotalProvider>
      </ApolloProvider>
    </>
  )
}

export default MyApp
