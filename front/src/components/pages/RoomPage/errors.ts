import { createErrorMsg } from '@/utils/errors'

const CREATE_ROOM_ERROR_MSGS: { [key in string]: string } = {
  'not found room': 'ルームが存在しません。',
}

export const getErrorMsg = createErrorMsg(CREATE_ROOM_ERROR_MSGS)
