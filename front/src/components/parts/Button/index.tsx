import React from 'react'
import cn from 'classnames'

type Color = 'gray' | 'green' | 'red'
type FontSize = 'ss' | 's' | 'm' | 'l'

export type ButtonProps = {
  children: React.ReactNode
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void
  type?: 'button' | 'submit' | 'reset'
  disabled?: boolean
  fontBold?: boolean
  fontSize?: FontSize
  fullWidth?: boolean
  fullHeight?: boolean
  outline?: boolean
  color?: Color
  elementType?: 'button' | 'label'
  classNames?: string
}

const fillColorClasses = (color: Color): string => {
  return `bg-${color}-600 hover:bg-${color}-700 text-white`
}

const outlineColorClasses = (color: Color): string => {
  return `border border-${color}-400 text-${color}-700 hover:border-${color}-900`
}

const fontSizeClasses = (fontSize: FontSize): string => {
  switch (fontSize) {
    case 'ss':
      return 'px-4 py-2 text-xs'
    case 's':
      return 'px-6 py-2 text-sm'
    case 'm':
      return 'px-6 py-2 text-base'
    case 'l':
      return 'px-6 py-2 text-lg'
  }
}

const Button: React.FC<ButtonProps> = ({
  children,
  onClick = () => {},
  type = 'button',
  disabled = false,
  fontBold = false,
  fontSize = 's',
  fullWidth = false,
  fullHeight = false,
  outline = false,
  color = 'gray',
  elementType = 'button',
  classNames = '',
}) => {
  return React.createElement(
    elementType,
    {
      className: cn(
        fontSizeClasses(fontSize),
        { 'w-full': fullWidth },
        { 'h-full': fullHeight },
        outline ? outlineColorClasses(color) : fillColorClasses(color),
        'duration-75',
        'rounded-sm',
        'focus:outline-none',
        'disabled:opacity-40',
        'cursor-pointer',
        { 'font-semibold': fontBold },
        { 'cursor-default': disabled },
        classNames,
      ),
      type,
      onClick,
      disabled,
    },
    children,
  )
}

export default Button
