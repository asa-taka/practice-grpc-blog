import * as React from 'react'

import { grpc } from 'grpc-web-client'

import * as pb from 'src/proto'

export interface Props {
  children: (state: State) => React.ReactNode
}

export interface State {
  readonly users: pb.User.AsObject[]
}

export class WithUsers extends React.Component<Props, State> {
  public readonly state: State = {
    users: [
      { id: 1, 'name': 'asa-taka', email: 'asa-taka@example.com' }
    ],
  }

  public render() {
    return this.props.children(this.state)
  }

  public componentDidMount() {
    const request = new pb.QueryUsersRequest()
    grpc.unary(pb.AccountService.QueryUsers, {
      request,
      host: 'http://localhost:10000',
      onEnd: res => {
        console.log(res, res.message)
      }
    })
  }
}

export default WithUsers
