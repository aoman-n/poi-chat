import React, { useEffect } from 'react'
import { AppPageProps } from 'next'
import Head from 'next/head'
import { ApolloProvider } from '@apollo/client'
import { createApolloClient } from '@/lib/apolloClient'
import { globalUsersVar } from '@/lib/users'
import { TotalProvider } from '@/contexts'
import { useAuthContext } from '@/contexts/auth'
import Main from '@/components/templates/Main'
import Entrance from '@/components/templates/Entrance'
import GuestLogin from '@/components/templates/GuestLogin'
import '../styles/globals.css'

import { useCommonQuery } from '@/graphql'

const WithCurrentUser: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const { setCurrentUser } = useAuthContext()
  const { data, error } = useCommonQuery()

  useEffect(() => {
    if (data) {
      setCurrentUser(data.me)
      globalUsersVar(data.globalUsers)
    }

    if (error) {
      setCurrentUser(null)
    }
  }, [data, error, setCurrentUser])

  // MEMO: unauthorizedの場合もここでerrorが入るので素通りさせる
  // if (error) return <div>error</div>

  return <>{children}</>
}

const Noop: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <>{children}</>
)

// TODO:
// main以外のLayoutではwebsocketでアクセスしないようにする
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
      <ApolloProvider client={createApolloClient()}>
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
