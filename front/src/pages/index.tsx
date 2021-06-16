import React from 'react'
import { NextPage } from 'next'
import IndexPageComponent from '@/components/pages/IndexPage'
import { AppGetStaticProps } from '@/types'

const IndexPage: NextPage = () => {
  return <IndexPageComponent />
}

export const getStaticProps: AppGetStaticProps = async () => {
  return {
    props: {
      title: 'ルーム一覧',
      layout: 'Main',
    },
  }
}

export default IndexPage
