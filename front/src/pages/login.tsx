import React from 'react'
import { AppGetStaticProps } from '@/types'
import LoginForm from '@/components/organisms/LoginForm'
import { useRequireNotLogin } from '@/hooks'

const Login = () => {
  useRequireNotLogin()

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
