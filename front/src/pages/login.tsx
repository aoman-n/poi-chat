import React from 'react'
import { AppGetStaticProps } from '@/types'
import LoginPageComponent from '@/components/pages/LoginPage'
import { useRequireNotLogin } from '@/hooks'

const LoginPage = () => {
  useRequireNotLogin()

  return <LoginPageComponent />
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ログイン',
      layout: 'Entrance',
    },
  }
}

export default LoginPage
