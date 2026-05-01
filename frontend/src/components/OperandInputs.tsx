import type { Operation } from '../types/calculator';
import { getOperationDefinition } from '../utils/validation';

interface OperandInputsProps {
  operation: Operation;
  values: string[];
  onChange: (index: number, value: string) => void;
}

export function OperandInputs({ operation, values, onChange }: OperandInputsProps) {
  const definition = getOperationDefinition(operation);

  return (
    <div className="operand-grid">
      {definition.placeholders.map((placeholder, index) => (
        <label key={`${operation}-${index}`} className="field">
          <span className="field__label">Operand {index + 1}</span>
          <input
            className="field__control"
            type="number"
            inputMode="decimal"
            step="any"
            placeholder={placeholder}
            value={values[index] ?? ''}
            onChange={(event) => onChange(index, event.target.value)}
          />
        </label>
      ))}
    </div>
  );
}
