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
            Property Price:
            <input
              type="number"
              value={price}
              onChange={(e) => setPrice(e.target.value)}
              required
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
          <button type="submit">Calculate</button>
        </form>
        {lbtt !== null && <p>LBTT: £{lbtt}</p>}
        {error && <p className="error-message">{error}</p>}
      </main>
      <footer className="app-footer">
        <p>© 2024 Your Company Name</p>
      </footer>
    </div>
  );
}

export default App;
