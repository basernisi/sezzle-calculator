import { useState } from 'react';
import { calculate } from '../api/calculatorApi';
import type { CalculatorSubmission } from '../types/calculator';
import { HttpError } from '../api/client';

interface CalculatorState {
  result: number | null;
  error: string | null;
  isLoading: boolean;
}

const initialState: CalculatorState = {
  result: null,
  error: null,
  isLoading: false,
};

export function useCalculator() {
  const [state, setState] = useState<CalculatorState>(initialState);

  async function submit(submission: CalculatorSubmission) {
    setState((current) => ({
      ...current,
      isLoading: true,
      error: null,
    }));

    try {
      const response = await calculate(
        {
          operation: submission.operation,
          operands: submission.operands,
        },
        submission.token,
      );
      setState({
        result: response.result,
        error: null,
        isLoading: false,
      });
    } catch (error) {
      const message =
        error instanceof HttpError ? error.message : 'Something went wrong while calling the API.';

      setState({
        result: null,
        error: message,
        isLoading: false,
      });
    }
  }

  function reset() {
    setState(initialState);
  }

  return {
    ...state,
    submit,
    reset,
  };
}
