import { AppGetStaticProps } from '@/types'
import LoginForm from '@/components/organisms/LoginForm'

const Login = () => {
  return <LoginForm />
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ログイン',
      layout: 'Entrance',
    },
  }
}

export default Login
