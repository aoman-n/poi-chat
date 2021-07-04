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
      {rooms.map((room) => (
        <Link key={room.id} href={`/rooms/${getRoomIdParam(room.id)}`}>
          <a>
            <div className="h-full flex items-center border-gray-200 border-b px-2 py-2 hover:bg-gray-200 duration-75 cursor-pointer">
              <div className="flex-grow">
                <h4 className="text-gray-900 font-medium">{room.name}</h4>
                <div className="text-gray-500 my-1 flex">
                  <div className="ml-auto flex space-x-4">
                    <div className="flex items-center space-x-1">
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
                          strokeWidth="2"
                          d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"
                        />
                      </svg>
                      <span>{room.totalMessageCount}</span>
                    </div>
                    <div className="flex items-center space-x-1">
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
                      <span>{room.totalUserCount}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </a>
        </Link>
      ))}
    </div>
  )
}

export default RoomList
