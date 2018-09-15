import * as React from 'react'

import { grpc } from 'grpc-web-client'

import * as pb from 'src/proto'
import config from 'src/config'

import { RequestState } from './RequestIndicator'

export interface Props {
  children: (data: pb.User.AsObject[]) => React.ReactNode
}

export type State = RequestState<pb.User.AsObject[]>

export class WithUsers extends React.Component<Props, State> {

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
    const request = new pb.QueryUsersRequest()
    grpc.unary(pb.AccountService.QueryUsers, {
      request,
      host: config.grpc,

      // NOTE: Type vanished!
      onEnd: (res: any) => {
        console.log('QueryUsers', res)
        const data = res.message.getUsersList().map((u: any) => u.toObject())
        this.setState({ data, status: 'success' })
      }
    })
  }
}

export default WithUsers
