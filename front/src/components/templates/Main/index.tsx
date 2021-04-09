import React from 'react'

export type MainTemplateProps = {
  HeaderComponent: React.ReactNode
  MainComponent: React.ReactNode
  MyProfileComponent: React.ReactNode
  OnlineUserListComponent: React.ReactNode
}

const MainTemplate: React.FC<MainTemplateProps> = ({
  HeaderComponent,
  MainComponent,
  MyProfileComponent,
  OnlineUserListComponent,
}) => {
  return (
    <div>
      <header className="bg-white h-16 border-b border-gray-200">
        <div className="max-w-screen-xl mx-auto h-full">{HeaderComponent}</div>
      </header>
      <main className="flex space-x-10 max-w-screen-lg mx-auto mt-24">
        <div className="flex-grow" style={{ minWidth: '600px' }}>
          {MainComponent}
        </div>
        <div className="w-80 min-w-80">
          <div className="border border-gray-400 mb-8">
            {MyProfileComponent}
          </div>
          <div className="overflow-y-auto">{OnlineUserListComponent}</div>
        </div>
      </main>
    </div>
  )
}

export default MainTemplate
