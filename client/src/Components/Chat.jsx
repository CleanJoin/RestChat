import { useCallback, useEffect, useState } from "react";

import { AUTO_UPDATE_INTERVAL_SEC } from "../settings";

import MessageList from "./MessageList";
import MemberList from "./MemberList";


function Chat({
  apiClient,
  isAuthorized,
  memberName,
  setIsAuthorized,
  setMemberName
}) {

  const [members, setMembers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [autoUpdateIntervalId, setAutoUpdateIntervalId] = useState(null);

  // TODO: Implement errors visualization
  const onError = useCallback((error) => {
    console.error(error);
  }, []);

  const updateMembers = useCallback(async () => {
    try {
      const newMembers = await apiClient.getMembers();
      console.log("getMembers resolved, newMembers:", newMembers);
      setMembers(newMembers);
    } catch (exception) {
      onError(`getMembers failed with error: ${exception.message}`);
    }
  }, [apiClient, onError]);

  const updateMessages = useCallback(async () => {
    try {
      const newMessages = await apiClient.getMessages();
      console.log("getMessages resolved, newMessages", newMessages);
      setMessages(newMessages);
    } catch (exception) {
      onError(`getMessages failed with error: ${exception.message}`);
    }
  }, [apiClient, onError]);

  const onLogout = useCallback(async () => {
    try {
      await apiClient.logout();
      console.log("Successfully logout from chat");
    } catch (exception) {
      onError(`Failed to logout from chat with error: ${exception.message}`);
    } finally {
      if (autoUpdateIntervalId !== null) {
        clearInterval(autoUpdateIntervalId);
        setAutoUpdateIntervalId(null);
      }
      setIsAuthorized(false);
      setMemberName(null);
    }
  }, [
    apiClient,
    autoUpdateIntervalId,
    setAutoUpdateIntervalId,
    setIsAuthorized,
    setMemberName,
    onError
  ]);

  const startAutoFetch = useCallback(() => {
    const fetchData = () => {
      updateMembers();
      updateMessages();
    };

    if (autoUpdateIntervalId === null && isAuthorized) {
      fetchData();
      const intervalId = setInterval(fetchData, AUTO_UPDATE_INTERVAL_SEC * 1000);
      setAutoUpdateIntervalId(intervalId);
      console.log("Started auto data fetch loop");
    }
  }, [autoUpdateIntervalId, updateMembers, updateMessages, isAuthorized]);

  useEffect(() => {
    if (isAuthorized) {
      startAutoFetch();
    }
  }, [startAutoFetch, isAuthorized]);

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