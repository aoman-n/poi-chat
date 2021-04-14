import { AppGetStaticProps } from '@/types'
import GuestLoginForm from '@/components/organisms/GuestLoginForm'

const Login = () => {
  return <GuestLoginForm />
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
