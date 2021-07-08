import React from 'react'
import cn from 'classnames'
import styles from './index.module.scss'

export type FrameProps = {
  screen: React.ReactNode
  settings: React.ReactNode
  comments: React.ReactNode
}

const Frame: React.VFC<FrameProps> = ({ screen, settings, comments }) => {
  return (
    <div className="h-full">
      <div className={cn('mb-4', styles.screenWrapper)}>
        <div className={styles.screenContent}>{screen}</div>
      </div>
      <div className="mb-4">{settings}</div>
      <div>{comments}</div>
    </div>
  )
}

export default Frame
