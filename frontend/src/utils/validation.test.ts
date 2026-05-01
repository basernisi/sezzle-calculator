import { describe, expect, it } from 'vitest';
import { toOperandNumbers, validateOperandInputs } from './validation';

describe('validation', () => {
  it('rejects division by zero', () => {
    expect(validateOperandInputs('divide', ['10', '0'])).toBe('Division by zero is not allowed.');
  });

  it('rejects negative square root', () => {
    expect(validateOperandInputs('sqrt', ['-9'])).toBe(
      'Square root of a negative number is not allowed.',
    );
  });

  it('converts operand values to numbers', () => {
    expect(toOperandNumbers('add', ['10', '5'])).toEqual([10, 5]);
  });
});
