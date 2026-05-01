import { postJSON } from './client';
import type { CalculateRequest, CalculateResponse } from '../types/calculator';

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:18080';

export async function calculate(request: CalculateRequest, token: string): Promise<CalculateResponse> {
  return postJSON<CalculateResponse>(`${apiBaseUrl}/api/v1/calculate`, request, token);
}
