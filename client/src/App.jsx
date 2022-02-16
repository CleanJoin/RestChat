import { React, useState } from 'react';
import { CssBaseline } from '@mui/material';

import LoginForm from './Components/LoginForm';
import Chat from './Components/Chat';

import './App.css';
import '@fontsource/roboto/400.css';

function App({ apiClient }) {
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [memberName, setMemberName] = useState(null);

  const loginHandler = (username, password) => {
    console.log("loginHandler called");
    const memberName = apiClient.login(username, password);
    console.log("Login successful with username:", username);
    setMemberName(memberName);
    setIsAuthorized(true);
  };

  const registerHandler = (username, password) => {
    console.log("registerHandler called");
    const memberName = apiClient.register(username, password);
    console.log("Registration successful with memberName:", memberName);
  };

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
            loginHandler={loginHandler}
            registerHandler={registerHandler}
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
