import { makeVar } from '@apollo/client'
import { CommonQuery } from '@/graphql'

export const globalUsersVar = makeVar<CommonQuery['globalUsers']>([])
