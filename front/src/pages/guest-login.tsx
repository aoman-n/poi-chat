import Head from 'next/head'
import GuestLogin from '@/components/templates/GuestLogin'
import Header from '@/components/organisms/Header'

const Login = () => {
  return (
    <>
      <Head>
        <title>ログイン</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <GuestLogin HeaderComponent={<Header isLoggedIn={false} />} />
    </>
  )
}

export default Login
