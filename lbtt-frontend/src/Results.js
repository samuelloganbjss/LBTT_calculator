import React from 'react';

function Results({ showResults, price, isAdditionalDwelling, isFirstTimeBuyer, lbtt }) {
  if (!showResults) return null;

  return (
    <section className="results-container">
      <h3>Values Entered:</h3>
      <div className="values-entered">
        <p className="results-item"><strong>Purchase price:</strong> £{price}</p>
        <p className="results-item"><strong>Is Additional Dwelling Supplement liable?</strong> {isAdditionalDwelling ? "Yes" : "No"}</p>
        <p className="results-item"><strong>Claiming First-Time Buyer Relief:</strong> {isFirstTimeBuyer ? "Yes" : "No"}</p>
      </div>

      <h3>Results:</h3>
      <div className="results">
        {lbtt !== null && (
          <div role="alert" aria-live="polite">
            <p className="results-item"><strong>LBTT Amount:</strong> £{lbtt}</p>
          </div>
        )}
        <button className="reset-button" onClick={() => window.location.reload()}>
          Reset calculation
        </button>
      </div>
    </section>
  );
}

export default Results;
