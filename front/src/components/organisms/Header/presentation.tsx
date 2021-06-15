import React from 'react'
import Link from 'next/link'
import styles from './index.module.scss'
import config from '@/config'

export type HeaderProps = {
  isLoggedIn: boolean
}

const Header: React.FC<HeaderProps> = ({ isLoggedIn }) => {
  return (
    <div
      className={[
        'h-full',
        'px-6',
        'flex',
        'items-center',
        styles.container,
      ].join(' ')}
    >
      <Link href="/">
        <a>
          <h1 className="h-full m-0 flex items-center font-extrabold font-sans text-2xl text-gray-800">
            <span className="pr-1">poi chat</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="w-7"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z"
                clipRule="evenodd"
              />
            </svg>
          </h1>
        </a>
      </Link>
      <div className="ml-auto">
        {isLoggedIn ? (
          <Link href={`${config.apiBaseUrl}/logout`}>
            <a className="font-semibold text-gray-700 py-2 px-3 hover:bg-gray-100 duration-100 rounded-md">
              ログアウト
            </a>
          </Link>
        ) : (
          <Link href="/login">
            <a className="font-semibold text-gray-700 py-2 px-3 hover:bg-gray-100 duration-100 rounded-md">
              ログイン
            </a>
          </Link>
        )}
      </div>
    </div>
  )
}

export default React.memo(Header)
