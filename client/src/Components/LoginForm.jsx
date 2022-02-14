function LoginForm(props) {
  return (
    <div>
      <p>Login Form</p>
      <button type="button" onClick={() => props.loginHandler()}>Login</button>
      <button type="button" onClick={() => props.registerHandler()}>Register</button>
    </div>
  );
}

export default LoginForm;