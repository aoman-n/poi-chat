import React from 'react'

export type ProfileProps = {
  name: string
  avatarUrl: string
}

const Profile: React.FC<ProfileProps> = ({ name, avatarUrl }) => {
  return (
    <div className="text-center py-8">
      <div className="w-20 h-20 rounded-full overflow-hidden ring-2 ring-white mx-auto">
        <img src={avatarUrl} alt="my avatar" height="100" width="100" />
      </div>
      <div className="flex flex-col items-center text-center justify-center">
        <h2 className="font-medium title-font mt-4 text-gray-900 text-xl">
          {name}
        </h2>
        <div className={`w-12 h-1 bg-gray-500 rounded mt-2`}></div>
      </div>
    </div>
  )
}

export default Profile
