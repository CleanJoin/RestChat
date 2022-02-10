import ApiClientRest from './ApiClientRest';

describe('ApiClientRest authorization requirements', () => {

    test('client instance created unauthorized', () => {
        const apiClient = new ApiClientRest()
        expect(apiClient.isAuthorized()).toEqual(false);
    });

})