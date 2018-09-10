import * as React from 'react'

import * as pb from 'src/proto'

export interface Props {
  users: pb.User.AsObject[]
}

export function UserList({ users }: Props) {
  return (
    <ul>
      {users.map(u => (
        <li key={u.id}>{u.id}: {u.name}</li>
      ))}
    </ul>
  )
}

export default UserList