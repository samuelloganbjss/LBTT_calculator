import React, { useState } from 'react';
import './App.css'; 

function App() {
  const [price, setPrice] = useState('');
  const [isFirstTimeBuyer, setIsFirstTimeBuyer] = useState(false);
  const [isAdditionalDwelling, setIsAdditionalDwelling] = useState(false);
  const [lbtt, setLbtt] = useState(null);
  const [error, setError] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/calculate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          price: parseFloat(price),
          isFirstTimeBuyer: isFirstTimeBuyer,
          isAdditionalDwelling: isAdditionalDwelling,
        }),
      });

      if (!response.ok) {
        throw new Error('Error calculating LBTT');
      }

      const data = await response.json();
      setLbtt(data.lbtt);
      setError(null); 
    } catch (err) {
      setError('Error calculating LBTT');
      setLbtt(null);
    }
  };

  return (
    <div className="app-container">
      {/* Introductory Section */}
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
            To access the calculators - please click the 
            <a href="https://revenue.scot/calculate-property-transactions" target="_blank" rel="noopener noreferrer"> Property transactions calculator </a>
            or <a href="https://revenue.scot/lease-transactions-calculator" target="_blank" rel="noopener noreferrer">Lease transactions calculator</a> tab below.
          </p>
        </div>
      </section>

      {/* Section for Subheading */}
      <section className="property-transactions-section">
        <h2>Calculate property transactions</h2>
        <p>This calculator provides the LBTT liability on a property transaction based on the rates and bands in force at the effective date.</p>
      </section>

      {/* Property Transactions Calculator Section */}
      <section className="calculator-section">
        <div className="calculator-container">
          <h3>Property Transactions Calculator</h3>
          <form onSubmit={handleSubmit} className="calculator-form">
            <label>
              Purchase price (£):
              <input
                type="number"
                value={price}
                onChange={(e) => setPrice(e.target.value)}
                required
                className="purchase-price-input"
                placeholder="£0.00"
                min="0"
              />
            </label>
            <div className="additional-dwelling-section">
              <p>Is this purchase liable to the Additional Dwelling Supplement?</p>
              <div className="toggle-button-group">
                <button
                  type="button"
                  className={`toggle-button ${isAdditionalDwelling ? 'active' : ''}`}
                  onClick={() => setIsAdditionalDwelling(true)}
                >
                  Yes
                </button>
                <button
                  type="button"
                  className={`toggle-button ${!isAdditionalDwelling ? 'active' : ''}`}
                  onClick={() => setIsAdditionalDwelling(false)}
                >
                  No
                </button>
              </div>
            </div>
            <div className="first-time-buyer-section">
              <p>Are you claiming First-Time Buyer relief?</p>
              <div className="toggle-button-group">
                <button
                  type="button"
                  className={`toggle-button ${isFirstTimeBuyer ? 'active' : ''}`}
                  onClick={() => setIsFirstTimeBuyer(true)}
                >
                  Yes
                </button>
                <button
                  type="button"
                  className={`toggle-button ${!isFirstTimeBuyer ? 'active' : ''}`}
                  onClick={() => setIsFirstTimeBuyer(false)}
                >
                  No
                </button>
              </div>
            </div>
            <button type="submit" className="calculate-button">Calculate your tax</button>
          </form>
          {lbtt !== null && (
            <div role="alert" aria-live="polite">
              <p>LBTT: £{lbtt}</p>
            </div>
          )}
          {error && (
            <div role="alert" aria-live="assertive">
              <p className="error-message">{error}</p>
            </div>
          )}
          {/* Footer Section */}
          <footer className="app-footer">
            <p>For non-residential or mixed property transactions, please include VAT chargeable where appropriate.
            LBTT is a self-assessed tax and therefore Revenue Scotland does not accept liability for the use by taxpayers or agents of this calculator.</p>
          </footer>
        </div>
      </section>
    </div>
  );
}

export default App;

