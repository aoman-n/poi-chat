import React from 'react'

const CommentsSkeleton: React.VFC = () => {
  return (
    <div className="animate-pulse">
      <div
        className="bg-gray-300 mb-2"
        style={{ width: '85px', height: '20px' }}
      />
      <div className="bg-gray-300 mb-2" style={{ height: '210px' }} />
      <div className="bg-gray-300" style={{ height: '40px' }} />
    </div>
  )
}

export default CommentsSkeleton
