import { AppGetStaticProps } from '@/types'
import GuestLoginPageComponent from '@/components/pages/GuestLoginPage'

const GuesLoginPage = () => {
  return <GuestLoginPageComponent />
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ゲストログイン',
      layout: 'GuestLogin',
    },
  }
}

export default GuesLoginPage
