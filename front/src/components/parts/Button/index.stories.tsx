import React from 'react'
import { Story, Meta } from '@storybook/react'
import Button, { ButtonProps } from '.'

export default {
  title: 'parts/Button',
  component: Button,
} as Meta<ButtonProps>

const Template: Story<ButtonProps> = (args) => (
  <div style={{ width: '400px' }}>
    <div className="mb-2">
      <Button {...args} color="gray">
        Gray Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="green">
        Green Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="red">
        Red Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="gray" outline>
        Outline Gray Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="green" outline>
        Outline Green Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="red" outline>
        Outline Red Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} color="gray" fontBold>
        Font Bold Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} fullWidth>
        Full Width Button
      </Button>
    </div>
    <div className="mb-2 h-20">
      <Button {...args} fullHeight>
        Full Height Button
      </Button>
    </div>
    <div className="mb-2">
      <Button {...args} disabled>
        Disabled Button
      </Button>
    </div>
  </div>
)

export const Default = Template.bind({})
Default.args = {
  onClick: () => {},
}
