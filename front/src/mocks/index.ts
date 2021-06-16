import { OnlineUser } from '@/components/organisms/OnlineUserList'
import { Room } from '@/components/organisms/RoomList'
import { PlaygroundProps } from '@/components/organisms/Playground/presentation'
import { RoomFragment, BalloonPosition } from '@/graphql'

export const mockUser = {
  id: '1',
  avatarUrl:
    'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png',
}

export const mockUsers: RoomFragment['room']['users'] = [
  {
    id: '1',
    name: 'ユーザー1',
    avatarUrl:
      'https://pbs.twimg.com/profile_images/685155144363737088/wJtJ2OlA_400x400.jpg',
    x: 60,
    y: 60,
    lastMessage: {
      id: '1',
      userId: '1',
      userName: '名無しさん',
      userAvatarUrl: '',
      body: 'こんにちは',
      createdAt: '2020-09-07T15:31:07Z',
    },
    balloonPosition: BalloonPosition.BottomLeft,
  },
  {
    id: '2',
    name: 'ユーザー2',
    avatarUrl:
      'https://pbs.twimg.com/profile_images/1130684542732230656/pW77OgPS_400x400.png',
    x: 300,
    y: 300,
    lastMessage: {
      id: '2',
      userId: '2',
      userName: 'はるまきくん',
      userAvatarUrl: '',
      body: 'こんにちは2',
      createdAt: '2020-09-07T15:31:07Z',
    },
    balloonPosition: BalloonPosition.BottomRight,
  },
  {
    id: '3',
    name: 'ユーザー3',
    avatarUrl:
      'https://avatars.githubusercontent.com/u/16658556?s=400&u=d90077a02b620f83ac0876cfe0b15bd696c415ec&v=4',
    x: 500,
    y: 200,
    lastMessage: {
      id: '3',
      userId: '3',
      userName: 'からあげさん',
      userAvatarUrl: '',
      body: 'こんにちは3',
      createdAt: '2020-09-07T15:31:07Z',
    },
    balloonPosition: BalloonPosition.TopLeft,
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
    id: 'Room:1',
    name: 'サンプルチャットルーム1',
    createdAt: '',
    totalUserCount: 8,
    totalMessageCount: 100,
  },
  {
    id: 'Room:2',
    name: 'サンプルチャットルーム2',
    createdAt: '',
    totalUserCount: 10,
    totalMessageCount: 100,
  },
  {
    id: 'Room:3',
    name: 'サンプルチャットルーム3',
    createdAt: '',
    totalUserCount: 3,
    totalMessageCount: 100,
  },
  {
    id: 'Room:4',
    name:
      'サンプルチャットルーム4サンプルチャットルーム4サンプルチャットルーム4サンプルチャットルーム4',
    createdAt: '',
    totalUserCount: 60,
    totalMessageCount: 100,
  },
]

export const mockMessages: PlaygroundProps['messages'] = [
  {
    id: '1',
    userName: '名無しさん',
    userId: '1',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'こんな部屋おちつかないわ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '2',
    userName: 'とりポンタ',
    userId: '2',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'ザリガニブームは中国のしわざか？',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '3',
    userName: '岡田毅',
    userId: '3',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '昔は第一ホテルのスイートもそんな値段だったな',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '4',
    userName: '名無しさん',
    userId: '1',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'とりポンタ そんなこと言ったら餃子もラーメンも食べれんよ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '5',
    userName: '名無しさん',
    userId: '5',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '未だに部屋２８℃なんだが・・',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '6',
    userName: '名無しさん',
    userId: '6',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​餃子&ラーメンは最高よ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '7',
    userName: 'デレデレ',
    userId: '7',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​戦後の食糧難時代の団塊世代、ザリガニ大好物やった。',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '8',
    userName: 'TS',
    userId: '8',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​腹減ってきたな',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '9',
    userName: 'Taka Take',
    userId: '9',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​安く仕入れて高く売る。こいつらはそれが命だ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '10',
    userName: '佐倉さおり',
    userId: '10',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​24時間以内にちゃんと石鹸で洗えばうつらんよ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '11',
    userName: 'Taka Take',
    userId: '11',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '高い金払わせられてんだぞ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '12',
    userName: 'Taka Take',
    userId: '12',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '高い金払わせられてんだぞ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '13',
    userName: 'Lエネ',
    userId: '13',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '​茶そば、大好き',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '14',
    userName: '大根玉子',
    userId: '14',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'デパ地下',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '15',
    userName: 'Lエネ',
    userId: '15',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '銭がねーんだよ！舐めてんのか！',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '16',
    userName: '日共マミさんの非凡な日常',
    userId: '16',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '知ってるけど、ザリガニとどういう関係が？',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '17',
    userName: 'Lエネ',
    userId: '17',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'ANN嘘くさいよね。こんな高いの買えないでしょ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '18',
    userName: '岡田毅',
    userId: '18',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '駅弁が一番だな、弁当は',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '19',
    userName: '名無しさん',
    userId: '19',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '傷口あったら12時間らしいｗ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '20',
    userName: '名無しさん',
    userId: '20',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body:
      '顔あったら食べにくいです。ホラー映画。今まで悪いことやってきた憎まれキャラの最期の叫びみたい',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '21',
    userName: 'Lエネ',
    userId: '21',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '銭がねーんだよ！舐めてんのか！',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '22',
    userName: '日共マミさんの非凡な日常',
    userId: '22',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '知ってるけど、ザリガニとどういう関係が？',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '23',
    userName: 'Lエネ',
    userId: '23',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'ANN嘘くさいよね。こんな高いの買えないでしょ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '24',
    userName: '岡田毅',
    userId: '24',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '駅弁が一番だな、弁当は',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '25',
    userName: '名無しさん',
    userId: '25',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body:
      '顔あったら食べにくいです。ホラー映画。今まで悪いことやってきた憎まれキャラの最期の叫びみたい',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '26',
    userName: '手嶌。',
    userId: '26',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: 'いよいよ牛とウニ',
    createdAt: '2020-09-07T15:31:07Z',
  },
  {
    id: '27',
    userName: '名無しさん',
    userId: '27',
    userAvatarUrl: 'https://img.icons8.com/cotton/2x/person-male.png',
    body: '傷口あったら12時間らしいｗ',
    createdAt: '2020-09-07T15:31:07Z',
  },
]
