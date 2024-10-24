import React from 'react';
import { render, fireEvent, screen } from '@testing-library/react';
import App from './App';

beforeEach(() => {
  jest.clearAllMocks();
});

test('calculates LBTT from backend', async () => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve({ lbtt: 5000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '100000' },
  });

  fireEvent.click(screen.getByText(/calculate/i));

  const lbttElement = await screen.findByText(/LBTT: £5000/i);
  expect(lbttElement).toBeInTheDocument();
});

test('sends first-time buyer status to the backend', async () => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve({ lbtt: 2000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '150000' },
  });

  fireEvent.click(screen.getByLabelText(/first-time buyer/i));

  fireEvent.click(screen.getByText(/calculate/i));

  expect(global.fetch).toHaveBeenCalledWith(
    'http://localhost:8080/calculate',
    expect.objectContaining({
      body: JSON.stringify({
        price: 150000,
        isFirstTimeBuyer: true,
        isAdditionalDwelling: false,
      }),
    })
  );

  const lbttElement = await screen.findByText(/LBTT: £2000/i);
  expect(lbttElement).toBeInTheDocument();
});

test('sends additional dwelling status to the backend', async () => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve({ lbtt: 6000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '300000' },
  });

  fireEvent.click(screen.getByLabelText(/additional dwelling/i));

  fireEvent.click(screen.getByText(/calculate/i));

  expect(global.fetch).toHaveBeenCalledWith(
    'http://localhost:8080/calculate',
    expect.objectContaining({
      body: JSON.stringify({
        price: 300000,
        isFirstTimeBuyer: false,
        isAdditionalDwelling: true,
      }),
    })
  );

  const lbttElement = await screen.findByText(/LBTT: £6000/i);
  expect(lbttElement).toBeInTheDocument();
});

test('displays error message when backend returns an error', async () => {
  global.fetch = jest.fn(() =>
    Promise.reject(new Error('Server error'))
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '300000' },
  });

  fireEvent.click(screen.getByText(/calculate/i));

  const errorElement = await screen.findByText(/Error calculating LBTT/i);
  expect(errorElement).toBeInTheDocument();
});

test('displays validation error for invalid input', async () => {
  render(<App />);

  fireEvent.change(screen.getByLabelText(/property price/i), {
    target: { value: '-50000' },
  });

  fireEvent.click(screen.getByText(/calculate/i));

  const errorElement = await screen.findByText(/Error calculating LBTT/i);
  expect(errorElement).toBeInTheDocument();
});
