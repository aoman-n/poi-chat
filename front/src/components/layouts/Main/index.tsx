import React from 'react'
import { ROOM_SCREEN_SIZE } from '@/constants'
import Header from '@/components/domainParts/Header'
import UsersNav from '@/components/domainParts/UsersNav'

export type MainTemplateProps = {
  children: React.ReactNode
}

const MainTemplate: React.FC<MainTemplateProps> = ({ children }) => {
  return (
    <div>
      <header className="bg-white h-16 border-b border-gray-200">
        <div className="max-w-screen-xl mx-auto h-full">
          <Header />
        </div>
      </header>
      <main className="flex space-x-10 max-w-screen-lg mx-auto mt-24">
        <aside className="flex-1" style={{ minWidth: '240px' }}>
          <UsersNav />
        </aside>
        {/* 一旦決め打ちのサイズで */}
        <div style={{ width: `${ROOM_SCREEN_SIZE.WIDTH}px` }}>{children}</div>
      </main>
    </div>
  )
}

export default MainTemplate
