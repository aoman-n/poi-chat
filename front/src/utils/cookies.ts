import nookies, { parseCookies, setCookie, destroyCookie } from 'nookies'
import { GetServerSidePropsContext } from 'next'

const accessPathKey = 'prev_access_path'
const maxAge = 60 * 5

export const setAccessPathOnClient = (path: string) => {
  setCookie(null, accessPathKey, path, {
    maxAge,
    path: '/',
  })
}

export const destroyAccessPathOnClient = () => {
  destroyCookie(null, accessPathKey)
}

export const getAccessPathOnClient = () => {
  const cookies = parseCookies()
  if (cookies && cookies[accessPathKey]) {
    return cookies[accessPathKey]
  }

  return ''
}

export const setAccessPathOnServer = (
  ctx: GetServerSidePropsContext,
  path: string,
) => {
  nookies.set(ctx, accessPathKey, path, {
    maxAge,
    path: '/',
  })
}

export const destroyAccessPathOnServer = (ctx: GetServerSidePropsContext) => {
  nookies.destroy(ctx, accessPathKey)
}

export const getAccessPathOnServer = (ctx: GetServerSidePropsContext) => {
  const cookies = nookies.get(ctx)
  if (cookies && cookies[accessPathKey]) {
    return cookies[accessPathKey]
  }

  return ''
}
