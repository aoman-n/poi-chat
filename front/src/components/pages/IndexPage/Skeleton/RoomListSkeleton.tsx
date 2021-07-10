import React from 'react'

const MessagesSkeleton: React.VFC = () => {
  return (
    <div className="animate-pulse">
      <div className="bg-gray-300 mb-2" style={{ height: '500px' }} />
    </div>
  )
}

export default MessagesSkeleton
