import type { ApiError } from '../types/calculator';

const defaultHeaders = {
  'Content-Type': 'application/json',
};

export class HttpError extends Error {
  code: string;

  constructor(message: string, code = 'HTTP_ERROR') {
    super(message);
    this.code = code;
  }
}

export async function postJSON<TResponse>(
  url: string,
  body: unknown,
  token: string,
): Promise<TResponse> {
  const headers: Record<string, string> = {
    ...defaultHeaders,
  };
  if (token.trim() !== '') {
    headers.Authorization = `Bearer ${token}`;
  }

  const response = await fetch(url, {
    method: 'POST',
    headers,
    body: JSON.stringify(body),
  });

  const payload = (await response.json().catch(() => null)) as TResponse | ApiError | null;
  if (!response.ok) {
    const error = payload as ApiError | null;
    throw new HttpError(
      error?.error?.message ?? 'The request could not be completed.',
      error?.error?.code ?? 'HTTP_ERROR',
    );
  }

  return payload as TResponse;
}
