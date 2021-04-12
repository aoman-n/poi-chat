import { AppGetStaticProps } from '@/types'
import Entrance from '@/components/templates/Entrance'
import LoginForm from '@/components/organisms/LoginForm'

const Login = () => {
  return (
    <>
      <Entrance MainComponent={<LoginForm />} />
    </>
  )
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
