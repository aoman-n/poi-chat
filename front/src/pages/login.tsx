import Head from 'next/head'
import Entrance from '@/components/templates/Entrance'
import LoginForm from '@/components/organisms/LoginForm'

const Login = () => {
  return (
    <>
      <Head>
        <title>ログイン</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Entrance MainComponent={<LoginForm />} />
    </>
  )
}

export default Login
