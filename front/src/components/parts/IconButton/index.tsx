import React from 'react'
import cn from 'classnames'

export type IconButtonProps = {
  classNames?: string
  children: React.ReactNode
}

const IconButton: React.FC<IconButtonProps> = ({ classNames, children }) => {
  return (
    <button
      className={cn(
        'focus:outline-none',
        'rounded-full',
        'p-2',
        'hover:bg-gray-200',
        'duration-75',
        'cursor-pointer',
        classNames,
      )}
    >
      {children}
    </button>
  )
}

export default IconButton
