import React from 'react'

export type OnlineUser = {
  id: string
  name: string
  avatarUrl: string
}

export type OnlineUserListProps = {
  users: OnlineUser[]
}

const OnlineUserList: React.FC<OnlineUserListProps> = ({ users }) => {
  return (
    <div>
      <h3 className="mb-4 text-lg">ログイン中のユーザー</h3>
      <ul className="px-0 flex flex-col space-y-3">
        {users.map((user) => (
          <li
            key={user.id}
            className="flex items-center mb-0 space-x-3 text-gray-700"
          >
            <div className="w-12 border border-gray-200 rounded-full overflow-hidden">
              <img src={user.avatarUrl} alt={user.name + ' avatar'} />
            </div>
            <div>{user.name}</div>
            <div>
              <div className="bg-green-500 w-3 h-3 rounded-full"></div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default OnlineUserList
