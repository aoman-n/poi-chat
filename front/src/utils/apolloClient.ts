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
      Room: {
        fields: {
          messages: {
            keyArgs: false,
            merge(existing, incoming, { readField }) {
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
                const aId = readField('id', a)
                const bId = readField('id', b)

                if (typeof aId === 'string' && typeof bId === 'string') {
                  if (aId > bId) return 1
                }
                return -1
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

  return split(
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
