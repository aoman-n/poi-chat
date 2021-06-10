import React, { useState } from 'react'
import cn from 'classnames'
import CropIcon from '@/components/organisms/CropIcon'
import { useInputImage } from '@/hooks'
import { loginByGuest } from '@/utils/api'

export type GuestLoginFormProps = {
  noop?: string
}

// TODO: リファクタリング/imageのupdateに再レンダリングが多く走っている
const GuestLoginForm: React.FC<GuestLoginFormProps> = () => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [username, setUsername] = useState('名無しさん')
  const [imageBlob, setImageBlob] = useState<Blob | null>(null)
  const { fileRef, handleChangeFile, imageUrl } = useInputImage(
    'http://placekitten.com/500/800',
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
        window.location.href = '/'
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
            <div className="text-sm">
              <label htmlFor="username" className="block text-black mb-3">
                アイコン画像
              </label>
              <label className="flex justify-center space-x-1 rounded-sm px-4 py-3 bg-green-500 shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue text-white w-full text-center hover:opacity-90">
                <span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
                    />
                  </svg>
                </span>
                <span className="text-base leading-normal">
                  アイコン画像を選択
                </span>
                <input
                  type="file"
                  className="hidden"
                  ref={fileRef}
                  onChange={() => handleChangeFile()}
                />
              </label>
            </div>
          </div>
          <div className="my-6">
            <button
              disabled={isSubmitting}
              type="submit"
              className={cn(
                'py-4',
                'w-full',
                'text-white',
                'rounded-sm',
                'tracking-wide',
                'bg-gray-800',
                'duration-200',
                'hover:opacity-90',
                'focus:outline-none',
              )}
            >
              ログインする
            </button>
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
