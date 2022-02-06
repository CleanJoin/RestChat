import { StatusCodes } from 'http-status-codes';

describe('integration smoke tests', () => {

    it('application frontend is accessible through proxy', () => {
        cy.visit('http://localhost:10000');
    });

    it('application backend is accessible through proxy', () => {
        cy.request('http://localhost:10000/api/');
    });

    it('backend successfully returns health status', () => {
        cy.request('http://localhost:10000/api/health').should((response) => {
            expect(response.status).to.be.equal(StatusCodes.OK);
            expect(response.body).to.have.property("success").equal(true);
            expect(response.body).to.have.property("time");
        });
    });
});