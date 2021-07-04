import React from 'react'
import { ArrowRight, ArrowLeft, Check, Upload } from './icons'

const icons = {
  arrowRight: ArrowRight,
  arrowLeft: ArrowLeft,
  check: Check,
  upload: Upload,
}

type Color = 'gray' | 'green' | 'white'

const getColor = (color: Color) => {
  switch (color) {
    case 'gray':
      return 'text-gray-600'
    case 'green':
      return 'text-green-600'
    case 'white':
      return 'text-white'
  }
}

type IconProps = {
  type: keyof typeof icons
  color?: Color
  className?: string
}

const Icon: React.FC<IconProps> = ({ type, color = 'gray', className }) => {
  const Component = icons[type]

  return (
    <Component
      className={['w-7', 'h-7', getColor(color), className].join(' ')}
    />
  )
}

export default Icon
