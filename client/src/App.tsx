import * as React from 'react'
import './App.css'

import { UserList, WithUsers } from './components'

function App() {
  return (
    <div className="App">
      <WithUsers>
        {({ users }) => (
          <UserList users={users} />
        )}
      </WithUsers>
    </div>
  )
}

export default App
