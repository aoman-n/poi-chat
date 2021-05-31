import React from 'react'
import Link from 'next/link'
import { RoomListFragment } from '@/graphql'
import { getRoomIdParam } from '@/utils/ids'

export type Room = RoomListFragment['rooms']['nodes'][0]

export type RoomListProps = {
  rooms: RoomListFragment['rooms']['nodes']
}

const RoomList: React.FC<RoomListProps> = ({ rooms }) => {
  return (
    <div>
      <h2 className="text-gray-900 font-medium title-font text-2xl mb-4">
        チャットルーム一覧
      </h2>
      <p className="leading-relaxed text-base mb-8">
        一覧からルームをクリックすると入室することができます。
      </p>
      <div>
        {rooms.map((room) => (
          <Link key={room.id} href={`/rooms/${getRoomIdParam(room.id)}`}>
            <a>
              <div className="h-full flex items-center border-gray-200 border-b px-2 py-4 hover:bg-gray-100 duration-150 cursor-pointer">
                <img
                  alt="team"
                  className="w-12 h-12 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4"
                  src="https://dummyimage.com/80x80"
                />
                <div className="flex-grow">
                  <h4 className="text-gray-900 font-medium text-lg">
                    {room.name}
                  </h4>
                  <div className="text-gray-500 my-1 flex">
                    {/* UI Designer */}
                    <div className="ml-auto flex items-center space-x-1">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-5 w-5"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"
                        />
                      </svg>
                      <span>{room.userCount}</span>
                    </div>
                  </div>
                </div>
              </div>
            </a>
          </Link>
        ))}
      </div>
    </div>
  )
}

export default RoomList
