export const mockImageUrl =
  'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png'
export const SAMPLE_BG_IMAGE = 'https://pbs.twimg.com/media/EVUqmD3U4AABXgv.jpg'
export const DEFAULT_ROOM_BG_COLOR = '#5f9ea0'
export const ROOM_SCREEN_SIZE = {
  // 16:9
  WIDTH: 720,
  HEIGHT: 405,
}

export type RoomBgImage = {
  name: string
  url: string
}

export const ROOM_BG_IMAGES: RoomBgImage[] = [
  {
    name: 'theme1',
    url: 'https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg1.jpg',
  },
  {
    name: 'theme3',
    url: 'https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg3.jpg',
  },
]
