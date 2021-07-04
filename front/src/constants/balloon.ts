import { ValueOf } from '@/utils/types'
import { BalloonPosition as GraphBalloonPosition } from '@/graphql'

export const BALLOON_POSITIONS = {
  TOP_RIGHT: 'TOP_RIGHT',
  TOP_LEFT: 'TOP_LEFT',
  BOTTOM_RIGHT: 'BOTTOM_RIGHT',
  BOTTOM_LEFT: 'BOTTOM_LEFT',
} as const

export type BalloonPosition = ValueOf<typeof BALLOON_POSITIONS>

export const convertToFrontBalloonPos = (
  graphPosision: GraphBalloonPosition,
) => {
  switch (graphPosision) {
    case GraphBalloonPosition.TopRight:
      return BALLOON_POSITIONS.TOP_RIGHT
    case GraphBalloonPosition.TopLeft:
      return BALLOON_POSITIONS.TOP_LEFT
    case GraphBalloonPosition.BottomRight:
      return BALLOON_POSITIONS.BOTTOM_RIGHT
    case GraphBalloonPosition.BottomLeft:
      return BALLOON_POSITIONS.BOTTOM_LEFT
  }
}

export const convertToGraphBalloonPos = (position: BalloonPosition) => {
  switch (position) {
    case 'TOP_LEFT':
      return GraphBalloonPosition.TopLeft
    case 'TOP_RIGHT':
      return GraphBalloonPosition.TopRight
    case 'BOTTOM_LEFT':
      return GraphBalloonPosition.BottomLeft
    case 'BOTTOM_RIGHT':
      return GraphBalloonPosition.BottomRight
  }
}
