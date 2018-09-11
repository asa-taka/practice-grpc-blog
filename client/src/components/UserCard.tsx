import * as React from 'react'

import * as pb from 'src/proto'

export interface Props {
  user: pb.User.AsObject
}

export function UserCard({ user }: Props) {
  return (
    <div className="card">
      <div className="card-title">
        {user.name}
      </div>
      <div className="card-body">
        id: {user.id}, {user.email}
      </div>
    </div>
  )
}

export default UserCard