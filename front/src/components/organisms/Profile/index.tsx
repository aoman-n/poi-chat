import React from 'react'

export type ProfileProps = {
  profile: {
    name: string
    avatarUrl: string
  }
}

const Profile: React.FC<ProfileProps> = ({ profile }) => {
  return (
    <div className="text-center py-8">
      <div className="w-20 h-20 rounded-full inline-flex items-center justify-center bg-gray-200 text-gray-400">
        <img src={profile.avatarUrl} alt="my avatar" />
      </div>
      <div className="flex flex-col items-center text-center justify-center">
        <h2 className="font-medium title-font mt-4 text-gray-900 text-lg">
          {profile.name}
        </h2>
        <div className={`w-12 h-1 bg-gray-500 rounded mt-2`}></div>
      </div>
    </div>
  )
}

export default Profile
