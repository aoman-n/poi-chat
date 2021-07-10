import React from 'react'

const MessagesSkeleton: React.VFC = () => {
  return (
    <div className="animate-pulse">
      <div
        className="bg-gray-300 mb-2"
        style={{ width: '85px', height: '20px' }}
      />
      <div className="bg-gray-300 mb-2" style={{ height: '410px' }} />
    </div>
  )
}

export default MessagesSkeleton
