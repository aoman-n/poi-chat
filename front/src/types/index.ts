import 'next'
import { GetStaticProps, GetServerSideProps } from 'next'
import { AppProps } from 'next/app'

export type PageProps = {
  title: string
  layout: 'Main' | 'Entrance'
}

/* eslint @typescript-eslint/ban-types: 0 */
// @see: https://qiita.com/Takepepe/items/56acedaf94cb0d0c388f
declare module 'next' {
  // PageComponent に適用する型
  export type PageFC<P = {}, IP = P & PageProps> = NextPage<P, IP>
  type Override<T extends U, U> = Omit<T, keyof U> & U
  export type AppPageProps = Override<
    AppProps<PageProps>,
    {
      pageProps: PageProps
    }
  >
}

export type AppGetStaticProps = GetStaticProps<PageProps>
export type AppGetServerSideProps<P> = GetServerSideProps<PageProps & P>
