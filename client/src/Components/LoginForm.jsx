import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Paper from '@mui/material/Paper';

import Fingerprint from '@mui/icons-material/Fingerprint';
import AppRegistration from '@mui/icons-material/AppRegistration';

import logo from '../Images/RestChat_logo.png';

import { Container, TextField } from '@mui/material';
import { Box } from '@mui/system';


function LoginForm({ loginHandler, registerHandler }) {

  const handleSubmit = () => {
    console.log("Form has been submited");
  };

  const handleLogin = () => {
    console.log("Login button clicked");
  };

  const handleRegister = () => {
    console.log("Register button clicked");
  };

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
            autoFocus
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
              onClick={handleLogin}
              startIcon={<Fingerprint />}
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
              onClick={handleRegister}
            >
              Register
            </Button>
          </Stack>

          { /*
          <Alert severity="error" variant="filled">Error</Alert>
          <Alert severity="error" variant="filled">Error</Alert>
          <Alert severity="error" variant="filled">Error</Alert>
          <Alert severity="error" variant="filled">Error</Alert>
          */ }

        </Stack>

      </Box>

    </Container>
  );
}

export default LoginForm;