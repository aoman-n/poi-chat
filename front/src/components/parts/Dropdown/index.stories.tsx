import React from 'react'
import { Story, Meta } from '@storybook/react'
import Dropdown from '.'

export default {
  title: 'parts/Dropdown',
  component: Dropdown,
} as Meta

const Template: Story = () => (
  <div className="mx-24 my-14">
    <Dropdown
      button={
        <button className="px-6 py-4 bg-blue-400 text-white">click</button>
      }
      leftPx={-100}
    >
      <div className="px-14 py-8">
        <div>hoge</div>
        <div>miso</div>
        <div>huga</div>
      </div>
    </Dropdown>
  </div>
)

export const Default = Template.bind({})
