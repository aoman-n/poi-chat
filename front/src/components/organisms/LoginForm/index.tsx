import React from 'react'
import cn from 'classnames'
import Link from 'next/link'
import config from '@/config'

export type LoginFormProps = {
  noop?: string
}

const LoginForm: React.FC<LoginFormProps> = () => {
  return (
    <>
      <div className="text-center py-8 text-lg text-gray-800 border-b border-gray-200">
        ログインして始める
      </div>
      <div className="px-8">
        <div>
          <div className="text-center py-6">
            <Link href={new URL('/twitter/oauth', config.apiBaseUrl).href}>
              <a
                style={{ backgroundColor: 'rgb(29, 161, 242)' }}
                className={cn(
                  'inline-block',
                  'py-3',
                  'w-full',
                  'text-white',
                  'rounded-sm',
                  'font-semibold',
                  'tracking-wide',
                  'duration-200',
                  'hover:opacity-90',
                  'focus:outline-none',
                )}
              >
                Twitter Login
              </a>
            </Link>
          </div>

          <div className="flex md:justify-between justify-center items-center">
            <div
              style={{ height: '1px' }}
              className="bg-gray-300 md:block hidden w-5/12"
            ></div>
            <span className="mx-1 text-sm font-light text-gray-400">or</span>
            <div
              style={{ height: '1px' }}
              className="bg-gray-300 md:block hidden w-5/12"
            ></div>
          </div>

          <div className="text-center py-6">
            <Link href="/guest-login">
              <a
                className={cn(
                  'inline-block',
                  'py-3',
                  'w-full',
                  'text-white',
                  'rounded-sm',
                  'font-semibold',
                  'tracking-wide',
                  'bg-gray-800',
                  'duration-200',
                  'hover:opacity-90',
                  'focus:outline-none',
                )}
              >
                Guest Login
              </a>
            </Link>
          </div>
        </div>
      </div>
    </>
  )
}

export default LoginForm
