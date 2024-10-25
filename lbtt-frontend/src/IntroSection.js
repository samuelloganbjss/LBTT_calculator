import React from 'react';

function IntroSection() {
  return (
    <section className="intro-section">
      <header className="intro-header">
        <h1>Calculate tax</h1>
      </header>
      <div className="intro-text">
        <p>
          The tax calculator allows taxpayers and agents to work out the amount of LBTT payable
          on residential, non-residential or mixed property transactions, and non-residential lease
          transactions based on the rates and thresholds.
        </p>
        <p>
          LBTT is a self-assessed tax and therefore Revenue Scotland does not accept liability for
          the use by taxpayers or agents of this calculator.
        </p>
        <p>
          For more information on rates and bands or how the tax works see 
          <a href="https://revenue.scot/land-buildings-transaction-tax" target="_blank" rel="noopener noreferrer"> Land and Buildings Transaction Tax (LBTT) guidance</a>.
        </p>
        <p>
          To access the calculators - please use the below
          <a href="https://revenue.scot/calculate-tax/calculate-property-transactions" target="_blank" rel="noopener noreferrer"> Property transactions calculator  </a> below.
        </p>
      </div>
    </section>
  );
}

export default IntroSection;
