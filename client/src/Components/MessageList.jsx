function MessageList({ messages }) {
    return (
        <div>
            <p>Chat messages:</p>
            <ul>
                {
                    messages.map((message, index) => {
                        return <li key={index}>
                            #{index} ({message.id}) [{message.time}]: [{message.member_name}]:{message.text}
                        </li>
                    })
                }
            </ul>
        </div>
    );
}

export default MessageList;