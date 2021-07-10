import React from 'react'

export type FrameProps = {
  contentHeader: React.ReactNode
  roomList: React.ReactNode
}

const Frame: React.VFC<FrameProps> = ({ contentHeader, roomList }) => {
  return (
    <div style={{ width: '600px' }}>
      <div className="mb-8">{contentHeader}</div>
      <div>{roomList}</div>
    </div>
  )
}

export default Frame
