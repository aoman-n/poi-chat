import React from 'react'
import cn from 'classnames'

export type EntranceProps = {
  noop?: string
}

const Entrance: React.FC<EntranceProps> = ({ noop }) => {
  return (
    <div className="h-screen bg-gray-100">
      <header className="pt-24 pb-6 text-center">
        <h1 className="inline-flex items-center text-gray-700 text-center">
          <span className="pr-2 text-4xl">poi chat</span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="w-11"
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
      </header>
      <main className="w-96 mx-auto bg-white">
        <div className="text-center py-8 text-lg text-gray-800 border-b border-gray-200">
          ログインして始める
        </div>
        <div className="px-8">
          <div>
            <div className="text-center py-6">
              <button
                style={{ backgroundColor: 'rgb(29, 161, 242)' }}
                className={cn(
                  'py-3',
                  'w-full',
                  'text-white',
                  'rounded-sm',
                  'font-semibold',
                  'tracking-wide',
                  'duration-200',
                  'hover:opacity-90',
                  'focus:outline-none',
                )}
              >
                Twitter Login
              </button>
            </div>

            <div className="flex md:justify-between justify-center items-center">
              <div
                style={{ height: '1px' }}
                className="bg-gray-300 md:block hidden w-5/12"
              ></div>
              <span className="mx-1 text-sm font-light text-gray-400">or</span>
              <div
                style={{ height: '1px' }}
                className="bg-gray-300 md:block hidden w-5/12"
              ></div>
            </div>

            <div className="text-center py-6">
              <button
                className={cn(
                  'py-3',
                  'w-full',
                  'text-white',
                  'rounded-sm',
                  'font-semibold',
                  'tracking-wide',
                  'bg-gray-800',
                  'duration-200',
                  'hover:opacity-90',
                  'focus:outline-none',
                )}
              >
                Guest Login
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default Entrance
