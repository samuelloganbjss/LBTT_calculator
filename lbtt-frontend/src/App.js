import React, { useState } from 'react';

function App() {
  const [price, setPrice] = useState('');
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
        body: JSON.stringify({ price: parseFloat(price) }),
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
        <button type="submit">Calculate</button>
      </form>
      {lbtt !== null && <p>LBTT: £{lbtt}</p>}
      {error && <p>{error}</p>}
    </div>
  );
}

export default App;
