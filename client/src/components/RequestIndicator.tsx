import * as React from 'react'

export type RequestState<Data> =
  RequestSuccessState<Data> |
  RequestLoadingState |
  RequestFailedState

interface RequestSuccessState<Data> {
  status: 'success'
  data: Data
  error: null
}

interface RequestLoadingState {
  status: 'loading'
  data: null
  error: null
}

interface RequestFailedState {
  status: 'failed'
  data: null
  error: Error
}

interface RenderProps<Data> {
  children: (data: Data) => React.ReactNode
}

export type Props<Data> = RenderProps<Data> & RequestState<Data>

export function RequestIndicator<E>(props: Props<E>) {

  if (props.status === 'failed') {
    return <div>Error: {props.error.toString()}</div>
  }

  if (props.status === 'loading') {
    return <div>Now Loading...</div>
  }

  if (props.status === 'success') {
    return props.children(props.data)
  }

  return null

}

export default RequestIndicator as React.SFC