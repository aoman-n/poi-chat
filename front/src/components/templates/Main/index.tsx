import React from 'react'
import { ROOM_SIZE } from '@/constants'
import Header from '@/components/organisms/Header'
import UsersNav from '@/components/organisms/UsersNav'

export type MainTemplateProps = {
  children: React.ReactNode
}

const MainTemplate: React.FC<MainTemplateProps> = ({ children }) => {
  return (
    <div>
      <header className="bg-white h-16 border-b border-gray-200">
        <div className="max-w-screen-xl mx-auto h-full">
          <Header isLoggedIn={true} />
        </div>
      </header>
      <main className="flex space-x-10 max-w-screen-lg mx-auto mt-24">
        {/* Canvasのレスポンシブは考慮する点が多いので一旦決め打ちのサイズで */}
        <div style={{ width: `${ROOM_SIZE.WIDTH}px` }}>{children}</div>
        <div className="w-80 min-w-80">
          <UsersNav />
        </div>
      </main>
    </div>
  )
}

export default MainTemplate
