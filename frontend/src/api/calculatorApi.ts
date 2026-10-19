import { appConfig } from '../config/appConfig'
import { apiRoutes } from '../config/apiRoutes'

export type ArithmeticOperation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'

export type MathOperation =
  | 'power'
  | 'sqrt'
  | 'percentage'

export interface CalculationResponse {
  result: number
}

export interface ExpressionResponse {
  expression: string
  result: number
}

async function postJson<T>(
  url: string,
  payload: object
): Promise<T> {
  const controller = new AbortController()
  const timeout = setTimeout(
    () => controller.abort(),
    appConfig.apiTimeoutMs
  )

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
      signal: controller.signal,
    })

    const data: unknown = await response.json()

    if (!response.ok) {
      if (
        typeof data === 'object' &&
        data !== null &&
        'error' in data &&
        typeof data.error === 'string'
      ) {
        throw new Error(data.error)
      }

      throw new Error(`Request failed: HTTP ${response.status}`)
    }

    return data as T
  } catch (error) {
    if (
      error instanceof Error &&
      error.name === 'AbortError'
    ) {
      throw new Error('Request timed out. Please try again.')
    }

    throw error
  } finally {
    clearTimeout(timeout)
  }
}

export async function calculateArithmetic(
  operation: ArithmeticOperation,
  a: number,
  b: number
): Promise<CalculationResponse> {
  return postJson<CalculationResponse>(
    apiRoutes.arithmetic[operation],
    { a, b }
  )
}

export async function calculateMath(
  operation: MathOperation,
  a: number,
  b?: number
): Promise<CalculationResponse> {
  const payload =
    operation === 'sqrt'
      ? { a }
      : { a, b }

  return postJson<CalculationResponse>(
    apiRoutes.math[operation],
    payload
  )
}

export async function evaluateExpression(
  expression: string
): Promise<ExpressionResponse> {
  return postJson<ExpressionResponse>(
    apiRoutes.expression.evaluate,
    { expression }
  )
}