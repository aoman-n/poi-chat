import React from 'react'
import Profile, { ProfileProps } from '@/components/parts/Profile'
import OnlineUserList, {
  OnlineUserListProps,
} from '@/components/parts/OnlineUserList'

export type UsersNavProps = {
  profile: ProfileProps | null
  onlineUserList: OnlineUserListProps
}

const UsersNav: React.FC<UsersNavProps> = ({ profile, onlineUserList }) => {
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
