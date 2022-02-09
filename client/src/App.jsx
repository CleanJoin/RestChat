import { React, useState } from 'react';
import LoginForm from './Components/LoginForm';
import Chat from './Components/Chat';


import logo from './logo.svg';
import './App.css';

function App(props) {
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [memberName, setMemberName] = useState(null);

  const loginHandler = (username, password) => {
    console.log("loginHandler called");
    const [memberName, error] = props.apiClient.login(username, password);
    console.log("username, error:", username, error);

    if (error == null) {
      setMemberName(memberName);
      setIsAuthorized(true);
    }
  };

  const registerHandler = (username, password) => {
    console.log("registerHandler called");
    const [memberName, error] = props.apiClient.register(username, password);
    console.log("username, error:", username, error);

    if (error == null) {
      console.log("Registration successful");
    }
  };

  const logoutHandler = () => {
    console.log("logoutHandler called");
  };

  const sendMessageHandler = (text) => {
    console.log("sendMessageHandler called");
  };

  const getMembersHandler = () => {
    console.log("getMembersHandler called");
  };

  const getMessagesHandler = () => {
    console.log("getMessagesHandler called");
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />

        <p>isAuthorized: {isAuthorized.toString()}</p>

        {!isAuthorized &&
          <LoginForm
            loginHandler={loginHandler}
            registerHandler={registerHandler}
          />
        }

        {isAuthorized &&
          <Chat
            memberName={memberName}
            logoutHandler={logoutHandler}
            sendMessageHandler={sendMessageHandler}
            getMembersHandler={getMembersHandler}
            getMessagesHandler={getMessagesHandler}
          />
        }

      </header>

    </div>
  );
}

export default App;
