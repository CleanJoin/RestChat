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

describe('ApiClientRest new user registration', () => {

    test('creating unique new user should succeed', async () => {
        const newUserName = 'AbsolutelyUnknownNewUser';
        const newPassword = 'p-word';

        await expect(async () => {
            await client.login(newUserName, newPassword);
        }).rejects.toThrow();

        const registeredUserName = await client.register(newUserName, newPassword);
        expect(registeredUserName).toEqual(newUserName);

        const memberName = await client.login(newUserName, newPassword);
        expect(memberName).toEqual(newUserName);
    });

});