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
    users: [],
  }

  public render() {
    return this.props.children(this.state)
  }

  public componentDidMount() {
    const request = new pb.QueryUsersRequest()
    grpc.unary(pb.AccountService.QueryUsers, {
      request,
      host: 'http://localhost:10000',

      // NOTE: Type vanished!
      onEnd: (res: any) => {
        console.log('QueryUsers', res)
        const users = res.message.getUsersList().map((u: any) => u.toObject())
        this.setState({ users })
      }
    })
  }
}

export default WithUsers
