import React from 'react';

function Calculator({ price, setPrice, isFirstTimeBuyer, setIsFirstTimeBuyer, isAdditionalDwelling, setIsAdditionalDwelling, handleSubmit, error }) {
  return (
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
        {error && (
          <div role="alert" aria-live="assertive">
            <p className="error-message">{error}</p>
          </div>
        )}
        <footer className="app-footer">
          <p>For non-residential or mixed property transactions, please include VAT chargeable where appropriate.
          LBTT is a self-assessed tax and therefore Revenue Scotland does not accept liability for the use by taxpayers or agents of this calculator.</p>
        </footer>
      </div>
    </section>
  );
}

export default Calculator;
