import * as React from 'react'

import { grpc } from 'grpc-web-client'

import * as pb from 'src/proto'

export interface Props {
  children: (state: State) => React.ReactNode
}

export interface State {
  readonly user: pb.User.AsObject | null
}

export class WithUser extends React.Component<Props, State> {
  public readonly state: State = {
    user: null,
  }

  public render() {
    return this.props.children(this.state)
  }

  public componentDidMount() {
    const request = new pb.GetUserRequest()
    grpc.unary(pb.AccountService.GetUser, {
      request,
      host: 'http://localhost:10000',
      onEnd: (res: any) => {
        console.log('GetUser', res)
        const user = res.message.getUser().toObject()
        this.setState({ user })
      }
    })
  }
}

export default WithUser
