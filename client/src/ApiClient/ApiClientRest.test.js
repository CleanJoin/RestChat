import ApiClientRest from './ApiClientRest';

let client = null;

beforeEach(() => {
    client = new ApiClientRest();
});

describe('ApiClientRest authorization requirements', () => {

    test('client instance created unauthorized', () => {
        expect(client.isAuthorized()).toEqual(false);
    });

    test('requireAuthorization() throws error', () => {
        expect(() => client.requireAuthorization()).toThrow();
    });

    test('logout() requires authorization', async () => {
        await expect(async () => {
            await client.logout();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('getMembers() requires authorization', async () => {
        await expect(async () => {
            await client.getMembers();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('getMessages() requires authorization', async () => {
        await expect(async () => {
            await client.getMessages();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('sendMessage() requires authorization', async () => {
        await expect(async () => {
            await client.sendMessage("Some message");
        }).rejects.toThrowError(/not authorized/i);
    });

});

describe('ApiClientRest login', () => {

    test('login for non-existing user should fail', async () => {
        await expect(async () => {
            await client.login('User-666', 'Wrong-User-666-password')
        }).rejects.toThrow();
    });

    test('login with wrong password should fail', async () => {
        await expect(async () => {
            await client.login('User-2', 'Wrong-User-2-password')
        }).rejects.toThrow();
    });

    test('login with valid password should succeed', async () => {
        const username = 'User-1';
        const password = 'User-1_password';

        expect(client.isAuthorized()).toEqual(false);

        let memberName = await client.login(username, password);
        expect(memberName).toEqual(username);

        expect(client.isAuthorized()).toEqual(true);
    });

    test('api token should change between after second login', async () => {
        const username = 'User-1';
        const password = 'User-1_password';

        expect(client.isAuthorized()).toEqual(false);

        await client.login(username, password);
        const firstApiToken = client.apiToken;
        expect(client.isAuthorized()).toEqual(true);

        await client.login(username, password);
        const secondApiToken = client.apiToken;
        expect(client.isAuthorized()).toEqual(true);

        expect(firstApiToken).not.toEqual(secondApiToken);
    });

});

describe('ApiClientRest new user registration', () => {

    test('creating unique new user should succeed', async () => {
        const newUserName = 'User-test-register';
        const newPassword = 'p-word';

        const registeredUserName = await client.register(newUserName, newPassword);
        expect(registeredUserName).toEqual(newUserName);
    });

    test('login with new created user should succeed', async () => {
        const newUserName = 'User-test-register-login';
        const newPassword = 'password12345';

        expect(client.isAuthorized()).toEqual(false);
        await expect(async () => {
            await client.login(newUserName, newPassword);
        }).rejects.toThrow();

        const registeredUserName = await client.register(newUserName, newPassword);
        expect(registeredUserName).toEqual(newUserName);

        expect(client.isAuthorized()).toEqual(false);
        const newMemberName = await client.login(newUserName, newPassword);
        expect(client.isAuthorized()).toEqual(true);
        expect(newMemberName).toEqual(newUserName);
    });

    test('request to create already existing user should return error', async () => {
        await expect(async () => {
            await client.register('User-1', 'SomePassword');
        }).rejects.toThrow();
    });

});

describe('ApiClientRest logout', () => {

    test('logout not authorized client should fail', async () => {
        expect(client.isAuthorized()).toEqual(false);
        await expect(async () => {
            await client.logout();
        }).rejects.toThrow();
    });

    test('logout authorized should succeed', async () => {
        expect(client.isAuthorized()).toEqual(false);
        await client.login('User-4', 'User-4_password');
        expect(client.isAuthorized()).toEqual(true);
        await client.logout();
        expect(client.isAuthorized()).toEqual(false);
    });

});

describe('ApiClientRest get members list', () => {

    test('authorized user should get online members list', async () => {
        await client.login('User-4', 'User-4_password');
        const members = await client.getMembers();
        expect(members.length).toEqual(3);
    });

    test('new user should find himself in users list', async () => {
        await client.login('User-1', 'User-1_password');
        const members = await client.getMembers();
        expect(members.length).toEqual(4);
    });

    test('user logout should decrease online users count', async () => {

        await client.login('User-4', 'User-4_password');
        expect(client.isAuthorized()).toEqual(true);

        const client2 = new ApiClientRest();
        await client2.login('User-5', 'User-5_password');
        expect(client2.isAuthorized()).toEqual(true);

        const allMembers1 = await client.getMembers();
        const allMembers2 = await client2.getMembers();

        expect(allMembers1.length).toEqual(allMembers2.length);

        const membersNumBeforeLogout = allMembers2.length;
        await client.logout();
        const membersNumAfterLogout = (await client2.getMembers()).length;

        expect(membersNumBeforeLogout - membersNumAfterLogout).toEqual(1);
    });

});

describe('ApiClientRest can receive and send messages', () => {

    test('authorized user should get previous chat messages', async () => {
        await client.login('User-1', 'User-1_password');
        const messages = await client.getMessages();
        expect(messages.length).toEqual(10);
    });

    test('different clients see same messages', async () => {
        await client.login('User-1', 'User-1_password');
        const client2 = new ApiClientRest();
        await client2.login('User-2', 'User-2_password');

        const messages = await client.getMessages();
        expect(messages.length).toEqual(10);

        const messages2 = await client.getMessages();
        expect(messages2.length).toEqual(10);

    });

    test('client can send message and receive it back', async () => {
        const text = "This is very new message sent by User-1";

        await client.login('User-1', 'User-1_password');
        let messagesBefore = await client.getMessages();

        // Message should not be present in message list yet
        expect(messagesBefore.filter(msg => msg.text === text).length).toEqual(0);

        await client.sendMessage(text);
        let messagesAfter = await client.getMessages();
        expect(messagesAfter.filter(msg => msg.text === text).length).toEqual(1);

        const newMessage = messagesAfter.filter(msg => msg.text === text)[0];
        expect(newMessage.member_name).toEqual('User-1');

    });

    test('other client should see new message', async () => {
        const text = "This is example of message sent by User-3 for User-6";

        await client.login('User-3', 'User-3_password');
        let messagesBefore = await client.getMessages();
        expect(messagesBefore.filter(msg => msg.text === text).length).toEqual(0);

        const client2 = new ApiClientRest();
        await client2.login('User-6', 'User-6_password');

        // Message should not be present in message list yet
        expect(messagesBefore.filter(msg => msg.text === text).length).toEqual(0);

        await client.sendMessage(text);

        let messagesAfter = await client2.getMessages();
        expect(messagesAfter.filter(msg => msg.text === text).length).toEqual(1);
    });

});
