export type Operation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'sqrt'
  | 'percentage';

export interface CalculateRequest {
  operation: Operation;
  operands: number[];
}

export interface CalculateResponse {
  result: number;
}

export interface ApiError {
 error: {
    code: string;
    message: string;
  };
}

export interface IssueTokenRequest {
  client_id: string;
  client_secret: string;
}

export interface IssueTokenResponse {
  access_token: string;
  token_type: string;
  expires_in: number;
}

export interface OperationDefinition {
  value: Operation;
  label: string;
  operandCount: 1 | 2;
  placeholders: string[];
  helpText: string;
}

export interface CalculatorSubmission {
  operation: Operation;
  operands: number[];
  token: string;
}
