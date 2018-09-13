import * as React from 'react'
import './App.css'

import { RequestIndicator, UserCard, UserList, WithUser, WithUsers } from './components'

function App() {
  return (
    <div className="App">
      <h2>QueryUsers</h2>
      <WithUsers>
        {data => (
          <UserList users={data} />
        )}
      </WithUsers>

      <h2>GetUser</h2>
      <WithUser>
        {data => (
          <UserCard user={data} />
        )}
      </WithUser>

    </div>
  )
}

export default App
