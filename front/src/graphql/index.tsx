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
  Time: Date
}

export enum BalloonPosition {
  TopLeft = 'TOP_LEFT',
  TopRight = 'TOP_RIGHT',
  BottomLeft = 'BOTTOM_LEFT',
  BottomRight = 'BOTTOM_RIGHT',
}

export type ChangeBalloonPositionInput = {
  roomId: Scalars['ID']
  balloonPosition: BalloonPosition
}

export type ChangeBalloonPositionPayload = {
  __typename?: 'ChangeBalloonPositionPayload'
  roomUser?: Maybe<RoomUser>
}

export type ChangedBalloonPositionPayload = {
  __typename?: 'ChangedBalloonPositionPayload'
  roomUser: RoomUser
}

export type Connection = {
  pageInfo: PageInfo
  edges: Array<Maybe<Edge>>
  nodes: Array<Maybe<Node>>
}

export type CreateRoomInput = {
  name: Scalars['String']
  bgUrl?: Maybe<Scalars['String']>
  bgColor?: Maybe<Scalars['String']>
}

export type CreateRoomPayload = {
  __typename?: 'CreateRoomPayload'
  room: Room
}

export type Edge = {
  cursor: Scalars['String']
  node: Node
}

export type EnteredPayload = {
  __typename?: 'EnteredPayload'
  roomUser: RoomUser
}

export type ExitedPayload = {
  __typename?: 'ExitedPayload'
  userId: Scalars['ID']
}

export type Message = Node & {
  __typename?: 'Message'
  id: Scalars['ID']
  user: User
  body: Scalars['String']
  createdAt: Scalars['Time']
}

export type MessageConnection = Connection & {
  __typename?: 'MessageConnection'
  pageInfo: PageInfo
  edges: Array<MessageEdge>
  nodes: Array<Message>
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

export type MovePayload = {
  __typename?: 'MovePayload'
  roomUser: RoomUser
}

export type MovedPayload = {
  __typename?: 'MovedPayload'
  roomUser: RoomUser
}

export type Mutation = {
  __typename?: 'Mutation'
  /** ルームの作成 */
  createRoom: CreateRoomPayload
  /** メッセージの送信 */
  sendMessage: SendMassagePaylaod
  /** ルーム内ユーザーのポジション移動 */
  move: MovePayload
  /** ルーム内ユーザーの吹き出し削除 */
  removeLastMessage: RemoveLastMessagePayload
  /** ルーム内ユーザーの吹き出し位置変更 */
  changeBalloonPosition: ChangeBalloonPositionPayload
}

export type MutationCreateRoomArgs = {
  input?: Maybe<CreateRoomInput>
}

export type MutationSendMessageArgs = {
  input?: Maybe<SendMessageInput>
}

export type MutationMoveArgs = {
  input: MoveInput
}

export type MutationRemoveLastMessageArgs = {
  input: RemoveLastMessageInput
}

export type MutationChangeBalloonPositionArgs = {
  input: ChangeBalloonPositionInput
}

export type Node = {
  id: Scalars['ID']
}

export type OfflinedPayload = {
  __typename?: 'OfflinedPayload'
  user: User
}

export type OnlinedPayload = {
  __typename?: 'OnlinedPayload'
  user: User
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
  /** ルーム一覧を取得 */
  rooms: RoomConnection
  /** ルーム情報を取得 */
  room: Room
  /** ログイン中のユーザーが自身の情報を取得 */
  me: User
  /** オンライン中のユーザー一覧を取得 */
  onlineUsers: Array<User>
  node?: Maybe<Node>
}

export type QueryRoomsArgs = {
  first?: Maybe<Scalars['Int']>
  after?: Maybe<Scalars['String']>
  orderBy?: Maybe<RoomOrderField>
}

export type QueryRoomArgs = {
  id: Scalars['ID']
}

export type QueryNodeArgs = {
  id: Scalars['ID']
}

export type RemoveLastMessageInput = {
  roomId: Scalars['ID']
}

export type RemoveLastMessagePayload = {
  __typename?: 'RemoveLastMessagePayload'
  roomUser?: Maybe<RoomUser>
}

export type RemovedLastMessagePayload = {
  __typename?: 'RemovedLastMessagePayload'
  roomUser: RoomUser
}

export type Room = Node & {
  __typename?: 'Room'
  id: Scalars['ID']
  name: Scalars['String']
  bgColor: Scalars['String']
  bgUrl: Scalars['String']
  createdAt: Scalars['Time']
  totalUserCount: Scalars['Int']
  totalMessageCount: Scalars['Int']
  messages: MessageConnection
  users: Array<RoomUser>
}

export type RoomMessagesArgs = {
  last?: Maybe<Scalars['Int']>
  before?: Maybe<Scalars['String']>
}

export type RoomConnection = Connection & {
  __typename?: 'RoomConnection'
  pageInfo: PageInfo
  edges: Array<RoomEdge>
  nodes: Array<Room>
  roomCount: Scalars['Int']
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

export type RoomUser = {
  __typename?: 'RoomUser'
  id: Scalars['ID']
  user: User
  x: Scalars['Int']
  y: Scalars['Int']
  lastMessage?: Maybe<Message>
  balloonPosition: BalloonPosition
}

/** ルーム内のユーザーの行動を取得するためのイベントタイプ */
export type RoomUserEvent =
  | EnteredPayload
  | ExitedPayload
  | MovedPayload
  | SentMassagePayload
  | RemovedLastMessagePayload
  | ChangedBalloonPositionPayload

export type SendMassagePaylaod = {
  __typename?: 'SendMassagePaylaod'
  message: Message
}

export type SendMessageInput = {
  roomID: Scalars['ID']
  body: Scalars['String']
}

export type SentMassagePayload = {
  __typename?: 'SentMassagePayload'
  roomUser: RoomUser
}

export type Subscription = {
  __typename?: 'Subscription'
  /**
   * ユーザーのオンラインステータスの更新を待ち受けるサブスクリプション
   * このサブスクリプションを待ち受けると同時に自身をオンライン状態にする
   */
  actedUserEvent?: Maybe<UserEvent>
  /**
   * ルーム内ユーザーのアクションを待ち受けるサブスクリプション
   * このサブスクリプションを待ち受けると同時に自身をルームに入室させる
   */
  actedRoomUserEvent?: Maybe<RoomUserEvent>
}

export type SubscriptionActedRoomUserEventArgs = {
  roomId: Scalars['ID']
}

export type User = {
  __typename?: 'User'
  id: Scalars['ID']
  name: Scalars['String']
  avatarUrl: Scalars['String']
  /**  ルームに入室していなかったらnull  */
  enteredRoom?: Maybe<Room>
}

/** ユーザーのオンライン・オフライン状態の変更を取得するためのイベントタイプ */
export type UserEvent = OnlinedPayload | OfflinedPayload

export type CreateRoomMutationVariables = Exact<{
  name: Scalars['String']
  bgUrl: Scalars['String']
}>

export type CreateRoomMutation = { __typename?: 'Mutation' } & {
  createRoom: { __typename?: 'CreateRoomPayload' } & {
    room: { __typename?: 'Room' } & Pick<Room, 'id' | 'name'>
  }
}

export type IndexPageQueryVariables = Exact<{
  first: Scalars['Int']
  after?: Maybe<Scalars['String']>
}>

export type IndexPageQuery = { __typename?: 'Query' } & {
  rooms: { __typename?: 'RoomConnection' } & Pick<
    RoomConnection,
    'roomCount'
  > & {
      pageInfo: { __typename?: 'PageInfo' } & Pick<
        PageInfo,
        'startCursor' | 'endCursor' | 'hasNextPage' | 'hasPreviousPage'
      >
      nodes: Array<{ __typename?: 'Room' } & RoomFieldsForListFragment>
      edges: Array<
        { __typename?: 'RoomEdge' } & {
          node: { __typename?: 'Room' } & RoomFieldsForListFragment
        }
      >
    }
}

export type RoomFieldsForListFragment = { __typename?: 'Room' } & Pick<
  Room,
  'id' | 'name' | 'createdAt' | 'totalUserCount' | 'totalMessageCount'
>

export type RoomPageQueryVariables = Exact<{
  roomId: Scalars['ID']
  before?: Maybe<Scalars['String']>
}>

export type RoomPageQuery = { __typename?: 'Query' } & {
  room: { __typename?: 'Room' } & Pick<
    Room,
    'id' | 'name' | 'bgColor' | 'bgUrl'
  > & {
      users: Array<{ __typename?: 'RoomUser' } & RoomUserFieldsFragment>
      messages: { __typename?: 'MessageConnection' } & {
        pageInfo: { __typename?: 'PageInfo' } & Pick<
          PageInfo,
          'startCursor' | 'endCursor' | 'hasNextPage' | 'hasPreviousPage'
        >
        nodes: Array<{ __typename?: 'Message' } & MessageFieldsFragment>
      }
    }
}

export type MoreRoomMessagesQueryVariables = Exact<{
  roomId: Scalars['ID']
  before?: Maybe<Scalars['String']>
}>

export type MoreRoomMessagesQuery = { __typename?: 'Query' } & {
  room: { __typename?: 'Room' } & Pick<Room, 'id'> & PageMessagesFieldFragment
}

export type SendMessageMutationVariables = Exact<{
  roomId: Scalars['ID']
  body: Scalars['String']
}>

export type SendMessageMutation = { __typename?: 'Mutation' } & {
  sendMessage: { __typename?: 'SendMassagePaylaod' } & {
    message: { __typename?: 'Message' } & MessageFieldsFragment
  }
}

export type MoveMutationVariables = Exact<{
  roomId: Scalars['ID']
  x: Scalars['Int']
  y: Scalars['Int']
}>

export type MoveMutation = { __typename?: 'Mutation' } & {
  move: { __typename?: 'MovePayload' } & {
    roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
  }
}

export type RemoveLastMessageMutationVariables = Exact<{
  roomId: Scalars['ID']
}>

export type RemoveLastMessageMutation = { __typename?: 'Mutation' } & {
  removeLastMessage: { __typename?: 'RemoveLastMessagePayload' } & {
    roomUser?: Maybe<{ __typename?: 'RoomUser' } & RoomUserFieldsFragment>
  }
}

export type ChangeBalloonPositionMutationVariables = Exact<{
  roomId: Scalars['ID']
  balloonPosition: BalloonPosition
}>

export type ChangeBalloonPositionMutation = { __typename?: 'Mutation' } & {
  changeBalloonPosition: { __typename?: 'ChangeBalloonPositionPayload' } & {
    roomUser?: Maybe<{ __typename?: 'RoomUser' } & RoomUserFieldsFragment>
  }
}

export type ActedRoomUserEventSubscriptionVariables = Exact<{
  roomId: Scalars['ID']
}>

export type ActedRoomUserEventSubscription = { __typename?: 'Subscription' } & {
  actedRoomUserEvent?: Maybe<
    | ({ __typename?: 'EnteredPayload' } & {
        roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
      })
    | ({ __typename?: 'ExitedPayload' } & Pick<ExitedPayload, 'userId'>)
    | ({ __typename?: 'MovedPayload' } & {
        roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
      })
    | ({ __typename?: 'SentMassagePayload' } & {
        roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
      })
    | ({ __typename?: 'RemovedLastMessagePayload' } & {
        roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
      })
    | ({ __typename?: 'ChangedBalloonPositionPayload' } & {
        roomUser: { __typename?: 'RoomUser' } & RoomUserFieldsFragment
      })
  >
}

export type PageMessagesFieldFragment = { __typename?: 'Room' } & {
  messages: { __typename?: 'MessageConnection' } & {
    pageInfo: { __typename?: 'PageInfo' } & Pick<
      PageInfo,
      'startCursor' | 'endCursor' | 'hasNextPage' | 'hasPreviousPage'
    >
    nodes: Array<{ __typename?: 'Message' } & MessageFieldsFragment>
  }
}

export type MessageFieldsFragment = { __typename?: 'Message' } & Pick<
  Message,
  'id' | 'body' | 'createdAt'
> & { user: { __typename?: 'User' } & Pick<User, 'id' | 'name' | 'avatarUrl'> }

export type RoomUserFieldsFragment = { __typename?: 'RoomUser' } & Pick<
  RoomUser,
  'id' | 'x' | 'y' | 'balloonPosition'
> & {
    user: { __typename?: 'User' } & Pick<User, 'id' | 'name' | 'avatarUrl'>
    lastMessage?: Maybe<{ __typename?: 'Message' } & MessageFieldsFragment>
  }

export type CommonQueryVariables = Exact<{ [key: string]: never }>

export type CommonQuery = { __typename?: 'Query' } & {
  me: { __typename?: 'User' } & UserFieldsFragment
  onlineUsers: Array<{ __typename?: 'User' } & UserFieldsFragment>
}

export type UserFieldsFragment = { __typename?: 'User' } & Pick<
  User,
  'id' | 'name' | 'avatarUrl'
>

export type ActedUserEventSubscriptionVariables = Exact<{
  [key: string]: never
}>

export type ActedUserEventSubscription = { __typename?: 'Subscription' } & {
  actedUserEvent?: Maybe<
    | ({ __typename?: 'OnlinedPayload' } & {
        user: { __typename?: 'User' } & UserFieldsFragment
      })
    | ({ __typename?: 'OfflinedPayload' } & {
        user: { __typename?: 'User' } & UserFieldsFragment
      })
  >
}

export const RoomFieldsForListFragmentDoc = gql`
  fragment RoomFieldsForList on Room {
    id
    name
    createdAt
    totalUserCount
    totalMessageCount
  }
`
export const MessageFieldsFragmentDoc = gql`
  fragment MessageFields on Message {
    id
    body
    createdAt
    user {
      id
      name
      avatarUrl
    }
  }
`
export const PageMessagesFieldFragmentDoc = gql`
  fragment PageMessagesField on Room {
    messages(last: 10, before: $before) {
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      nodes {
        ...MessageFields
      }
    }
  }
  ${MessageFieldsFragmentDoc}
`
export const RoomUserFieldsFragmentDoc = gql`
  fragment RoomUserFields on RoomUser {
    id
    user {
      id
      name
      avatarUrl
    }
    x
    y
    lastMessage {
      ...MessageFields
    }
    balloonPosition
  }
  ${MessageFieldsFragmentDoc}
`
export const UserFieldsFragmentDoc = gql`
  fragment UserFields on User {
    id
    name
    avatarUrl
  }
`
export const CreateRoomDocument = gql`
  mutation CreateRoom($name: String!, $bgUrl: String!) {
    createRoom(input: { name: $name, bgUrl: $bgUrl }) {
      room {
        id
        name
      }
    }
  }
`
export type CreateRoomMutationFn = Apollo.MutationFunction<
  CreateRoomMutation,
  CreateRoomMutationVariables
>

/**
 * __useCreateRoomMutation__
 *
 * To run a mutation, you first call `useCreateRoomMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateRoomMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createRoomMutation, { data, loading, error }] = useCreateRoomMutation({
 *   variables: {
 *      name: // value for 'name'
 *      bgUrl: // value for 'bgUrl'
 *   },
 * });
 */
export function useCreateRoomMutation(
  baseOptions?: Apollo.MutationHookOptions<
    CreateRoomMutation,
    CreateRoomMutationVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<CreateRoomMutation, CreateRoomMutationVariables>(
    CreateRoomDocument,
    options,
  )
}
export type CreateRoomMutationHookResult = ReturnType<
  typeof useCreateRoomMutation
>
export type CreateRoomMutationResult = Apollo.MutationResult<CreateRoomMutation>
export type CreateRoomMutationOptions = Apollo.BaseMutationOptions<
  CreateRoomMutation,
  CreateRoomMutationVariables
>
export const IndexPageDocument = gql`
  query IndexPage($first: Int!, $after: String) {
    rooms(first: $first, after: $after, orderBy: LATEST) {
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      nodes {
        ...RoomFieldsForList
      }
      edges {
        node {
          ...RoomFieldsForList
        }
      }
      roomCount
    }
  }
  ${RoomFieldsForListFragmentDoc}
`

/**
 * __useIndexPageQuery__
 *
 * To run a query within a React component, call `useIndexPageQuery` and pass it any options that fit your needs.
 * When your component renders, `useIndexPageQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useIndexPageQuery({
 *   variables: {
 *      first: // value for 'first'
 *      after: // value for 'after'
 *   },
 * });
 */
export function useIndexPageQuery(
  baseOptions: Apollo.QueryHookOptions<IndexPageQuery, IndexPageQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<IndexPageQuery, IndexPageQueryVariables>(
    IndexPageDocument,
    options,
  )
}
export function useIndexPageLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    IndexPageQuery,
    IndexPageQueryVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<IndexPageQuery, IndexPageQueryVariables>(
    IndexPageDocument,
    options,
  )
}
export type IndexPageQueryHookResult = ReturnType<typeof useIndexPageQuery>
export type IndexPageLazyQueryHookResult = ReturnType<
  typeof useIndexPageLazyQuery
>
export type IndexPageQueryResult = Apollo.QueryResult<
  IndexPageQuery,
  IndexPageQueryVariables
>
export const RoomPageDocument = gql`
  query RoomPage($roomId: ID!, $before: String) {
    room(id: $roomId) {
      id
      name
      bgColor
      bgUrl
      users {
        ...RoomUserFields
      }
      messages(last: 10, before: $before) {
        pageInfo {
          startCursor
          endCursor
          hasNextPage
          hasPreviousPage
        }
        nodes {
          ...MessageFields
        }
      }
    }
  }
  ${RoomUserFieldsFragmentDoc}
  ${MessageFieldsFragmentDoc}
`

/**
 * __useRoomPageQuery__
 *
 * To run a query within a React component, call `useRoomPageQuery` and pass it any options that fit your needs.
 * When your component renders, `useRoomPageQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useRoomPageQuery({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      before: // value for 'before'
 *   },
 * });
 */
export function useRoomPageQuery(
  baseOptions: Apollo.QueryHookOptions<RoomPageQuery, RoomPageQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<RoomPageQuery, RoomPageQueryVariables>(
    RoomPageDocument,
    options,
  )
}
export function useRoomPageLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    RoomPageQuery,
    RoomPageQueryVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<RoomPageQuery, RoomPageQueryVariables>(
    RoomPageDocument,
    options,
  )
}
export type RoomPageQueryHookResult = ReturnType<typeof useRoomPageQuery>
export type RoomPageLazyQueryHookResult = ReturnType<
  typeof useRoomPageLazyQuery
>
export type RoomPageQueryResult = Apollo.QueryResult<
  RoomPageQuery,
  RoomPageQueryVariables
>
export const MoreRoomMessagesDocument = gql`
  query MoreRoomMessages($roomId: ID!, $before: String) {
    room(id: $roomId) {
      id
      ...PageMessagesField
    }
  }
  ${PageMessagesFieldFragmentDoc}
`

/**
 * __useMoreRoomMessagesQuery__
 *
 * To run a query within a React component, call `useMoreRoomMessagesQuery` and pass it any options that fit your needs.
 * When your component renders, `useMoreRoomMessagesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMoreRoomMessagesQuery({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      before: // value for 'before'
 *   },
 * });
 */
export function useMoreRoomMessagesQuery(
  baseOptions: Apollo.QueryHookOptions<
    MoreRoomMessagesQuery,
    MoreRoomMessagesQueryVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<MoreRoomMessagesQuery, MoreRoomMessagesQueryVariables>(
    MoreRoomMessagesDocument,
    options,
  )
}
export function useMoreRoomMessagesLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    MoreRoomMessagesQuery,
    MoreRoomMessagesQueryVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    MoreRoomMessagesQuery,
    MoreRoomMessagesQueryVariables
  >(MoreRoomMessagesDocument, options)
}
export type MoreRoomMessagesQueryHookResult = ReturnType<
  typeof useMoreRoomMessagesQuery
>
export type MoreRoomMessagesLazyQueryHookResult = ReturnType<
  typeof useMoreRoomMessagesLazyQuery
>
export type MoreRoomMessagesQueryResult = Apollo.QueryResult<
  MoreRoomMessagesQuery,
  MoreRoomMessagesQueryVariables
>
export const SendMessageDocument = gql`
  mutation SendMessage($roomId: ID!, $body: String!) {
    sendMessage(input: { roomID: $roomId, body: $body }) {
      message {
        ...MessageFields
      }
    }
  }
  ${MessageFieldsFragmentDoc}
`
export type SendMessageMutationFn = Apollo.MutationFunction<
  SendMessageMutation,
  SendMessageMutationVariables
>

/**
 * __useSendMessageMutation__
 *
 * To run a mutation, you first call `useSendMessageMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSendMessageMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [sendMessageMutation, { data, loading, error }] = useSendMessageMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      body: // value for 'body'
 *   },
 * });
 */
export function useSendMessageMutation(
  baseOptions?: Apollo.MutationHookOptions<
    SendMessageMutation,
    SendMessageMutationVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<SendMessageMutation, SendMessageMutationVariables>(
    SendMessageDocument,
    options,
  )
}
export type SendMessageMutationHookResult = ReturnType<
  typeof useSendMessageMutation
>
export type SendMessageMutationResult = Apollo.MutationResult<SendMessageMutation>
export type SendMessageMutationOptions = Apollo.BaseMutationOptions<
  SendMessageMutation,
  SendMessageMutationVariables
>
export const MoveDocument = gql`
  mutation Move($roomId: ID!, $x: Int!, $y: Int!) {
    move(input: { roomId: $roomId, x: $x, y: $y }) {
      roomUser {
        ...RoomUserFields
      }
    }
  }
  ${RoomUserFieldsFragmentDoc}
`
export type MoveMutationFn = Apollo.MutationFunction<
  MoveMutation,
  MoveMutationVariables
>

/**
 * __useMoveMutation__
 *
 * To run a mutation, you first call `useMoveMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useMoveMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [moveMutation, { data, loading, error }] = useMoveMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      x: // value for 'x'
 *      y: // value for 'y'
 *   },
 * });
 */
export function useMoveMutation(
  baseOptions?: Apollo.MutationHookOptions<MoveMutation, MoveMutationVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<MoveMutation, MoveMutationVariables>(
    MoveDocument,
    options,
  )
}
export type MoveMutationHookResult = ReturnType<typeof useMoveMutation>
export type MoveMutationResult = Apollo.MutationResult<MoveMutation>
export type MoveMutationOptions = Apollo.BaseMutationOptions<
  MoveMutation,
  MoveMutationVariables
>
export const RemoveLastMessageDocument = gql`
  mutation RemoveLastMessage($roomId: ID!) {
    removeLastMessage(input: { roomId: $roomId }) {
      roomUser {
        ...RoomUserFields
      }
    }
  }
  ${RoomUserFieldsFragmentDoc}
`
export type RemoveLastMessageMutationFn = Apollo.MutationFunction<
  RemoveLastMessageMutation,
  RemoveLastMessageMutationVariables
>

/**
 * __useRemoveLastMessageMutation__
 *
 * To run a mutation, you first call `useRemoveLastMessageMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveLastMessageMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeLastMessageMutation, { data, loading, error }] = useRemoveLastMessageMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useRemoveLastMessageMutation(
  baseOptions?: Apollo.MutationHookOptions<
    RemoveLastMessageMutation,
    RemoveLastMessageMutationVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    RemoveLastMessageMutation,
    RemoveLastMessageMutationVariables
  >(RemoveLastMessageDocument, options)
}
export type RemoveLastMessageMutationHookResult = ReturnType<
  typeof useRemoveLastMessageMutation
>
export type RemoveLastMessageMutationResult = Apollo.MutationResult<RemoveLastMessageMutation>
export type RemoveLastMessageMutationOptions = Apollo.BaseMutationOptions<
  RemoveLastMessageMutation,
  RemoveLastMessageMutationVariables
>
export const ChangeBalloonPositionDocument = gql`
  mutation ChangeBalloonPosition(
    $roomId: ID!
    $balloonPosition: BalloonPosition!
  ) {
    changeBalloonPosition(
      input: { roomId: $roomId, balloonPosition: $balloonPosition }
    ) {
      roomUser {
        ...RoomUserFields
      }
    }
  }
  ${RoomUserFieldsFragmentDoc}
`
export type ChangeBalloonPositionMutationFn = Apollo.MutationFunction<
  ChangeBalloonPositionMutation,
  ChangeBalloonPositionMutationVariables
>

/**
 * __useChangeBalloonPositionMutation__
 *
 * To run a mutation, you first call `useChangeBalloonPositionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangeBalloonPositionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changeBalloonPositionMutation, { data, loading, error }] = useChangeBalloonPositionMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      balloonPosition: // value for 'balloonPosition'
 *   },
 * });
 */
export function useChangeBalloonPositionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    ChangeBalloonPositionMutation,
    ChangeBalloonPositionMutationVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    ChangeBalloonPositionMutation,
    ChangeBalloonPositionMutationVariables
  >(ChangeBalloonPositionDocument, options)
}
export type ChangeBalloonPositionMutationHookResult = ReturnType<
  typeof useChangeBalloonPositionMutation
>
export type ChangeBalloonPositionMutationResult = Apollo.MutationResult<ChangeBalloonPositionMutation>
export type ChangeBalloonPositionMutationOptions = Apollo.BaseMutationOptions<
  ChangeBalloonPositionMutation,
  ChangeBalloonPositionMutationVariables
>
export const ActedRoomUserEventDocument = gql`
  subscription actedRoomUserEvent($roomId: ID!) {
    actedRoomUserEvent(roomId: $roomId) {
      ... on EnteredPayload {
        roomUser {
          ...RoomUserFields
        }
      }
      ... on ExitedPayload {
        userId
      }
      ... on MovedPayload {
        roomUser {
          ...RoomUserFields
        }
      }
      ... on SentMassagePayload {
        roomUser {
          ...RoomUserFields
        }
      }
      ... on RemovedLastMessagePayload {
        roomUser {
          ...RoomUserFields
        }
      }
      ... on ChangedBalloonPositionPayload {
        roomUser {
          ...RoomUserFields
        }
      }
    }
  }
  ${RoomUserFieldsFragmentDoc}
`

/**
 * __useActedRoomUserEventSubscription__
 *
 * To run a query within a React component, call `useActedRoomUserEventSubscription` and pass it any options that fit your needs.
 * When your component renders, `useActedRoomUserEventSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useActedRoomUserEventSubscription({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useActedRoomUserEventSubscription(
  baseOptions: Apollo.SubscriptionHookOptions<
    ActedRoomUserEventSubscription,
    ActedRoomUserEventSubscriptionVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useSubscription<
    ActedRoomUserEventSubscription,
    ActedRoomUserEventSubscriptionVariables
  >(ActedRoomUserEventDocument, options)
}
export type ActedRoomUserEventSubscriptionHookResult = ReturnType<
  typeof useActedRoomUserEventSubscription
>
export type ActedRoomUserEventSubscriptionResult = Apollo.SubscriptionResult<ActedRoomUserEventSubscription>
export const CommonDocument = gql`
  query Common {
    me {
      ...UserFields
    }
    onlineUsers {
      ...UserFields
    }
  }
  ${UserFieldsFragmentDoc}
`

/**
 * __useCommonQuery__
 *
 * To run a query within a React component, call `useCommonQuery` and pass it any options that fit your needs.
 * When your component renders, `useCommonQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCommonQuery({
 *   variables: {
 *   },
 * });
 */
export function useCommonQuery(
  baseOptions?: Apollo.QueryHookOptions<CommonQuery, CommonQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<CommonQuery, CommonQueryVariables>(
    CommonDocument,
    options,
  )
}
export function useCommonLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<CommonQuery, CommonQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<CommonQuery, CommonQueryVariables>(
    CommonDocument,
    options,
  )
}
export type CommonQueryHookResult = ReturnType<typeof useCommonQuery>
export type CommonLazyQueryHookResult = ReturnType<typeof useCommonLazyQuery>
export type CommonQueryResult = Apollo.QueryResult<
  CommonQuery,
  CommonQueryVariables
>
export const ActedUserEventDocument = gql`
  subscription actedUserEvent {
    actedUserEvent {
      ... on OnlinedPayload {
        user {
          ...UserFields
        }
      }
      ... on OfflinedPayload {
        user {
          ...UserFields
        }
      }
    }
  }
  ${UserFieldsFragmentDoc}
`

/**
 * __useActedUserEventSubscription__
 *
 * To run a query within a React component, call `useActedUserEventSubscription` and pass it any options that fit your needs.
 * When your component renders, `useActedUserEventSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useActedUserEventSubscription({
 *   variables: {
 *   },
 * });
 */
export function useActedUserEventSubscription(
  baseOptions?: Apollo.SubscriptionHookOptions<
    ActedUserEventSubscription,
    ActedUserEventSubscriptionVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useSubscription<
    ActedUserEventSubscription,
    ActedUserEventSubscriptionVariables
  >(ActedUserEventDocument, options)
}
export type ActedUserEventSubscriptionHookResult = ReturnType<
  typeof useActedUserEventSubscription
>
export type ActedUserEventSubscriptionResult = Apollo.SubscriptionResult<ActedUserEventSubscription>
