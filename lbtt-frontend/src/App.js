import React, { useState } from 'react';

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
          isAdditionalDwelling: isAdditionalDwelling
        }),
      });
      const data = await response.json();
      setLbtt(data.lbtt);
    } catch (err) {
      setError('Error calculating LBTT');
    }
  };

  return (
    <div>
      <h1>LBTT Calculator</h1>
      <form onSubmit={handleSubmit}>
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
          First-Time Buyer:
          <input
            type="checkbox"
            checked={isFirstTimeBuyer}
            onChange={(e) => setIsFirstTimeBuyer(e.target.checked)}
          />
        </label>
        <label>
          Additional Dwelling:
          <input
            type="checkbox"
            checked={isAdditionalDwelling}
            onChange={(e) => setIsAdditionalDwelling(e.target.checked)}
          />
        </label>
        <button type="submit">Calculate</button>
      </form>
      {lbtt !== null && <p>LBTT: £{lbtt}</p>}
      {error && <p>{error}</p>}
    </div>
  );
}

export default App;
