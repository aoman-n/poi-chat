import { ApolloError } from '@apollo/client'

const INTERNAL_SERVER_ERROR_MSG = 'サーバーエラーが発生しました。'

export const createErrorMsg = (errorMsgMap: { [key in string]: string }) => (
  apolloError: ApolloError | undefined,
): string[] => {
  const errors: string[] = []

  if (apolloError && apolloError.graphQLErrors.length > 0) {
    for (const err of apolloError.graphQLErrors) {
      const msg = errorMsgMap[err.message]
      if (msg) errors.push(msg)
    }

    if (errors.length === 0) {
      errors.push(INTERNAL_SERVER_ERROR_MSG)
    }
  }

  return errors
}
