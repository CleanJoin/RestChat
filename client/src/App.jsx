import { React, useState } from 'react';
import { CssBaseline } from '@mui/material';

import LoginForm from './Components/LoginForm';
import Chat from './Components/Chat';

import './App.css';
import '@fontsource/roboto/400.css';

function App({ apiClient }) {
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [memberName, setMemberName] = useState(null);

  const logoutHandler = () => {
    console.log("logoutHandler called");
    setIsAuthorized(false);
    setMemberName(null);
  };

  const sendMessageHandler = (text) => {
    console.log("sendMessageHandler called");
  };

  const getMembersHandler = () => {
    console.log("getMembersHandler called");
    return null;
  };

  const getMessagesHandler = () => {
    console.log("getMessagesHandler called");
    return null;
  };

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
            memberName={memberName}
            updateIntervalSec={5}
            logoutHandler={logoutHandler}
            sendMessageHandler={sendMessageHandler}
            getMembersHandler={getMembersHandler}
            getMessagesHandler={getMessagesHandler}
          />
        }

    </div>
  );
}

export default App;
