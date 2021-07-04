import React from 'react'
import { NextPage, GetServerSideProps } from 'next'
import IndexPageComponent from '@/components/pages/IndexPage'
import {
  getAccessPathOnServer,
  destroyAccessPathOnServer,
} from '@/utils/cookies'

const IndexPage: NextPage = () => {
  return <IndexPageComponent />
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const prevAccessPath = getAccessPathOnServer(ctx)

  if (prevAccessPath) {
    destroyAccessPathOnServer(ctx)
    return {
      redirect: {
        permanent: false,
        destination: prevAccessPath,
      },
    }
  }

  return {
    props: {
      title: 'ルーム一覧',
      layout: 'Main',
    },
  }
}

export default IndexPage
