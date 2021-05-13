# データ設計メモ

- Type
  - GlobalUserStatus: オンライン/オフラインの管理。ログインユーザーすべてに配信。
    - Online
    - Offline
  - RoomUserEvent: 入室/退室/位置移動/発話etc...を管理。同ルーム内のみに配信。
    - Join
    - Exit
    - Move
    - Message
- Subscription
  - changedGlobalUserStatus: GlobalUserStatus!
  - actedRoomUser: RoomUserEvent!

## すべてのオンラインユーザーに配信するデータ

### オンラインユーザー情報

userStatusとしてRedisへ保存する。
有効期限は`60000ms`。
`del` or `expired` イベントを受け取ったらオフライン状態とする。

```bash
# key
globalUserStatuses:<userId>
```

```json
{
  "id": "xxx",
  "name": "xxx",
  "avatarUrl": "xxx"
}
```

## ルーム内ユーザーへ配信するデータ

### ルームにいるユーザー情報

Redisに保存する。

```bash
# key
roomUsers:<roomId>:<userId>
# set
set roomUsers:100:1
# sadd
sadd roomUsers:100:index
# smembers
sadd roomUsers:100:index
```

```bash
127.0.0.1:6379> set roomUsers:100:1 aoba EX 1000
OK
127.0.0.1:6379> set roomUsers:100:2 hiroshi EX 1000
OK
127.0.0.1:6379> sadd roomUsers:100:index roomUsers:100:1
(integer) 1
127.0.0.1:6379> sadd roomUsers:100:index roomUsers:100:2
(integer) 1
127.0.0.1:6379> smembers roomUsers:100:index
1) "roomUsers:100:1"
2) "roomUsers:100:2"
127.0.0.1:6379> scard roomUsers:100:index
(integer) 2
127.0.0.1:6379> srem roomUsers:100:index roomUsers:100:1
(integer) 1
127.0.0.1:6379> smembers roomUsers:100:index
1) "roomUsers:100:2"
```

```json
{
  "id": "xxx",
  "name": "xxx",
  "avatarUrl": "xxx",
  "posX": 100,
  "posY": 100
}
```

## GraphQL Schema

```graphql
type MovePayload {}
type SendMessagePayload {}

extend type Mutation {
  move(roomId: ID!, input: MoveInput!): MovePayload!
  sendMessage(roomId: ID!, input: SendMessageInput!): SendMessagePayload!
}

extend type Subscription {
  # 接続でオンライン状態/切断でオフライン状態にする
  keepOnline: Boolean!
  actedGlobalUserEvent: GlobalUserEvent!
  # roomに入室している状態も管理する
  actedRoomUserEvent(roomId: ID!): RoomUserEvent!
}

input MoveInput {
  x: Int!
  y: Int!
}
input SendMessageInput {
  message: String!
}

union GlobalUserEvent = Onlined | Offlined
type Onlined {
  userId: ID!
  name: String!
  avatarUrl: String!
}
type Offlined {
  userId: ID!
}

union RoomUserEvent = Joined | Exited | Moved | SendedMassage
type Joined {
  userId: ID!
  name: String!
  avatarUrl: String!
  x: Int!
  y: Int!
}
type Exited {
  userId: ID!
}
type Moved {
  userId: ID!
  x: Int!
  y: Int!
}
type SendedMassage {
  userId: ID!
  message: String!
}
```
