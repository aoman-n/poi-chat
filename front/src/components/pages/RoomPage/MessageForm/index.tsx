import React, { useState, useCallback } from 'react'

export type MessageFormProps = {
  handleSubmitMessage: (values: { body: string }) => void
}

const MessageForm: React.VFC<MessageFormProps> = ({ handleSubmitMessage }) => {
  const [inputMesage, setInputMessage] = useState('')
  const wrappedHandleSubmitMessage = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault()
      handleSubmitMessage({ body: inputMesage })
      setInputMessage('')
    },
    [handleSubmitMessage, inputMesage],
  )

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputMessage(e.target.value)
  }, [])

  return (
    <form
      className={['mt-6', 'text-gray-900', 'flex'].join(' ')}
      onSubmit={wrappedHandleSubmitMessage}
    >
      <input
        id="username"
        type="text"
        className="px-3 py-2 bg-white focus:outline-none text-sm flex-grow"
        placeholder="コメントを入力"
        autoComplete="off"
        onChange={handleChange}
        value={inputMesage}
      />
      <button className="px-4 bg-green-600 text-white font-medium opacity-80 hover:opacity-100 duration-100 focus:outline-none">
        送信
      </button>
    </form>
  )
}

export default MessageForm
