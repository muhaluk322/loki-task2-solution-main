describe('Loki Test', () => {
  it('Loki shows logs of "webapp"', () => {
    cy.visit(`${Cypress.env('GRAFANA_URL')}/explore`);
    cy.contains('Label browser').click();
    cy.get('button[title="varlogs"]').click();
    cy.get('button[aria-label="Use selector as logs button"]').click();
    cy.get('[class$="positionRelative"]').should('contain', 'GET');
  });
});
