import React from 'react'
import Button from '@/components/parts/Button'

export type ContentHeaderProps = {
  isLoggedIn: boolean
  handleOpenModal: () => void
}

const ContentHeader: React.FC<ContentHeaderProps> = ({
  isLoggedIn,
  handleOpenModal,
}) => {
  return (
    <div className="flex">
      <div>
        <h2 className="text-gray-900 font-medium title-font text-xl mb-4">
          チャットルーム一覧
        </h2>
        <p className="leading-relaxed text-sm mb-8 text-gray-700">
          一覧からルームをクリックすると入室することができます。
        </p>
      </div>
      {isLoggedIn && (
        <div className="ml-auto">
          <Button onClick={handleOpenModal}>ルーム作成</Button>
        </div>
      )}
    </div>
  )
}

export default ContentHeader
