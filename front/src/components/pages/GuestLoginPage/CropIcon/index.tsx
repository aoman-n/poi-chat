import React, { useEffect, useRef } from 'react'
import Croppie from 'croppie'
import { isExistsRef } from '@/utils/elements'
import 'croppie/croppie.css'

export type CropIconProps = {
  imageUrl: string
  height: number
  width: number
  handleUpdateImage: (blob: Blob) => void
}

// TODO: refactor
const CropIcon: React.FC<CropIconProps> = ({
  imageUrl,
  height,
  width,
  handleUpdateImage,
}) => {
  const croppieEl = useRef<HTMLDivElement | null>(null)
  const [croppie, setCroppie] = React.useState<Croppie | null>(null)

  /* eslint react-hooks/exhaustive-deps: 0 */
  useEffect(() => {
    if (isExistsRef(croppieEl)) {
      const croppieInstance = new Croppie(croppieEl.current, {
        enableExif: true,
        viewport: {
          height: 360,
          width: 360,
          type: 'circle',
        },
        boundary: {
          height,
          width,
        },
      })
      croppieInstance.bind({
        url: imageUrl,
      })
      setCroppie(croppieInstance)

      const updateEventCallback = () => {
        croppieInstance
          .result({
            type: 'blob',
            size: {
              width: 480,
              height: 480,
            },
          })
          .then((blob) => {
            handleUpdateImage(blob)
          })
      }

      croppieEl.current.addEventListener('update', updateEventCallback)

      croppieInstance
        .result({
          type: 'blob',
          size: {
            width: 480,
            height: 480,
          },
        })
        .then((blob) => {
          handleUpdateImage(blob)
        })

      return () => {
        if (isExistsRef(croppieEl)) {
          croppieEl.current.removeEventListener('update', updateEventCallback)
        }
      }
    }
  }, [])

  useEffect(() => {
    if (croppie && imageUrl) {
      croppie.bind({
        url: imageUrl,
      })
    }
  }, [imageUrl, croppie])

  return <div ref={croppieEl} />
}

export default CropIcon
