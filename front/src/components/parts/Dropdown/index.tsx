import React, { useRef } from 'react'
import cn from 'classnames'
import { useOpen } from '@/hooks'

type DropdownProps = {
  button: React.ReactNode
  classNames?: string
  children: React.ReactNode
  leftPx?: number
}

const Dropdown: React.FC<DropdownProps> = ({
  button,
  classNames,
  children,
  leftPx = 0,
}) => {
  const targetRef = useRef<null | HTMLDivElement>(null)
  const { isOpen, handleOpen } = useOpen(targetRef)

  return (
    <div className={cn('relative', classNames)}>
      <div className="flex items-center" onClick={handleOpen}>
        {button}
      </div>
      <div
        ref={targetRef}
        style={{ left: `${leftPx}px` }}
        className={cn('absolute', 'top-10', 'z-10', {
          hidden: !isOpen,
        })}
      >
        <div className="bg-white border border-gray-200 shadow-xl rounded-sm">
          {children}
        </div>
      </div>
    </div>
  )
}

export default Dropdown
