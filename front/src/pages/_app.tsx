import React from 'react'
import { AppPageProps } from 'next'
import Head from 'next/head'
import Main from '@/components/templates/Main'
import Entrance from '@/components/templates/Entrance'
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
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </>
  )
}

export default MyApp
