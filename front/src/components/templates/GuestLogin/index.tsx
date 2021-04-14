import React from 'react'
import Header from '@/components/organisms/Header'

export type GuestLoginProps = {
  children: React.ReactNode
}

const GuestLogin: React.FC<GuestLoginProps> = ({ children }) => {
  return (
    <div className="h-screen bg-gray-100 flex flex-col">
      <header className="bg-white h-16 border-b border-gray-200">
        <div className="max-w-screen-xl mx-auto h-full">
          <Header isLoggedIn={false} />
        </div>
      </header>
      <main className="pt-24 mx-auto">{children}</main>
    </div>
  )
}

export default GuestLogin
