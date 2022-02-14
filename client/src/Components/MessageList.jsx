function MessageList({ messages }) {
  return (
      <div>
          <p>Chat messages:</p>
          <ul>
              {
                  messages.map((message, index) => {
                      return <li key={index}>{message}</li>
                  })
              }
          </ul>
      </div>
  );
}

export default MessageList;