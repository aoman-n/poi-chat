import React from 'react'
import Link from 'next/link'
import Profile, { ProfileProps } from '@/components/parts/Profile'
import Button from '@/components/parts/Button'
import OnlineUserList, {
  OnlineUserListProps,
} from '@/components/parts/OnlineUserList'

export type UsersNavProps = {
  isLoggedIn: boolean
  profile: ProfileProps | null
  onlineUserList: OnlineUserListProps
}

const UsersNav: React.FC<UsersNavProps> = ({
  isLoggedIn,
  profile,
  onlineUserList,
}) => {
  if (!isLoggedIn) {
    return (
      <div className="border border-gray-400 text-center py-6 px-6">
        <p className="text-gray-800 mb-3">
          ルームに入室してチャットを開始するには、ログインしてください。
        </p>
        <Link href="/login">
          <a>
            <Button>ログインする</Button>
          </a>
        </Link>
      </div>
    )
  }

  return (
    <>
      {profile && (
        <div className="border border-gray-400 mb-8">
          <Profile {...profile} />
        </div>
      )}
      <div className="overflow-y-auto">
        <OnlineUserList {...onlineUserList} />
      </div>
    </>
  )
}

export default UsersNav
