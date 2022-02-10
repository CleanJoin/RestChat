import ApiClientRest from './ApiClientRest';

describe('ApiClientRest authorization requirements', () => {

    test('client instance created unauthorized', () => {
        const client = new ApiClientRest();
        expect(client.isAuthorized()).toEqual(false);
    });

    test('requireAuthorization() throws error', () => {
        const client = new ApiClientRest();
        expect(() => client.requireAuthorization()).toThrow();
    });

    test('logout() requires authorization', async () => {
        const client = new ApiClientRest();
        await expect(async () => {
            await client.logout();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('getMembers() requires authorization', async () => {
        const client = new ApiClientRest();
        await expect(async () => {
            await client.getMembers();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('getMessages() requires authorization', async () => {
        const client = new ApiClientRest();
        await expect(async () => {
            await client.getMessages();
        }).rejects.toThrowError(/not authorized/i);
    });

    test('sendMessage() requires authorization', async () => {
        const client = new ApiClientRest();
        await expect(async () => {
            await client.getMessages();
        }).rejects.toThrowError(/not authorized/i);
    });

});