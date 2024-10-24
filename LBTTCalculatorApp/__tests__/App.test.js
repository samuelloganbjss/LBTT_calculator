import React from 'react';
import { render, fireEvent } from '@testing-library/react-native';
import App from '../App'; 

test('renders property price input', () => {
  const { getByPlaceholderText } = render(<App />);
  const priceInput = getByPlaceholderText('Enter property price');
  expect(priceInput).toBeTruthy();
});

test('renders first-time buyer and additional dwelling checkboxes', () => {
  const { getByText } = render(<App />);
  const firstTimeBuyerCheckbox = getByText('First-Time Buyer');
  const additionalDwellingCheckbox = getByText('Additional Dwelling');
  
  expect(firstTimeBuyerCheckbox).toBeTruthy();
  expect(additionalDwellingCheckbox).toBeTruthy();
});

test('renders calculate button', () => {
  const { getByText } = render(<App />);
  const calculateButton = getByText('Calculate LBTT');
  
  expect(calculateButton).toBeTruthy();
});

test('calculates LBTT and displays the result', () => {
  const { getByPlaceholderText, getByText, getByDisplayValue } = render(<App />);
  
  // Simulate entering a price
  const priceInput = getByPlaceholderText('Enter property price');
  fireEvent.changeText(priceInput, '300000');
  
  // Simulate clicking the calculate button
  const calculateButton = getByText('Calculate LBTT');
  fireEvent.press(calculateButton);
  
  // Ensure that the LBTT amount is displayed
  const lbttDisplay = getByText(/LBTT Amount: £/);
  expect(lbttDisplay).toBeTruthy();
});
