import React from 'react'
import { Story, Meta } from '@storybook/react'
import MainTemplate, { MainTemplateProps } from '.'

export default {
  title: 'templates/Main',
  component: MainTemplate,
} as Meta

const Template: Story<MainTemplateProps> = (args) => <MainTemplate {...args} />

export const Default = Template.bind({})
Default.args = {
  MainComponent: <div className="bg-green-300 h-full">メインエリア</div>,
  MyProfileComponent: <div className="bg-red-300 h-72">プロフィール</div>,
  OnlineUserListComponent: (
    <div className="h-80">
      <div>ユーザー一覧</div>
      {[...Array(30)].map((_, i) => (
        <div key={i}>{'ユーザー' + i}</div>
      ))}
    </div>
  ),
}
