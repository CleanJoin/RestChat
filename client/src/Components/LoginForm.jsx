import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Alert from '@mui/material/Alert';
import Fingerprint from '@mui/icons-material/Fingerprint';
import AppRegistration from '@mui/icons-material/AppRegistration';

import logo from '../Images/RestChat_logo.png';

import { Container, TextField } from '@mui/material';
import { Box } from '@mui/system';
import { useState } from 'react';


const MAX_LOGIN_LENGTH = 15;
const LOGIN_REGEX = /^([a-zA-Z0-9_-]){1,}$/;
const MAX_PASSWORD_LENGTH = 32;
const PASSWORD_REGEX = /^([a-zA-Z0-9_-]){1,}$/;

const MESSAGE_ERROR = "error";
const MESSAGE_SUCCESS = "success";

function LoginForm({ apiClient, setIsAuthorized, setMemberName }) {

  const [messages, setMessages] = useState([]);

  const handleSubmit = (event) => {
    event.preventDefault();
  };

  const addMessage = (text, type) => {
    setMessages([...messages, { text: text, type: type }]);
  }

  const deleteMessage = (index) => {
    setMessages([...messages.filter((x, i) => i !== index)]);
  }

  const getFormParams = () => {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    return [username, password]
  }

  const sanitizeFormParams = () => {
    const [username, password] = getFormParams();

    if (username.length < 1 || username.length > MAX_LOGIN_LENGTH) {
      throw new Error(`Username length should be 1 to ${MAX_LOGIN_LENGTH} symbols.`);
    }

    if (password.length < 1 || password.length > MAX_PASSWORD_LENGTH) {
      throw new Error(`Password length should be 1 to ${MAX_PASSWORD_LENGTH} symbols.`);
    }

    if(!LOGIN_REGEX.test(username)) {
      throw new Error('Allowed username symbols is [a-zA-Z0-9_].');
    }

    if(!PASSWORD_REGEX.test(password)) {
      throw new Error('Allowed password symbols is [a-zA-Z0-9_].');
    }

    return [username, password];
  }

  const handleLogin = async () => {
    try {
      const [username, password] = sanitizeFormParams();
      const memberName = await apiClient.login(username, password);
      setIsAuthorized(true);
      setMemberName(memberName);
    } catch (exception) {
      addMessage(exception.message, MESSAGE_ERROR);
    }
  }

  const handleRegister = async () => {
    const [username, password] = getFormParams();
    try {
      const memberName = await apiClient.register(username, password);
      addMessage(`Successfully registered user ${memberName}. Login to proceed.`, MESSAGE_SUCCESS);
    } catch (exception) {
      addMessage(exception.message, MESSAGE_ERROR);
    }
  }

  return (
    <Container
      component="main"
      maxWidth="sm"
      sx={{
        minHeight: '100vh',
      }}
    >

      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          minHeight: '100vh',
          display: 'flex',
          flexDirection: 'column',
          textAlign: 'center',
          paddingTop: '20vh'
        }}
      >

        <Stack direction="column" spacing={2}>

          <img src={logo} alt="RestChat logo"></img>

          <TextField
            fullWidth
            required
            id="username"
            name="username"
            label="Username"
          />

          <TextField
            fullWidth
            required
            id="password"
            name="password"
            label="Password"
            type="password"
          />

          <Stack direction="row" spacing={1}>
            <Button
              type="submit"
              variant="contained"
              fullWidth
              size="large"
              startIcon={<Fingerprint />}
              onClick={async () => { handleLogin() }}
            >
              Login
            </Button>

            <Button
              type="submit"
              variant="contained"
              color="warning"
              fullWidth
              size="large"
              startIcon={<AppRegistration />}
              onClick={async () => { handleRegister() }}
            >
              Register
            </Button>
          </Stack>


          <Stack direction="column-reverse" spacing={1}>

            {
              messages.map((msg, i) => {
                return (
                  <Alert
                    severity={msg.type}
                    key={i}
                    onClose={() => deleteMessage(i)}
                    >
                    {msg.text}
                  </Alert>
                )
              })
            }

          </Stack>

        </Stack>

      </Box>

    </Container>
  );
}

export default LoginForm;