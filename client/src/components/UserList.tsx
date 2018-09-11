import * as React from 'react'

import * as pb from 'src/proto'

import UserCard from './UserCard'

export interface Props {
  users: pb.User.AsObject[]
}

export function UserList({ users }: Props) {
  return (
    <ul>
      {users.map(u => (
        <UserCard key={u.id} user={u} />
      ))}
    </ul>
  )
}

export default UserList