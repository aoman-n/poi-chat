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
      <div className="flex items-center justify-center py-5 space-x-2 border border-gray-200">
        <span className="inline-block bg-green-500 w-2 h-2 rounded-full" />
        <p className="text-gray-700">ログイン中のユーザー一覧</p>
      </div>
      <ul className="overflow-y-auto py-2" style={{ height: '400px' }}>
        {users.map((user) => (
          <li
            key={user.id}
            className="flex items-center px-4 py-2 hover:bg-gray-50"
          >
            <div
              className="border border-gray-200 rounded-full overflow-hidden mr-3"
              style={{ width: '40px' }}
            >
              <img
                src={user.avatarUrl}
                alt={user.name + ' avatar'}
                height="100"
                width="100"
              />
            </div>
            <div
              className="flex-1 text-sm text-gray-600"
              style={{
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
              }}
            >
              {user.name}
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default OnlineUserList
