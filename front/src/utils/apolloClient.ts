import {
  NormalizedCacheObject,
  ApolloClient,
  HttpLink,
  InMemoryCache,
  split,
  Reference,
} from '@apollo/client'
import { WebSocketLink } from '@apollo/client/link/ws'
import { getMainDefinition } from '@apollo/client/utilities'

const getHttpLink = () => {
  const uri = process.browser
    ? 'http://localhost:8080/query'
    : 'http://api:8080/query'

  return new HttpLink({ uri, credentials: 'include' })
}

const getWsLink = () => {
  return new WebSocketLink({
    uri: 'ws://localhost:8080/query',
    options: {
      reconnect: true,
      lazy: true,
    },
  })
}

const makeCache = () => {
  return new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          isLoggedIn: {
            read() {
              return true
            },
          },
        },
      },
      Message: {
        fields: {
          createdAt: {
            // MEMO: https://github.com/apollographql/apollo-client/issues/585
            // https://stackoverflow.com/questions/66085887/how-can-you-retrieve-a-date-field-in-apollo-client-from-a-query-as-a-date-and-no
            read(existing) {
              return new Date(existing)
            },
          },
        },
      },
      Room: {
        fields: {
          messages: {
            keyArgs: false,
            merge(existing, incoming, { readField }) {
              console.log({ existing, incoming })
              if (!incoming) return existing
              if (!existing) return incoming
              const { nodes, ...rest } = incoming
              const result = rest

              const merged: Reference[] = [...nodes, ...existing.nodes]
              const filteredDup: Reference[] = Array.from(
                merged
                  .reduce(
                    (map, current) =>
                      map.set(readField('id', current), current),
                    new Map(),
                  )
                  .values(),
              )
              const sorted = filteredDup.sort((a, b) => {
                const aCreatedAt = readField<Date>('createdAt', a)
                const bCreatedAt = readField<Date>('createdAt', b)
                return (
                  (aCreatedAt?.getTime() || 0) - (bCreatedAt?.getTime() || 0)
                )
              })

              result.nodes = sorted
              return result
            },
          },
        },
      },
    },
  })
}

const createLinks = () => {
  const httpLink = getHttpLink()

  if (!process.browser) return httpLink

  const links = split(
    ({ query }) => {
      const definition = getMainDefinition(query)
      return (
        definition.kind === 'OperationDefinition' &&
        definition.operation === 'subscription'
      )
    },
    getWsLink(),
    httpLink,
  )

  return links
}

let apolloClient: ApolloClient<NormalizedCacheObject> | null = null

export const createApolloClient = () => {
  if (apolloClient) {
    return apolloClient
  }

  apolloClient = new ApolloClient({
    link: createLinks(),
    cache: makeCache(),
    connectToDevTools: true, // TODO: import environment
  })

  return apolloClient
}
