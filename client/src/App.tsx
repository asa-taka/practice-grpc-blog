import * as React from 'react'
import './App.css'

import { UserCard, UserList, WithUser, WithUsers } from './components'

function App() {
  return (
    <div className="App">
      <WithUsers>
        {({ users }) => (
          <UserList users={users} />
        )}
      </WithUsers>
      <WithUser>
        {({ user }) => (
          <UserCard user={user} />
        )}
      </WithUser>
    </div>
  )
}

export default App
