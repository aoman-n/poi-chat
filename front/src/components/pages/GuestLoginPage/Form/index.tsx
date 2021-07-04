import React, { useState } from 'react'
import { useInputImage } from '@/hooks'
import { loginByGuest } from '@/utils/api'
import {
  getAccessPathOnClient,
  destroyAccessPathOnClient,
} from '@/utils/cookies'
import config from '@/config'
import Button from '@/components/parts/Button'
import Icon from '@/components/parts/Icon'
import CropIcon from '../CropIcon'

export type GuestLoginFormProps = {
  noop?: string
}

// TODO: リファクタリング/imageのupdateに再レンダリングが多く走っている
const GuestLoginForm: React.FC<GuestLoginFormProps> = () => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [username, setUsername] = useState('名無しさん')
  const [imageBlob, setImageBlob] = useState<Blob | null>(null)
  const { fileRef, handleChangeFile, imageUrl } = useInputImage(
    `${config.frontBaseUrl}/defaultAvatar1.png`,
  )

  const handleChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value)
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (imageBlob && username) {
      setIsSubmitting(true)
      try {
        await loginByGuest({ name: username, image: imageBlob })
        const accessPath = getAccessPathOnClient()
        if (accessPath) {
          destroyAccessPathOnClient()
          window.location.href = accessPath
        } else {
          window.location.href = '/'
        }
      } catch (err) {
        setIsSubmitting(false)
        console.log(err)
        alert('error')
      }
    }
  }

  const handleUpdateImage = (blob: Blob) => {
    setImageBlob(blob)
  }

  return (
    <div className="flex space-x-8">
      <div className="bg-white px-8" style={{ width: '420px' }}>
        <p className="text-lg text-center text-gray-800 my-10">
          ゲストログイン
        </p>
        <form className="" onSubmit={handleSubmit}>
          <div className="border-b border-gray-200 pb-6">
            <div className="mb-5 text-sm">
              <label htmlFor="username" className="block text-black mb-3">
                ユーザー名
              </label>
              <input
                id="username"
                type="text"
                className="rounded-sm px-4 py-3 bg-gray-100 w-full focus:outline-none"
                placeholder="ユーザー名"
                value={username}
                onChange={handleChangeUsername}
              />
            </div>
            <div>
              <label
                htmlFor="username"
                className="block text-black mb-3 text-sm"
              >
                アイコン画像
              </label>
              <Button
                elementType="label"
                color="green"
                classNames="flex justify-center items-center space-x-1 h-12"
              >
                <span>
                  <Icon type="upload" color="white" />
                </span>
                <span className="text-base">アイコン画像を選択</span>
                <input
                  type="file"
                  className="hidden"
                  ref={fileRef}
                  onChange={handleChangeFile}
                />
              </Button>
            </div>
          </div>
          <div className="my-6">
            <Button
              disabled={isSubmitting}
              type="submit"
              fullWidth
              fontSize="m"
              classNames="h-12"
            >
              ログインする
            </Button>
          </div>
        </form>
      </div>

      {/* right content */}
      <div className="px-8 pt-6 border border-gray-200">
        <CropIcon
          height={400}
          width={400}
          imageUrl={imageUrl}
          handleUpdateImage={handleUpdateImage}
        />
      </div>
    </div>
  )
}

export default GuestLoginForm
