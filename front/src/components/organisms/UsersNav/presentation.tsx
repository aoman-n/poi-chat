import React from 'react'
import Profile, { ProfileProps } from '@/components/organisms/Profile'
import OnlineUserList, {
  OnlineUserListProps,
} from '@/components/organisms/OnlineUserList'

export type UsersNavProps = {
  profile: ProfileProps
  onlineUserList: OnlineUserListProps
}

const UsersNav: React.FC<UsersNavProps> = ({ profile, onlineUserList }) => {
  return (
    <>
      <div className="border border-gray-400 mb-8">
        <Profile {...profile} />
      </div>
      <div className="overflow-y-auto">
        <OnlineUserList {...onlineUserList} />
      </div>
    </>
  )
}

export default UsersNav
