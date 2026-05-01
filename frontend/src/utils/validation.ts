import type { Operation, OperationDefinition } from '../types/calculator';

export const operationDefinitions: OperationDefinition[] = [
  {
    value: 'add',
    label: 'Addition',
    operandCount: 2,
    placeholders: ['First number', 'Second number'],
    helpText: 'Adds two numbers together.',
  },
  {
    value: 'subtract',
    label: 'Subtraction',
    operandCount: 2,
    placeholders: ['Minuend', 'Subtrahend'],
    helpText: 'Subtracts the second number from the first.',
  },
  {
    value: 'multiply',
    label: 'Multiplication',
    operandCount: 2,
    placeholders: ['First factor', 'Second factor'],
    helpText: 'Multiplies two numbers.',
  },
  {
    value: 'divide',
    label: 'Division',
    operandCount: 2,
    placeholders: ['Dividend', 'Divisor'],
    helpText: 'Divides the first number by the second.',
  },
  {
    value: 'power',
    label: 'Exponentiation',
    operandCount: 2,
    placeholders: ['Base', 'Exponent'],
    helpText: 'Raises the first number to the power of the second.',
  },
  {
    value: 'sqrt',
    label: 'Square Root',
    operandCount: 1,
    placeholders: ['Number'],
    helpText: 'Returns the square root of a number.',
  },
  {
    value: 'percentage',
    label: 'Percentage',
    operandCount: 2,
    placeholders: ['Percent', 'Value'],
    helpText: 'Calculates percent of a value.',
  },
];

export function getOperationDefinition(operation: Operation): OperationDefinition {
  const definition = operationDefinitions.find((item) => item.value === operation);
  if (!definition) {
    throw new Error(`Unsupported operation: ${operation}`);
  }

  return definition;
}

export function sanitizeOperandInputs(values: string[], operation: Operation): string[] {
  return values.slice(0, getOperationDefinition(operation).operandCount);
}

export function validateOperandInputs(operation: Operation, values: string[]): string | null {
  const definition = getOperationDefinition(operation);
  const trimmedValues = values.slice(0, definition.operandCount);

  if (trimmedValues.length !== definition.operandCount || trimmedValues.some((value) => value.trim() === '')) {
    return `Please enter ${definition.operandCount} valid operand${definition.operandCount === 1 ? '' : 's'}.`;
  }

  const numbers = trimmedValues.map((value) => Number(value));
  if (numbers.some((value) => Number.isNaN(value) || !Number.isFinite(value))) {
    return 'Operands must be valid finite numbers.';
  }

  if (operation === 'divide' && numbers[1] === 0) {
    return 'Division by zero is not allowed.';
  }

  if (operation === 'sqrt' && numbers[0] < 0) {
    return 'Square root of a negative number is not allowed.';
  }

  return null;
}

export function toOperandNumbers(operation: Operation, values: string[]): number[] {
  return sanitizeOperandInputs(values, operation).map((value) => Number(value));
}
