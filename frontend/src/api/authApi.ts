import { postJSON } from './client';
import type { IssueTokenRequest, IssueTokenResponse } from '../types/calculator';

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:18080';

export async function issueToken(request: IssueTokenRequest): Promise<IssueTokenResponse> {
  return postJSON<IssueTokenResponse>(`${apiBaseUrl}/api/v1/auth/token`, request, '');
}
