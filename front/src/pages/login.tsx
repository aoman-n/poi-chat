import Head from 'next/head'
import Entrance from '@/components/templates/Entrance'

const Login = () => {
  return (
    <>
      <Head>
        <title>ログイン</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Entrance />
    </>
  )
}

export default Login
