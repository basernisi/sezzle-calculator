import { useEffect, useState, type FormEvent } from 'react';
import { issueToken } from '../api/authApi';
import { HttpError } from '../api/client';
import { ErrorAlert } from './ErrorAlert';
import { OperandInputs } from './OperandInputs';
import { OperationSelector } from './OperationSelector';
import type { CalculatorSubmission, Operation } from '../types/calculator';
import {
  getOperationDefinition,
  sanitizeOperandInputs,
  toOperandNumbers,
  validateOperandInputs,
} from '../utils/validation';

interface CalculatorFormProps {
  isLoading: boolean;
  error: string | null;
  onSubmit: (input: CalculatorSubmission) => Promise<void>;
  onResetRemoteState: () => void;
}

const tokenStorageKey = 'sezzle-calculator-token';
const clientIDStorageKey = 'sezzle-calculator-client-id';

export function CalculatorForm({
  isLoading,
  error,
  onSubmit,
  onResetRemoteState,
}: CalculatorFormProps) {
  const [operation, setOperation] = useState<Operation>('add');
  const [operandInputs, setOperandInputs] = useState<string[]>(['', '']);
  const [clientID, setClientID] = useState<string>(() => {
    if (typeof window === 'undefined') {
      return 'sezzle-demo-client';
    }

    return window.sessionStorage.getItem(clientIDStorageKey) ?? 'sezzle-demo-client';
  });
  const [clientSecret, setClientSecret] = useState<string>('');
  const [token, setToken] = useState<string>(() => {
    if (typeof window === 'undefined') {
      return '';
    }

    return window.sessionStorage.getItem(tokenStorageKey) ?? '';
  });
  const [validationError, setValidationError] = useState<string | null>(null);
  const [isGeneratingToken, setIsGeneratingToken] = useState(false);

  useEffect(() => {
    setOperandInputs((current) => {
      const definition = getOperationDefinition(operation);
      const next = sanitizeOperandInputs(current, operation);

      while (next.length < definition.operandCount) {
        next.push('');
      }

      return next;
    });
    setValidationError(null);
    onResetRemoteState();
  }, [operation]);

  function handleOperandChange(index: number, value: string) {
    setOperandInputs((current) => {
      const next = [...current];
      next[index] = value;
      return next;
    });

    if (validationError) {
      setValidationError(null);
    }
  }

  function handleTokenChange(value: string) {
    setToken(value);
    window.sessionStorage.setItem(tokenStorageKey, value);

    if (validationError) {
      setValidationError(null);
    }
  }

  function handleClientIDChange(value: string) {
    setClientID(value);
    window.sessionStorage.setItem(clientIDStorageKey, value);

    if (validationError) {
      setValidationError(null);
    }
  }

  async function handleGenerateToken() {
    if (clientID.trim() === '' || clientSecret.trim() === '') {
      setValidationError('Please provide the client ID and client secret to generate a token.');
      return;
    }

    setIsGeneratingToken(true);
    setValidationError(null);

    try {
      const response = await issueToken({
        client_id: clientID.trim(),
        client_secret: clientSecret,
      });

      handleTokenChange(response.access_token);
    } catch (error) {
      const message =
        error instanceof HttpError ? error.message : 'Could not generate a token from the API.';
      setValidationError(message);
    } finally {
      setIsGeneratingToken(false);
    }
  }

  async function runCalculation() {
    if (token.trim() === '') {
      setValidationError('Please provide a bearer token before calling the API.');
      return;
    }

    const nextError = validateOperandInputs(operation, operandInputs);
    if (nextError) {
      setValidationError(nextError);
      return;
    }

    setValidationError(null);
    await onSubmit({
      operation,
      operands: toOperandNumbers(operation, operandInputs),
      token: token.trim(),
    });
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    await runCalculation();
  }

  return (
    <form className="calculator-form" onSubmit={handleSubmit}>
      <div className="credential-grid">
        <label className="field">
          <span className="field__label">Client ID</span>
          <input
            className="field__control"
            type="text"
            placeholder="sezzle-demo-client"
            value={clientID}
            onChange={(event) => handleClientIDChange(event.target.value)}
          />
        </label>
        <label className="field">
          <span className="field__label">Client secret</span>
          <input
            className="field__control"
            type="password"
            placeholder="Paste the demo client secret"
            value={clientSecret}
            onChange={(event) => setClientSecret(event.target.value)}
          />
        </label>
      </div>
      <div className="button-row">
        <button
          className="button button--secondary"
          type="button"
          disabled={isGeneratingToken}
          onClick={() => {
            void handleGenerateToken();
          }}
        >
          {isGeneratingToken ? 'Generating token...' : 'Generate token'}
        </button>
      </div>
      <label className="field">
        <span className="field__label">Bearer token</span>
        <textarea
          className="field__control field__control--token"
          placeholder="Paste a token generated from the backend helper"
          rows={4}
          value={token}
          onChange={(event) => handleTokenChange(event.target.value)}
        />
      </label>
      <OperationSelector value={operation} onChange={setOperation} />
      <p className="form-help">{getOperationDefinition(operation).helpText}</p>
      <OperandInputs operation={operation} values={operandInputs} onChange={handleOperandChange} />
      <ErrorAlert message={validationError ?? error} />
      <div className="button-row">
        <button
          className="button button--primary"
          type="button"
          disabled={isLoading || isGeneratingToken}
          onClick={() => {
            void runCalculation();
          }}
        >
          {isLoading ? 'Calculating...' : 'Calculate'}
        </button>
      </div>
    </form>
  );
}
