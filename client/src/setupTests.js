// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import mockDbFabric from './MockApi/db';
import mockHandlersFabric from './MockApi/handlers';
import { MockApiServer, ApiBuilderServer } from './MockApi/server'

const mockServer = new MockApiServer(
    new ApiBuilderServer(),
    mockDbFabric,
    mockHandlersFabric
);

beforeEach(() => {
    mockServer.reset();
    mockServer.start();
});

afterEach(() => {
    mockServer.stop();
});