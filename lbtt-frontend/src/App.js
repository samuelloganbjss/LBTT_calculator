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
      setError(null); // Clear any previous error
    } catch (err) {
      setError('Error calculating LBTT');
      setLbtt(null);
    }
  };

  return (
    <div className="app-container">
      <header className="app-header">
        <h1>LBTT Calculator</h1>
      </header>
      <main className="app-main">
        <form onSubmit={handleSubmit} className="calculator-form">
          <label>
            Purchase price (£):
            <input
              type="number"
              value={price}
              onChange={(e) => setPrice(e.target.value)}
              required
              className="purchase-price-input" // Added class for styling
              placeholder="£0.00" // Updated placeholder for formatted look
              min="0" // Ensuring only positive values are allowed
            />
          </label>
          <label>
            <input
              type="checkbox"
              checked={isFirstTimeBuyer}
              onChange={(e) => setIsFirstTimeBuyer(e.target.checked)}
            />
            First-Time Buyer
          </label>
          <label>
            <input
              type="checkbox"
              checked={isAdditionalDwelling}
              onChange={(e) => setIsAdditionalDwelling(e.target.checked)}
            />
            Additional Dwelling
          </label>
          <button type="submit">Calculate your tax</button>
        </form>
        {lbtt !== null && <p>LBTT: £{lbtt}</p>}
        {error && <p className="error-message">{error}</p>}
      </main>
      <footer className="app-footer">
        <p>For non-residential or mixed property transactions, please include VAT chargeable where appropriate.
        LBTT is a self-assessed tax and therefore Revenue Scotland does not accept liability for the use by taxpayers or agents of this calculator.</p>
      </footer>
    </div>
  );
}

export default App;
