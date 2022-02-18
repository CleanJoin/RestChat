import { React, useState } from 'react';
import { CssBaseline } from '@mui/material';

import { AUTO_UPDATE_INTERVAL_SEC } from './restrictions';
import LoginForm from './Components/LoginForm';
import Chat from './Components/Chat';

import './App.css';
import '@fontsource/roboto/400.css';


function App({ apiClient }) {
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [memberName, setMemberName] = useState(null);

  return (
    <div className="App">
    <CssBaseline />

        {!isAuthorized &&
          <LoginForm
            apiClient={apiClient}
            setIsAuthorized={setIsAuthorized}
            setMemberName={setMemberName}
          />
        }

        {isAuthorized &&
          <Chat
            apiClient={apiClient}
            memberName={memberName}
            updateIntervalSec={AUTO_UPDATE_INTERVAL_SEC}
            setIsAuthorized={setIsAuthorized}
            setMemberName={setMemberName}
          />
        }

    </div>
  );
}

export default App;
