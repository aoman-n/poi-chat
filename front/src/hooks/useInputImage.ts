import { useState, useRef, useCallback } from 'react'

export const useInputImage = (defaultImageUrl = '') => {
  const [imageUrl, setImageUrl] = useState(defaultImageUrl)
  const fileRef = useRef<HTMLInputElement>(null)

  const handleChangeFile = useCallback(() => {
    if (fileRef.current && fileRef.current.files) {
      const { createObjectURL } = window.URL || window.webkitURL
      const imageUrl = createObjectURL(fileRef.current.files[0])
      setImageUrl(imageUrl)
    }
  }, [fileRef])

  return { fileRef, handleChangeFile, imageUrl }
}
