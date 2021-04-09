import React from 'react'

export const isExistsRef = <T>(
  ref: React.MutableRefObject<T | null>,
): ref is React.MutableRefObject<T> => {
  if (ref && ref.current) {
    return true
  }

  return false
}
