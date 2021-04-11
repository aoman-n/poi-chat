const devConfig = {
  apiBaseUrl: 'http://localhost:8080',
}

const prdConfig: typeof devConfig = {
  apiBaseUrl: 'http://localhost:8080',
}

const isDev = process.env.NODE_ENV === 'development'

export default isDev ? devConfig : prdConfig
