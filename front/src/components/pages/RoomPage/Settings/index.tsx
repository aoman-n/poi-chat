import React from 'react'
import Button from '@/components/parts/Button'
import { BalloonPosition } from '@/constants'

export type BalloonState = {
  hasBalloon: boolean
  position: BalloonPosition | null
}

export type SettingsProps = {
  handleChangeBalloonPos: (pos: BalloonPosition) => void
  handleRemoveBalloon: () => void
  balloonState: BalloonState
}

const disabledBalloonPosButton = (
  balloonState: BalloonState,
  buttonType: BalloonPosition,
) => {
  return !balloonState.hasBalloon || balloonState.position === buttonType
}

const Settings: React.FC<SettingsProps> = ({
  handleChangeBalloonPos,
  handleRemoveBalloon,
  balloonState,
}) => {
  return (
    <div className="pt-4 flex">
      <div>
        <h4 className="mb-2">吹き出し位置変更</h4>
        <div className="space-x-2 pb-2">
          <Button
            onClick={() => handleChangeBalloonPos('TOP_LEFT')}
            disabled={disabledBalloonPosButton(balloonState, 'TOP_LEFT')}
          >
            ↖左上
          </Button>
          <Button
            onClick={() => handleChangeBalloonPos('TOP_RIGHT')}
            disabled={disabledBalloonPosButton(balloonState, 'TOP_RIGHT')}
          >
            右上↗
          </Button>
        </div>
        <div className="space-x-2">
          <Button
            onClick={() => handleChangeBalloonPos('BOTTOM_LEFT')}
            disabled={disabledBalloonPosButton(balloonState, 'BOTTOM_LEFT')}
          >
            ↙左下
          </Button>
          <Button
            onClick={() => handleChangeBalloonPos('BOTTOM_RIGHT')}
            disabled={disabledBalloonPosButton(balloonState, 'BOTTOM_RIGHT')}
          >
            右下↘
          </Button>
        </div>
      </div>
      <div className="ml-auto">
        <Button
          onClick={handleRemoveBalloon}
          color="red"
          disabled={!balloonState.hasBalloon}
        >
          吹き出しを消す
        </Button>
      </div>
    </div>
  )
}

export default Settings
