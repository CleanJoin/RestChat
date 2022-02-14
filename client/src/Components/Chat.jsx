import { useCallback, useEffect, useState } from "react";
import MessageList from "./MessageList";
import MemberList from "./MemberList";

function Chat({
  memberName,
  updateIntervalSec,
  logoutHandler,
  sendMessageHandler,
  getMembersHandler,
  getMessagesHandler
}) {

  const [members, setMembers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [autoUpdateIntervalId, setAutoUpdateIntervalId] = useState(null);

  const updateMembers = useCallback(() => {
    console.log("Updating members");
    const newMembers = getMembersHandler();
    // setMembers(newMembers);
  }, [getMembersHandler]);

  const updateMessages = useCallback(() => {
    console.log("Updating messages");
    const newMessages = getMessagesHandler();
    // setMessages(newMessages);
  }, [getMessagesHandler]);

  const onLogout = () => {
    if (autoUpdateIntervalId != null) {
      clearInterval(autoUpdateIntervalId);
    }
    setAutoUpdateIntervalId(null);
    logoutHandler();
  }

  const startAutoFetch = useCallback(() => {
    const fetchData = () => {
      updateMembers();
      updateMessages();
    };

    if (autoUpdateIntervalId === null) {
      fetchData();
      const intervalId = setInterval(fetchData, updateIntervalSec * 1000);
      setAutoUpdateIntervalId(intervalId);
      console.log("Started auto data fetch loop");
    }
  }, [autoUpdateIntervalId, updateIntervalSec, updateMembers, updateMessages]);

  useEffect(() => {
    startAutoFetch();
  }, [startAutoFetch]);


  return (
    <div>
      <p>Chat</p>
      <p>Member name: {memberName}</p>
      <button type="button" onClick={onLogout}>Logout</button>
      <button type="button" onClick={sendMessageHandler}>Send message</button>
      <MemberList members={members} />
      <MessageList messages={messages} />
    </div>
  );
}

export default Chat;