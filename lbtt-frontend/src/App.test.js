import React from 'react';
import { render, fireEvent, screen } from '@testing-library/react';
import App from './App';

test('calculates LBTT from backend', async () => {
  // Mock the fetch call
  global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve({ lbtt: 5000 }),
    })
  );

  render(<App />);

  // Enter a property price
  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '100000' },
  });

  // Submit the form
  fireEvent.click(screen.getByText(/calculate/i));

  // Check that LBTT result is displayed
  const lbttElement = await screen.findByText(/LBTT: £5000/i);
  expect(lbttElement).toBeInTheDocument();
});
