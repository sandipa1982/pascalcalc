import { afterEach, describe, expect, it, vi } from 'vitest'
import { cleanup, fireEvent, render, screen, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom/vitest'
import ClassicCalculator, { MAX_CLASSIC_EXPRESSION_LENGTH } from './ClassicCalculator'
import { evaluateExpression } from '../api/calculatorApi'
vi.mock('../api/calculatorApi', () => ({ evaluateExpression: vi.fn() }))
afterEach(() => { cleanup(); vi.clearAllMocks() })
const press = (...keys: string[]) => keys.forEach(key =>
  fireEvent.click(screen.getByRole('button', { name: key })))
describe('ClassicCalculator expression entry', () => {
  it('renders arithmetic, parentheses and square root controls', () => {
    render(<ClassicCalculator />)
    for (const key of ['+', '√', '(', ')', '=']) {
      expect(screen.getByRole('button', { name: key })).toBeInTheDocument()
    }
  })
  it('evaluates ordinary arithmetic', async () => {
    vi.mocked(evaluateExpression).mockResolvedValue({ expression: '2+3', result: 5 })
    render(<ClassicCalculator />)
    press('2', '+', '3', '=')
    await waitFor(() => expect(evaluateExpression).toHaveBeenCalledWith('2+3'))
    await waitFor(() => expect(screen.getByRole('status')).toHaveTextContent('5'))
  })
  it('builds and evaluates 2+4+sqrt(16)', async () => {
    vi.mocked(evaluateExpression).mockResolvedValue({ expression: '2+4+sqrt(16)', result: 10 })
    render(<ClassicCalculator />)
    press('2', '+', '4', '+', '√', '1', '6', ')')
    expect(screen.getByRole('status')).toHaveTextContent('2+4+sqrt(16)')
    expect(evaluateExpression).not.toHaveBeenCalled()
    press('=')
    await waitFor(() => expect(evaluateExpression).toHaveBeenCalledWith('2+4+sqrt(16)'))
    await waitFor(() => expect(screen.getByRole('status')).toHaveTextContent('10'))
  })
  it('supports nested square root', async () => {
    vi.mocked(evaluateExpression).mockResolvedValue({ expression: 'sqrt(sqrt(16))', result: 2 })
    render(<ClassicCalculator />)
    press('√', '√', '1', '6', ')', ')', '=')
    await waitFor(() => expect(evaluateExpression).toHaveBeenCalledWith('sqrt(sqrt(16))'))
  })
  it('shows service validation errors', async () => {
    vi.mocked(evaluateExpression).mockRejectedValue(new Error('square root of negative number'))
    render(<ClassicCalculator />)
    press('√', '−', '4', ')', '=')
    expect(await screen.findByRole('alert')).toHaveTextContent('square root of negative number')
  })
  it('accepts exactly 256 characters and rejects the next one', () => {
    render(<ClassicCalculator />)
    for (let i = 0; i < MAX_CLASSIC_EXPRESSION_LENGTH; i++) press('1')
    expect(screen.getByRole('status')).toHaveTextContent('1'.repeat(256))
    press('2')
    expect(screen.getByRole('alert')).toHaveTextContent('Maximum expression length is 256 characters.')
    expect(screen.getByRole('status')).toHaveTextContent('1'.repeat(256))
    expect(evaluateExpression).not.toHaveBeenCalled()
    press('C')
    expect(screen.queryByRole('alert')).not.toBeInTheDocument()
    expect(screen.getByRole('status')).toHaveTextContent('0')
  })
  it('counts sqrt( as five characters and rejects an overflowing function', () => {
    render(<ClassicCalculator />)
    for (let i = 0; i < 252; i++) press('1')
    press('√')
    expect(screen.getByRole('alert')).toHaveTextContent('Maximum expression length is 256 characters.')
    expect(screen.getByRole('status')).toHaveTextContent('1'.repeat(252))
  })
  it('does not allow sign change to exceed the limit', () => {
    render(<ClassicCalculator />)
    for (let i = 0; i < 255; i++) press('1')
    press('±')
    expect(screen.getByRole('alert')).toHaveTextContent('Maximum expression length is 256 characters.')
    expect(screen.getByRole('status')).toHaveTextContent('1'.repeat(255))
  })
it('rejects expressions exceeding 256 characters', () => {
  render(<ClassicCalculator />);

  const expression =
    '45+36+258+891+325+789-587+7896541+sqrt(1024)+789654123698745631478+36+' +
    '45+36+258+891+325+789-587+7896541+sqrt(1024)+789654123698745631478+36+' +
    '45+36+258+891+325+789-587+7896541+sqrt(1024)+789654123698745631478+36+' +
    '45+36+258+891+325+789-587+7896541+sqrt(1024)+789654123698745631478+36';

  expect(expression.length).toBeGreaterThan(
    MAX_CLASSIC_EXPRESSION_LENGTH
  );

  // Convert expression syntax to calculator button labels.
  // The √ button inserts "sqrt(", so remove the extra "(".
  const keys = expression
    .replaceAll('sqrt(', '√')
    .replaceAll('-', '−')
    .split('');

  keys.forEach(key => {
    fireEvent.click(screen.getByRole('button', { name: key }));
  });

  // The calculator must stop accepting input at 256 characters.
  const expectedExpression = expression.slice(
    0,
    MAX_CLASSIC_EXPRESSION_LENGTH
  );

  expect(screen.getByRole('status')).toHaveTextContent(
    expectedExpression
  );

  expect(screen.getByRole('alert')).toHaveTextContent(
    'Maximum expression length is 256 characters.'
  );

  expect(evaluateExpression).not.toHaveBeenCalled();
});
}); // Closes the outer describe() block