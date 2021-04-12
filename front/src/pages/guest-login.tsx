import { AppGetStaticProps } from '@/types'
import GuestLogin from '@/components/templates/GuestLogin'
import Header from '@/components/organisms/Header'

const Login = () => {
  return (
    <>
      <GuestLogin HeaderComponent={<Header isLoggedIn={false} />} />
    </>
  )
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ゲストログイン',
      layout: 'GuestLogin',
    },
  }
}

export default Login
