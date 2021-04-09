import React from 'react'

const Header: React.FC = () => {
  return (
    <div className="h-full mx-4 flex items-center">
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

    // </div>
    // </div>
  )
}

export default Header
