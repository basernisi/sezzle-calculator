import { fireEvent, render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import { CalculatorForm } from './CalculatorForm';

vi.mock('../api/authApi', () => ({
  issueToken: vi.fn(),
}));

describe('CalculatorForm', () => {
  it('shows validation feedback when operands are missing', async () => {
    const onSubmit = vi.fn().mockResolvedValue(undefined);

    render(
      <CalculatorForm
        isLoading={false}
        error={null}
        onSubmit={onSubmit}
        onResetRemoteState={vi.fn()}
      />,
    );

    fireEvent.change(screen.getByLabelText('Bearer token'), { target: { value: 'dev-token' } });
    fireEvent.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(await screen.findByRole('alert')).toHaveTextContent('Please enter 2 valid operands.');
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('requires a bearer token before submitting', async () => {
    const onSubmit = vi.fn().mockResolvedValue(undefined);

    render(
      <CalculatorForm
        isLoading={false}
        error={null}
        onSubmit={onSubmit}
        onResetRemoteState={vi.fn()}
      />,
    );

    fireEvent.change(screen.getByLabelText('Operand 1'), { target: { value: '10' } });
    fireEvent.change(screen.getByLabelText('Operand 2'), { target: { value: '5' } });
    fireEvent.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(await screen.findByRole('alert')).toHaveTextContent(
      'Please provide a bearer token before calling the API.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('requires client credentials before generating a token', async () => {
    render(
      <CalculatorForm
        isLoading={false}
        error={null}
        onSubmit={vi.fn().mockResolvedValue(undefined)}
        onResetRemoteState={vi.fn()}
      />,
    );

    fireEvent.click(screen.getByRole('button', { name: 'Generate token' }));

    expect(await screen.findByRole('alert')).toHaveTextContent(
      'Please provide the client ID and client secret to generate a token.',
    );
  });
});
