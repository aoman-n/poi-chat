import React from 'react'
import { Story, Meta } from '@storybook/react'
import EntranceTemplate, { EntranceProps } from '.'

export default {
  title: 'templates/Entrance',
  component: EntranceTemplate,
  argTypes: {
    onClick: { action: 'clicked' },
  },
  args: {
    text: { control: 'string' },
  },
} as Meta

const Template: Story<EntranceProps> = (args) => <EntranceTemplate {...args} />

export const Default = Template.bind({})
Default.args = {
  noop: 'ログイン',
}
