import {
  NormalizedCacheObject,
  ApolloClient,
  HttpLink,
  InMemoryCache,
  split,
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

              const isWriteQuery =
                nodes[0] &&
                existing.nodes[0] &&
                readField('id', nodes[0]) === readField('id', existing.nodes[0])

              if (isWriteQuery) {
                result.nodes = [...existing.nodes, nodes[nodes.length - 1]]
              } else {
                result.nodes = [...nodes, ...existing.nodes]
              }

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
