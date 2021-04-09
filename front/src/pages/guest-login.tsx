import Head from 'next/head'
import GuestLogin from '@/components/templates/GuestLogin'

const Login = () => {
  return (
    <>
      <Head>
        <title>ログイン</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <GuestLogin />
    </>
  )
}

export default Login
