import { AppPageProps } from 'next'
import Head from 'next/head'
import '../styles/globals.css'

function MyApp({ Component, pageProps }: AppPageProps) {
  return (
    <>
      <Head>
        <title>{pageProps.title} || poi-chat</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
