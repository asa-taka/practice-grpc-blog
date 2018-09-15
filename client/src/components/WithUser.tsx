import * as React from 'react'

import { grpc } from 'grpc-web-client'

import * as pb from 'src/proto'
import config from 'src/config'

import { RequestState } from './RequestIndicator'

export interface Props {
  children: (data: pb.User.AsObject) => React.ReactNode
}

export type State = RequestState<pb.User.AsObject>

export class WithUser extends React.Component<Props, State> {

  public readonly state: State = {
    status: 'loading',
    data: null,
    error: null
  }

  public render() {
    const { status, data, error } = this.state
    if (!data) {
      return null
    }
    return this.props.children(data)
  }

  public componentDidMount() {
    const request = new pb.GetUserRequest()
    request.setId(1)
    grpc.unary(pb.AccountService.GetUser, {
      request,
      host: config.grpc,

      // NOTE: Type vanished!
      onEnd: (res: any) => {
        console.log('GetUser', res)
        const data = res.message.getUser().toObject()
        this.setState({ data, status: 'success' })
      }
    })
  }
}

export default WithUser
