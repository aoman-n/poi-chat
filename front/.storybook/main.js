const path = require("path")

module.exports = {
  "stories": [
    "../src/**/*.stories.mdx",
    "../src/**/*.stories.@(js|jsx|ts|tsx)"
  ],
  "addons": [
    "@storybook/addon-links",
    "@storybook/addon-essentials"
  ],
  // @see: https://github.com/storybookjs/storybook/issues/15067#issuecomment-851959296
  // solution for typescript v4.3.2
  "typescript": {
    "reactDocgen": "none"
  },
  "webpackFinal": async (config) => {
    config.resolve.alias['@'] = path.resolve(__dirname, '../src')

    config.module.rules.push({
      test: /\.scss$/,
      use: [
        'style-loader',
        {
          loader: 'css-loader',
          options: {
            importLoaders: 1, // 1 => postcss-loader
            modules: {
              localIdentName: '[local]___[hash:base64:2]',
            },
          },
        },
        'sass-loader',
      ],
    })

    return config
  }
}