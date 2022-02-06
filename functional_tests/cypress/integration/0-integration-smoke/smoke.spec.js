describe('integration smoke tests', () => {

    it('application proxy is accessible', () => {
        cy.visit('http://localhost:10000');
        cy.request('http://localhost:10000/api/health');
    });

});