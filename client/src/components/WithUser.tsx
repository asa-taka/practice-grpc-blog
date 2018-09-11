import * as React from 'react'

import { grpc } from 'grpc-web-client'

import * as pb from 'src/proto'

export interface Props {
  children: (state: State) => React.ReactNode
}

export interface State {
  readonly user: pb.User.AsObject
}

export class WithUser extends React.Component<Props, State> {
  public readonly state: State = {
    user: {
      id: 1,
      name: 'asa-taka',
      email: 'asa-taka@example.com',
      status: pb.User.Status.STATUS_ACTIVE,
    }
  }

  public render() {
    return this.props.children(this.state)
  }

  public componentDidMount() {
    const request = new pb.GetUserRequest()
    grpc.unary(pb.AccountService.GetUser, {
      request,
      host: 'http://localhost:10000',
      onEnd: res => {
        console.log(res, res.message)
      }
    })
  }
}

export default WithUser
