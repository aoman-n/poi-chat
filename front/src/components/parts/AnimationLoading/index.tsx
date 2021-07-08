import React from 'react'
import styles from './index.module.scss'

const AnimationLoading: React.VFC = () => {
  return (
    <div>
      <div className={styles.loading}>
        <span>L</span>
        <span>o</span>
        <span>a</span>
        <span>d</span>
        <span>i</span>
        <span>n</span>
        <span>g</span>
        <span>...</span>
      </div>
    </div>
  )
}

export default AnimationLoading
