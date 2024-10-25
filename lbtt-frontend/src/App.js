import React, { useState } from 'react';
import './App.css'; 
import IntroSection from './IntroSection';
import PropertyTransactionsSection from './PropertyTransactionsSection';
import Calculator from './Calculator';
import Results from './Results';

function App() {
  const [price, setPrice] = useState('');
  const [isFirstTimeBuyer, setIsFirstTimeBuyer] = useState(false);
  const [isAdditionalDwelling, setIsAdditionalDwelling] = useState(false);
  const [lbtt, setLbtt] = useState(null);
  const [error, setError] = useState(null);
  const [showResults, setShowResults] = useState(false); 

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
      setShowResults(true); 
    } catch (err) {
      setError('Error calculating LBTT');
      setLbtt(null);
      setShowResults(false); 
    }
  };

  return (
    <div className="app-container">
      <IntroSection />
      <PropertyTransactionsSection />
      <div className="calculator-and-results-container">
        <Calculator 
          price={price}
          setPrice={setPrice}
          isFirstTimeBuyer={isFirstTimeBuyer}
          setIsFirstTimeBuyer={setIsFirstTimeBuyer}
          isAdditionalDwelling={isAdditionalDwelling}
          setIsAdditionalDwelling={setIsAdditionalDwelling}
          handleSubmit={handleSubmit}
          error={error}
        />
        <Results 
          showResults={showResults}
          price={price}
          isAdditionalDwelling={isAdditionalDwelling}
          isFirstTimeBuyer={isFirstTimeBuyer}
          lbtt={lbtt}
        />
      </div>
    </div>
  );
}

export default App;
