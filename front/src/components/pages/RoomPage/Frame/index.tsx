import React from 'react'
import cn from 'classnames'
import styles from './index.module.scss'

export type FrameProps = {
  screen: React.ReactNode
  settings: React.ReactNode
  messages: React.ReactNode
  messageFrom: React.ReactNode
}

const Frame: React.VFC<FrameProps> = ({
  screen,
  settings,
  messages,
  messageFrom,
}) => {
  return (
    <div className="h-full flex space-x-6">
      <div className="flex-1">
        <div className={cn('mb-4', styles.screenWrapper)}>
          <div className={styles.screenContent}>{screen}</div>
        </div>
        <div className="mb-4">{settings}</div>
        <div className="">{messageFrom}</div>
      </div>
      <div style={{ width: '300px' }}>
        <div>{messages}</div>
      </div>
    </div>
  )
}

export default Frame
