import React from 'react'
import cn from 'classnames'

export type GuestLoginProps = {
  noop?: string
}

const GuestLogin: React.FC<GuestLoginProps> = ({ noop }) => {
  return (
    <div className="h-screen bg-gray-100 flex flex-col">
      <header className="bg-white h-16 border-b border-gray-200">
        <div className="h-full px-4 max-w-screen-xl mx-auto">
          <h1 className="h-full m-0 flex items-center font-extrabold font-sans text-2xl text-gray-800">
            <span className="pr-1">poi chat</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="w-7"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z"
                clipRule="evenodd"
              />
            </svg>
          </h1>
        </div>
      </header>
      <main className="flex-grow flex justify-center items-center space-x-8">
        {/* left content */}
        <div className="bg-white px-8">
          <p className="text-lg text-gray-800 my-10">
            名前とアイコン画像を設定してログインしてください。
          </p>
          <form className="">
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
                />
              </div>
              <div className="text-sm">
                <label htmlFor="username" className="block text-black mb-3">
                  アイコン画像
                </label>
                <label className="block rounded-sm px-4 py-3 bg-green-500 shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue text-white w-full text-center hover:opacity-90">
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
                  <span className="text-base leading-normal">
                    アイコン画像を選択
                  </span>
                  <input type="file" className="hidden" />
                </label>
              </div>
            </div>
            <div className="my-6">
              <button
                type="submit"
                onClick={(e) => e.preventDefault()}
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
        <div className="bg-red-100">右(アイコンの設定)</div>
      </main>
    </div>
  )
}

export default GuestLogin
