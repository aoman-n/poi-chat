import React from 'react'

const SettingsSkeleton: React.VFC = () => {
  return (
    <div className="h-full flex animate-pulse">
      <div>
        <div
          className="bg-gray-300 mb-2"
          style={{ width: '130px', height: '20px' }}
        />
        <div
          className="bg-gray-300"
          style={{ width: '180px', height: '80px' }}
        />
      </div>
      <div
        className="bg-gray-300 ml-auto"
        style={{ width: '150px', height: '38px' }}
      />
    </div>
  )
}

export default SettingsSkeleton
