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

export type OfflineUserStatus = {
  __typename?: 'OfflineUserStatus'
  id: Scalars['ID']
}

export type OnlineUserStatus = {
  __typename?: 'OnlineUserStatus'
  id: Scalars['ID']
  displayName: Scalars['String']
  avatarUrl: Scalars['String']
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
  onlineUsers: Array<OnlineUserStatus>
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
  joinRoom: User
  subUserEvent: UserEvent
  changedUserStatus: UserStatus
}

export type SubscriptionSubMessageArgs = {
  roomId: Scalars['ID']
}

export type SubscriptionJoinRoomArgs = {
  roomID: Scalars['ID']
}

export type SubscriptionSubUserEventArgs = {
  roomId: Scalars['ID']
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

export type UserStatus = OnlineUserStatus | OfflineUserStatus

export type MessageFieldsFragment = { __typename?: 'Message' } & Pick<
  Message,
  'id' | 'userId' | 'userName' | 'userAvatarUrl' | 'body' | 'createdAt'
>

export type RoomDetailFragment = { __typename?: 'Query' } & {
  roomDetail: { __typename?: 'RoomDetail' } & Pick<
    RoomDetail,
    'id' | 'name'
  > & {
      users: Array<
        { __typename?: 'User' } & Pick<
          User,
          'id' | 'displayName' | 'avatarUrl' | 'x' | 'y'
        >
      >
      messages: { __typename?: 'MessageConnection' } & {
        nodes: Array<Maybe<{ __typename?: 'Message' } & MessageFieldsFragment>>
      }
    }
}

export type UserEventSubscriptionVariables = Exact<{
  roomId: Scalars['ID']
}>

export type UserEventSubscription = { __typename?: 'Subscription' } & {
  subUserEvent:
    | ({ __typename: 'MovedUser' } & Pick<MovedUser, 'id' | 'x' | 'y'>)
    | ({ __typename: 'ExitedUser' } & Pick<ExitedUser, 'id'>)
    | ({ __typename: 'JoinedUser' } & Pick<
        JoinedUser,
        'id' | 'displayName' | 'avatarUrl' | 'x' | 'y'
      >)
}

export type MessageSubscriptionVariables = Exact<{
  roomId: Scalars['ID']
}>

export type MessageSubscription = { __typename?: 'Subscription' } & {
  subMessage: { __typename?: 'Message' } & MessageFieldsFragment
}

export type SendMessageMutationVariables = Exact<{
  roomId: Scalars['ID']
  body: Scalars['String']
}>

export type SendMessageMutation = { __typename?: 'Mutation' } & {
  sendMessage: { __typename?: 'Message' } & Pick<
    Message,
    'id' | 'userId' | 'userName' | 'userAvatarUrl' | 'body' | 'createdAt'
  >
}

export type AuthQueryVariables = Exact<{ [key: string]: never }>

export type AuthQuery = { __typename?: 'Query' } & {
  me: { __typename?: 'Me' } & Pick<Me, 'id' | 'displayName' | 'avatarUrl'>
  onlineUsers: Array<
    { __typename?: 'OnlineUserStatus' } & Pick<
      OnlineUserStatus,
      'id' | 'displayName' | 'avatarUrl'
    >
  >
}

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

export type RoomsQueryVariables = Exact<{
  roomId: Scalars['ID']
}>

export type RoomsQuery = { __typename?: 'Query' } & RoomDetailFragment

export type JoinRoomSubscriptionVariables = Exact<{
  roomId: Scalars['ID']
}>

export type JoinRoomSubscription = { __typename?: 'Subscription' } & {
  joinRoom: { __typename?: 'User' } & Pick<
    User,
    'id' | 'displayName' | 'avatarUrl' | 'x' | 'y'
  >
}

export type MoveMutationVariables = Exact<{
  roomId: Scalars['ID']
  x: Scalars['Int']
  y: Scalars['Int']
}>

export type MoveMutation = { __typename?: 'Mutation' } & {
  move: { __typename?: 'MovedUser' } & Pick<MovedUser, 'id' | 'x' | 'y'>
}

export const MessageFieldsFragmentDoc = gql`
  fragment MessageFields on Message {
    id
    userId
    userName
    userAvatarUrl
    body
    createdAt
  }
`
export const RoomDetailFragmentDoc = gql`
  fragment RoomDetail on Query {
    roomDetail(id: $roomId) {
      id
      name
      users {
        id
        displayName
        avatarUrl
        x
        y
      }
      messages(last: 20) {
        nodes {
          ...MessageFields
        }
      }
    }
  }
  ${MessageFieldsFragmentDoc}
`
export const UserEventDocument = gql`
  subscription UserEvent($roomId: ID!) {
    subUserEvent(roomId: $roomId) {
      __typename
      ... on MovedUser {
        id
        x
        y
      }
      ... on ExitedUser {
        id
      }
      ... on JoinedUser {
        id
        displayName
        avatarUrl
        x
        y
      }
    }
  }
`

/**
 * __useUserEventSubscription__
 *
 * To run a query within a React component, call `useUserEventSubscription` and pass it any options that fit your needs.
 * When your component renders, `useUserEventSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserEventSubscription({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useUserEventSubscription(
  baseOptions: Apollo.SubscriptionHookOptions<
    UserEventSubscription,
    UserEventSubscriptionVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useSubscription<
    UserEventSubscription,
    UserEventSubscriptionVariables
  >(UserEventDocument, options)
}
export type UserEventSubscriptionHookResult = ReturnType<
  typeof useUserEventSubscription
>
export type UserEventSubscriptionResult = Apollo.SubscriptionResult<UserEventSubscription>
export const MessageDocument = gql`
  subscription Message($roomId: ID!) {
    subMessage(roomId: $roomId) {
      ...MessageFields
    }
  }
  ${MessageFieldsFragmentDoc}
`

/**
 * __useMessageSubscription__
 *
 * To run a query within a React component, call `useMessageSubscription` and pass it any options that fit your needs.
 * When your component renders, `useMessageSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMessageSubscription({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useMessageSubscription(
  baseOptions: Apollo.SubscriptionHookOptions<
    MessageSubscription,
    MessageSubscriptionVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useSubscription<
    MessageSubscription,
    MessageSubscriptionVariables
  >(MessageDocument, options)
}
export type MessageSubscriptionHookResult = ReturnType<
  typeof useMessageSubscription
>
export type MessageSubscriptionResult = Apollo.SubscriptionResult<MessageSubscription>
export const SendMessageDocument = gql`
  mutation SendMessage($roomId: ID!, $body: String!) {
    sendMessage(input: { roomID: $roomId, body: $body }) {
      id
      userId
      userName
      userAvatarUrl
      body
      createdAt
    }
  }
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
export const AuthDocument = gql`
  query Auth {
    me {
      id
      displayName
      avatarUrl
    }
    onlineUsers {
      id
      displayName
      avatarUrl
    }
  }
`

/**
 * __useAuthQuery__
 *
 * To run a query within a React component, call `useAuthQuery` and pass it any options that fit your needs.
 * When your component renders, `useAuthQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAuthQuery({
 *   variables: {
 *   },
 * });
 */
export function useAuthQuery(
  baseOptions?: Apollo.QueryHookOptions<AuthQuery, AuthQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<AuthQuery, AuthQueryVariables>(AuthDocument, options)
}
export function useAuthLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<AuthQuery, AuthQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<AuthQuery, AuthQueryVariables>(
    AuthDocument,
    options,
  )
}
export type AuthQueryHookResult = ReturnType<typeof useAuthQuery>
export type AuthLazyQueryHookResult = ReturnType<typeof useAuthLazyQuery>
export type AuthQueryResult = Apollo.QueryResult<AuthQuery, AuthQueryVariables>
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
export const RoomsDocument = gql`
  query Rooms($roomId: ID!) {
    ...RoomDetail
  }
  ${RoomDetailFragmentDoc}
`

/**
 * __useRoomsQuery__
 *
 * To run a query within a React component, call `useRoomsQuery` and pass it any options that fit your needs.
 * When your component renders, `useRoomsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useRoomsQuery({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useRoomsQuery(
  baseOptions: Apollo.QueryHookOptions<RoomsQuery, RoomsQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<RoomsQuery, RoomsQueryVariables>(
    RoomsDocument,
    options,
  )
}
export function useRoomsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<RoomsQuery, RoomsQueryVariables>,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<RoomsQuery, RoomsQueryVariables>(
    RoomsDocument,
    options,
  )
}
export type RoomsQueryHookResult = ReturnType<typeof useRoomsQuery>
export type RoomsLazyQueryHookResult = ReturnType<typeof useRoomsLazyQuery>
export type RoomsQueryResult = Apollo.QueryResult<
  RoomsQuery,
  RoomsQueryVariables
>
export const JoinRoomDocument = gql`
  subscription JoinRoom($roomId: ID!) {
    joinRoom(roomID: $roomId) {
      id
      displayName
      avatarUrl
      x
      y
    }
  }
`

/**
 * __useJoinRoomSubscription__
 *
 * To run a query within a React component, call `useJoinRoomSubscription` and pass it any options that fit your needs.
 * When your component renders, `useJoinRoomSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useJoinRoomSubscription({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useJoinRoomSubscription(
  baseOptions: Apollo.SubscriptionHookOptions<
    JoinRoomSubscription,
    JoinRoomSubscriptionVariables
  >,
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useSubscription<
    JoinRoomSubscription,
    JoinRoomSubscriptionVariables
  >(JoinRoomDocument, options)
}
export type JoinRoomSubscriptionHookResult = ReturnType<
  typeof useJoinRoomSubscription
>
export type JoinRoomSubscriptionResult = Apollo.SubscriptionResult<JoinRoomSubscription>
export const MoveDocument = gql`
  mutation Move($roomId: ID!, $x: Int!, $y: Int!) {
    move(input: { roomId: $roomId, x: $x, y: $y }) {
      id
      x
      y
    }
  }
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
