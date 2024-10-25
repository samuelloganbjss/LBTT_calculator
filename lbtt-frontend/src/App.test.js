import React from 'react';
import { render, fireEvent, screen, waitFor } from '@testing-library/react';
import App from './App';

beforeEach(() => {
  jest.clearAllMocks();
  jest.restoreAllMocks(); 
});

test('retrieves LBTT from backend', async () => {
  const fetchMock = jest.spyOn(global, 'fetch').mockImplementation(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ lbtt: 5000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/purchase price/i), {
    target: { value: '100000' },
  });

  fireEvent.click(screen.getByText(/calculate your tax/i));

  await waitFor(() => {
    const lbttElement = screen.getByRole('alert');
    expect(lbttElement).toHaveTextContent('LBTT Amount: £5000');
  });


  expect(fetchMock).toHaveBeenCalledWith(
    'http://localhost:8080/calculate',
    expect.objectContaining({
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        price: 100000,
        isFirstTimeBuyer: false,
        isAdditionalDwelling: false,
      }),
    })
  );
});

test('sends first-time buyer status to the backend', async () => {
  const fetchMock = jest.spyOn(global, 'fetch').mockImplementation(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ lbtt: 2000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/purchase price/i), {
    target: { value: '150000' },
  });

  fireEvent.click(screen.getByText(/yes/i, { selector: '.first-time-buyer-section button' }));

  fireEvent.click(screen.getByText(/calculate your tax/i));

  await waitFor(() => {
    const lbttElement = screen.getByRole('alert');
    expect(lbttElement).toHaveTextContent('LBTT Amount: £2000');
  });

  expect(fetchMock).toHaveBeenCalledWith(
    'http://localhost:8080/calculate',
    expect.objectContaining({
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        price: 150000,
        isFirstTimeBuyer: true,
        isAdditionalDwelling: false,
      }),
    })
  );
});

test('sends additional dwelling status to the backend', async () => {
  const fetchMock = jest.spyOn(global, 'fetch').mockImplementation(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ lbtt: 6000 }),
    })
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/purchase price/i), {
    target: { value: '300000' },
  });

  fireEvent.click(screen.getByText(/yes/i, { selector: '.additional-dwelling-section button' }));

  fireEvent.click(screen.getByText(/calculate your tax/i));

  await waitFor(() => {
    const lbttElement = screen.getByRole('alert');
    expect(lbttElement).toHaveTextContent('LBTT Amount: £6000');
  });

  expect(fetchMock).toHaveBeenCalledWith(
    'http://localhost:8080/calculate',
    expect.objectContaining({
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        price: 300000,
        isFirstTimeBuyer: false,
        isAdditionalDwelling: true,
      }),
    })
  );
});

test('displays error message when backend returns an error', async () => {
  jest.spyOn(global, 'fetch').mockImplementation(() =>
    Promise.reject(new Error('Server error'))
  );

  render(<App />);

  fireEvent.change(screen.getByLabelText(/purchase price/i), {
    target: { value: '300000' },
  });

  fireEvent.click(screen.getByText(/calculate your tax/i));

  await waitFor(() => {
    const errorElement = screen.getByRole('alert');
    expect(errorElement).toHaveTextContent('Error calculating LBTT');
  });
});

test('displays validation error for invalid input', async () => {
  render(<App />);

  fireEvent.change(screen.getByLabelText(/purchase price/i), {
    target: { value: '-50000' },
  });

  fireEvent.click(screen.getByText(/calculate your tax/i));

  await waitFor(() => {
    const errorElement = screen.getByRole('alert');
    expect(errorElement).toHaveTextContent('Error calculating LBTT');
  });
});
