import React from 'react'
import cn from 'classnames'
import { ROOM_SCREEN_SIZE } from '@/constants'
import styles from './index.module.scss'

export type FrameProps = {
  screen: React.ReactNode
  settings: React.ReactNode
  messageList: React.ReactNode
  messageFrom: React.ReactNode
}

const Frame: React.VFC<FrameProps> = ({
  screen,
  settings,
  messageList,
  messageFrom,
}) => {
  return (
    <div className="flex space-x-6">
      <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>
        <div className={cn('mb-4', styles.screenWrapper)}>
          <div className={styles.screenContent}>{screen}</div>
        </div>
        <div className="mb-4">{settings}</div>
        <div className="">{messageFrom}</div>
      </div>
      <div style={{ width: '300px' }}>{messageList}</div>
    </div>
  )
}

export default Frame
