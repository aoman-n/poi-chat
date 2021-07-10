import React from 'react'
import Header from '@/components/domainParts/Header'

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
      <main className="flex justify-center mt-12">{children}</main>
    </div>
  )
}

export default MainTemplate
