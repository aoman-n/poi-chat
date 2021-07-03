import { ApolloError } from '@apollo/client'

const INTERNAL_SERVER_ERROR_MSG = 'サーバーエラーが発生しました。'

const CREATE_ROOM_ERROR_MSGS: { [key in string]: string } = {
  'already exists room name':
    '既に存在するルーム名です。別の名前に変更してください。',
}

export const getCreateRoomErrorMsg = (
  apolloError: ApolloError | undefined,
): string[] => {
  const errors: string[] = []

  if (apolloError && apolloError.graphQLErrors.length > 0) {
    for (const err of apolloError.graphQLErrors) {
      const msg = CREATE_ROOM_ERROR_MSGS[err.message]
      if (msg) errors.push(msg)
    }

    if (errors.length === 0) {
      errors.push(INTERNAL_SERVER_ERROR_MSG)
    }
  }

  return errors
}
