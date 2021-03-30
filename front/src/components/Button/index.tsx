import React from 'react'

export type ButtonProps = {
  text: string
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
}

const Button: React.FC<ButtonProps> = ({ text, onClick }) => {
  return (
    <button className="w-1/2 bg-blue-300" onClick={onClick}>
      {text}
    </button>
  )
}

export default Button
