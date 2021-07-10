import React from 'react'

const CommentFormSkeleton: React.VFC = () => {
  return (
    <div className="animate-pulse">
      <div className="bg-gray-300" style={{ height: '40px' }} />
    </div>
  )
}

export default CommentFormSkeleton
