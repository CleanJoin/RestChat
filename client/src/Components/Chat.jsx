function Chat(props) {
  return (
    <div>
      <p>Chat</p>
      <p>Member name: {props.memberName}</p>
      <button type="button" onClick={() => props.logoutHandler()}>Logout</button>
      <button type="button" onClick={() => props.sendMessageHandler()}>Send message</button>
    </div>
  );
}

export default Chat;