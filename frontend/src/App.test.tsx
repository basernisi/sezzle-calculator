import { render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import App from './App';

vi.mock('./hooks/useCalculator', () => ({
  useCalculator: () => ({
    result: 15,
    error: null,
    isLoading: false,
    submit: vi.fn(),
    reset: vi.fn(),
  }),
}));

describe('App', () => {
  it('renders calculator heading and result', () => {
    render(<App />);

    expect(screen.getByText(/secure calculator with a clean service boundary/i)).toBeInTheDocument();
    expect(screen.getByText('15')).toBeInTheDocument();
  });
});
