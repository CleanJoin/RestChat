import { useCallback, useEffect, useState } from "react";

import { AUTO_UPDATE_INTERVAL_SEC } from "../settings";

import MessageList from "./MessageList";
import MemberList from "./MemberList";


function Chat({
  apiClient,
  memberName,
  setIsAuthorized,
  setMemberName
}) {

  const [members, setMembers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [autoUpdateIntervalId, setAutoUpdateIntervalId] = useState(null);

  // TODO: Handle exceptions
  const updateMembers = useCallback(() => {
    console.log("Updating members");
    apiClient.getMembers().then(newMembers => {
      console.log("getMembers resolved, newMembers:", newMembers);
      if (newMembers !== null && newMembers.length > 0) {
        setMembers(newMembers);
      }
    });
  }, [apiClient]);

  // TODO: Handle exceptions
  const updateMessages = useCallback(() => {
    console.log("Updating messages");
    apiClient.getMessages().then(newMessages => {
      console.log("getMessages resolved, newMessages", newMessages);
      if (newMessages !== null && newMessages.length > 0) {
        setMessages(newMessages);
      }
    });
  }, [apiClient]);

  // TODO: Handle exceptions
  const onLogout = useCallback(() => {
    apiClient.logout().then(() => {
      if (autoUpdateIntervalId != null) {
        clearInterval(autoUpdateIntervalId);
      }
      setAutoUpdateIntervalId(null);
      setIsAuthorized(false);
      setMemberName(null);
    });
  }, [apiClient, setAutoUpdateIntervalId, setIsAuthorized, setMemberName]);

  const startAutoFetch = useCallback(() => {
    const fetchData = () => {
      updateMembers();
      updateMessages();
    };

    if (autoUpdateIntervalId === null) {
      fetchData();
      const intervalId = setInterval(fetchData, AUTO_UPDATE_INTERVAL_SEC * 1000);
      setAutoUpdateIntervalId(intervalId);
      console.log("Started auto data fetch loop");
    }
  }, [autoUpdateIntervalId, updateIntervalSec, updateMembers, updateMessages]);

  useEffect(() => {
    startAutoFetch();
  }, [startAutoFetch]);

  return (
    <div>
      <p>Member name: {memberName}</p>
      <button type="button" onClick={onLogout}>Logout</button>
      <button type="button">Send message</button>
      <MemberList members={members} />
      <MessageList messages={messages} />
    </div>
  );
}

export default Chat;