function MemberList({ members }) {
  return (
      <div>
          <p>Members:</p>
          <ul>
              {
                  members.map((member, index) => {
                      return <li key={index}>{member}</li>
                  })
              }
          </ul>
      </div>
  );
}

export default MemberList;