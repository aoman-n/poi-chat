import React from 'react'
import Icon from '@/components/parts/Icon'

export type CircleArrowButtonProps = {
  arrowType: 'left' | 'right'
  onClick: () => void
  classNames: string
}

const CircleArrowButton: React.FC<CircleArrowButtonProps> = ({
  arrowType,
  onClick,
  classNames,
}) => {
  return (
    <button
      onClick={onClick}
      className={
        'focus:outline-none h-16 w-16 rounded-full border-2 border-gray-200 shadow-lg bg-white hover:bg-gray-100 duration-75 flex justify-center items-center ' +
        classNames
      }
    >
      {arrowType === 'left' && <Icon type="arrowLeft" />}
      {arrowType === 'right' && <Icon type="arrowRight" />}
    </button>
  )
}

export default CircleArrowButton
