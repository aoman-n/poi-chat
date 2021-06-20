import React, { useEffect } from 'react'
import { AppPageProps } from 'next'
import Head from 'next/head'
import { ApolloProvider } from '@apollo/client'
import { createApolloClient } from '@/utils/apolloClient'
import { TotalProvider } from '@/contexts'
import { useAuthContext } from '@/contexts/auth'
import Main from '@/components/layouts/Main'
import Entrance from '@/components/layouts/Entrance'
import GuestLogin from '@/components/layouts/GuestLogin'
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
    }

    if (error) {
      setCurrentUser(null)
    }
  }, [data, error, setCurrentUser])

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
