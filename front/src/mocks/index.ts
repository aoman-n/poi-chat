import { UserInfo } from '../painter/user'
import { OnlineUser } from '@/components/organisms/OnlineUserList'
import { Room } from '@/components/organisms/RoomList'

export const mockUsers: UserInfo[] = [
  {
    id: '1',
    avatarUrl:
      'https://pbs.twimg.com/profile_images/685155144363737088/wJtJ2OlA_400x400.jpg',
    currentX: 10,
    currentY: 10,
  },
  {
    id: '2',
    avatarUrl:
      'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png',
    currentX: 100,
    currentY: 100,
  },
  {
    id: '3',
    avatarUrl:
      'https://avatars.githubusercontent.com/u/16658556?s=400&u=d90077a02b620f83ac0876cfe0b15bd696c415ec&v=4',
    currentX: 200,
    currentY: 200,
  },
]

export const mockOnlineUsers: OnlineUser[] = [
  {
    id: '1',
    name: 'サンプルユーザー1',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '2',
    name: 'サンプルユーザー2',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '3',
    name: 'サンプルユーザー3',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
  {
    id: '4',
    name: 'サンプルユーザー4',
    avatarUrl:
      'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
  },
]

export const mockRooms: Room[] = [
  {
    id: '1',
    name: 'サンプルチャットルーム1',
    userCount: 8,
  },
  {
    id: '2',
    name: 'サンプルチャットルーム2',
    userCount: 10,
  },
  {
    id: '3',
    name: 'サンプルチャットルーム3',
    userCount: 3,
  },
  {
    id: '4',
    name: 'サンプルチャットルーム4',
    userCount: 60,
  },
]
