import { gql } from '@apollo/client'
import * as Apollo from '@apollo/client'
export type Maybe<T> = T | null
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]: Maybe<T[SubKey]> }
const defaultOptions = {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
  Time: string
}

export type Connection = {
  pageInfo: PageInfo
  edges: Array<Maybe<Edge>>
  nodes: Array<Maybe<Node>>
}

export type CreateRoomInput = {
  name: Scalars['String']
}

export type Edge = {
  cursor: Scalars['String']
  node: Node
}

export type ExitedUser = {
  __typename?: 'ExitedUser'
  id: Scalars['ID']
}

export type JoinedUser = {
  __typename?: 'JoinedUser'
  id: Scalars['ID']
  displayName: Scalars['String']
  avatarUrl: Scalars['String']
  x: Scalars['Int']
  y: Scalars['Int']
}

export type Me = {
  __typename?: 'Me'
  id: Scalars['ID']
  displayName: Scalars['String']
  avatarUrl: Scalars['String']
}

export type Message = Node & {
  __typename?: 'Message'
  id: Scalars['ID']
  userId: Scalars['ID']
  userName: Scalars['String']
  userAvatarUrl: Scalars['String']
  body: Scalars['String']
  createdAt: Scalars['Time']
}

export type MessageConnection = Connection & {
  __typename?: 'MessageConnection'
  pageInfo: PageInfo
  edges: Array<Maybe<MessageEdge>>
  nodes: Array<Maybe<Message>>
  messageCount: Scalars['Int']
}

export type MessageEdge = Edge & {
  __typename?: 'MessageEdge'
  cursor: Scalars['String']
  node: Message
}

export type MoveInput = {
  roomId: Scalars['ID']
  x: Scalars['Int']
  y: Scalars['Int']
}

export type MovedUser = {
  __typename?: 'MovedUser'
  id: Scalars['ID']
  x: Scalars['Int']
  y: Scalars['Int']
}

export type Mutation = {
  __typename?: 'Mutation'
  sendMessage: Message
  createRoom: Room
  move: MovedUser
}

export type MutationSendMessageArgs = {
  input?: Maybe<SendMessageInput>
}

export type MutationCreateRoomArgs = {
  input?: Maybe<CreateRoomInput>
}

export type MutationMoveArgs = {
  input: MoveInput
}

export type Node = {
  id: Scalars['ID']
}

export type NoopPayload = {
  __typename?: 'NoopPayload'
  clientMutationId?: Maybe<Scalars['String']>
}

export type PageInfo = {
  __typename?: 'PageInfo'
  startCursor?: Maybe<Scalars['String']>
  endCursor?: Maybe<Scalars['String']>
  hasNextPage: Scalars['Boolean']
  hasPreviousPage: Scalars['Boolean']
}

export type PaginationInput = {
  first?: Maybe<Scalars['Int']>
  after?: Maybe<Scalars['String']>
  last?: Maybe<Scalars['Int']>
  before?: Maybe<Scalars['String']>
}

export type Query = {
  __typename?: 'Query'
  node?: Maybe<Node>
  rooms: RoomConnection
  roomDetail: RoomDetail
  me: Me
}

export type QueryNodeArgs = {
  id: Scalars['ID']
}

export type QueryRoomsArgs = {
  first?: Maybe<Scalars['Int']>
  after?: Maybe<Scalars['String']>
  orderBy?: Maybe<RoomOrderField>
}

export type QueryRoomDetailArgs = {
  id: Scalars['ID']
}

export type Room = Node & {
  __typename?: 'Room'
  id: Scalars['ID']
  name: Scalars['String']
}

export type RoomConnection = Connection & {
  __typename?: 'RoomConnection'
  pageInfo: PageInfo
  edges: Array<Maybe<RoomEdge>>
  nodes: Array<Maybe<Room>>
  roomCount: Scalars['Int']
}

export type RoomDetail = Node & {
  __typename?: 'RoomDetail'
  id: Scalars['ID']
  name: Scalars['String']
  messages: MessageConnection
  users: Array<User>
}

export type RoomDetailMessagesArgs = {
  last?: Maybe<Scalars['Int']>
  before?: Maybe<Scalars['String']>
}

export type RoomEdge = Edge & {
  __typename?: 'RoomEdge'
  cursor: Scalars['String']
  node: Room
}

export enum RoomOrderField {
  Latest = 'LATEST',
  DescUserCount = 'DESC_USER_COUNT',
}

export type SendMessageInput = {
  roomID: Scalars['ID']
  body: Scalars['String']
}

export type Subscription = {
  __typename?: 'Subscription'
  subMessage: Message
  subUserEvent: UserEvent
  joinRoom: User
}

export type SubscriptionSubMessageArgs = {
  roomId: Scalars['ID']
}

export type SubscriptionSubUserEventArgs = {
  roomId: Scalars['ID']
}

export type SubscriptionJoinRoomArgs = {
  roomID: Scalars['ID']
}

export type User = {
  __typename?: 'User'
  id: Scalars['ID']
  displayName: Scalars['String']
  avatarUrl: Scalars['String']
  x: Scalars['Int']
  y: Scalars['Int']
}

export type UserEvent = MovedUser | ExitedUser | JoinedUser

export type IndexQueryVariables = Exact<{ [key: string]: never }>

export type IndexQuery = { __typename?: 'Query' } & {
  rooms: { __typename?: 'RoomConnection' } & Pick<
    RoomConnection,
    'roomCount'
  > & {
      pageInfo: { __typename?: 'PageInfo' } & Pick<
        PageInfo,
        'startCursor' | 'endCursor' | 'hasNextPage' | 'hasPreviousPage'
      >
      edges: Array<
        Maybe<
          { __typename?: 'RoomEdge' } & Pick<RoomEdge, 'cursor'> & {
              node: { __typename?: 'Room' } & Pick<Room, 'id' | 'name'>
            }
        >
      >
      nodes: Array<Maybe<{ __typename?: 'Room' } & Pick<Room, 'id' | 'name'>>>
    }
}

export const IndexDocument = gql`
  query Index {
    rooms {
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        cursor
        node {
          id
          name
        }
      }
      nodes {
        id
        name
      }
      roomCount
    }
  }
`

/**
 * __useIndexQuery__
 *
 * To run a query within a React component, call `useIndexQuery` and pass it any options that fit your needs.
 * When your component renders, `useIndexQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useIndexQuery({
 *   variables: {
 *   },
 * });
 */
export function useIndexQuery(
  baseOptions?: Apollo.QueryHookOptions<IndexQuery, IndexQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<IndexQuery, IndexQueryVariables>(
    IndexDocument,
    options,
  )
}
export function useIndexLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<IndexQuery, IndexQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<IndexQuery, IndexQueryVariables>(
    IndexDocument,
    options,
  )
}
export type IndexQueryHookResult = ReturnType<typeof useIndexQuery>
export type IndexLazyQueryHookResult = ReturnType<typeof useIndexLazyQuery>
export type IndexQueryResult = Apollo.QueryResult<
  IndexQuery,
  IndexQueryVariables
>
