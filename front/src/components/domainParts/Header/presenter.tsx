import React from 'react'
import Link from 'next/link'
import styles from './index.module.scss'
import Icon from '@/components/parts/Icon'
import IconButton from '@/components/parts/IconButton'
import Dropdown from '@/components/parts/Dropdown'
import OnlineUserList, {
  OnlineUserListProps,
} from '@/components/parts/OnlineUserList'
import config from '@/config'

export type Profile = {
  name: string
  avatarUrl: string
}

export type HeaderProps = {
  profile: Profile | null
  onlineUserList: OnlineUserListProps
}

const Header: React.FC<HeaderProps> = ({ profile, onlineUserList }) => {
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
      {profile && (
        <div className="ml-auto flex items-center space-x-6">
          {/* オペレータリスト */}
          <Dropdown
            leftPx={-324}
            button={
              <IconButton>
                <Icon type="users" />
              </IconButton>
            }
          >
            <div style={{ width: '360px' }}>
              <OnlineUserList {...onlineUserList} />
            </div>
          </Dropdown>
          <Dropdown
            leftPx={-234}
            button={
              <button className="focus:outline-none">
                <img
                  src={profile.avatarUrl}
                  className="rounded-full"
                  height={40}
                  width={40}
                />
              </button>
            }
          >
            <div style={{ width: '270px' }}>
              <div>
                <div className="flex items-center justify-center py-5 space-x-4 border border-gray-200">
                  <div className="inline-block rounded-full">
                    <img
                      src={profile.avatarUrl}
                      className="rounded-full"
                      height="54"
                      width="54"
                    />
                  </div>
                  <p className="text-gray-700 text-lg">{profile.name}</p>
                </div>
                <div>
                  <Link href={`${config.apiBaseUrl}/logout`}>
                    <a className="block py-4 text-center hover:bg-gray-100 duration-75">
                      ログアウト
                    </a>
                  </Link>
                </div>
              </div>
            </div>
          </Dropdown>
        </div>
      )}
    </div>
  )
}

export default React.memo(Header)
