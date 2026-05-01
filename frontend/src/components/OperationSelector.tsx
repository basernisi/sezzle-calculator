import { operationDefinitions } from '../utils/validation';
import type { Operation } from '../types/calculator';

interface OperationSelectorProps {
  value: Operation;
  onChange: (value: Operation) => void;
}

export function OperationSelector({ value, onChange }: OperationSelectorProps) {
  return (
    <label className="field">
      <span className="field__label">Operation</span>
      <select
        className="field__control"
        value={value}
        onChange={(event) => onChange(event.target.value as Operation)}
      >
        {operationDefinitions.map((operation) => (
          <option key={operation.value} value={operation.value}>
            {operation.label}
          </option>
        ))}
      </select>
    </label>
  );
}
