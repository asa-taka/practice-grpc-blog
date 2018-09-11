import * as React from 'react'
import './App.css'

import { UserCard, UserList, WithUser, WithUsers } from './components'

function App() {
  return (
    <div className="App">
      <h2>QueryUsers</h2>
      <WithUsers>
        {({ users }) => (
          <UserList users={users} />
        )}
      </WithUsers>

      <h2>GetUser</h2>
      <WithUser>
        {({ user }) => (
          <>
            {user && (
              <UserCard user={user} />
            )}
          </>
        )}
      </WithUser>
    </div>
  )
}

export default App
