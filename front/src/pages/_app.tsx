import React from 'react'
import { AppPageProps } from 'next'
import Head from 'next/head'
import { ApolloProvider } from '@apollo/client'
import { apolloClient } from '@/lib/apolloClient'
import Main from '@/components/templates/Main'
import Entrance from '@/components/templates/Entrance'
import GuestLogin from '@/components/templates/GuestLogin'
import '../styles/globals.css'

const Noop: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <>{children}</>
)

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
        <title>{pageProps.title} || poi-chat</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ApolloProvider client={apolloClient}>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </ApolloProvider>
    </>
  )
}

export default MyApp
