
import { afterEach, describe, expect, it, vi } from 'vitest'
import {
  cleanup,
  fireEvent,
  render,
  screen,
} from '@testing-library/react'
import '@testing-library/jest-dom/vitest'
import App from './App'

afterEach(() => {
  cleanup()
  vi.unstubAllGlobals()
  vi.restoreAllMocks()
})

// Open Precision Mode because Classic Mode is the default.
function renderPrecisionCalculator() {
  render(<App />)

  fireEvent.click(
    screen.getByRole('button', { name: 'Precision Mode' })
  )
}

// Enter a value into a Precision Calculator operand.
function enterOperand(
  operand: 1 | 2,
  value: string
) {
  fireEvent.change(
    screen.getByLabelText(
      new RegExp(`Operand ${operand}`, 'i')
    ),
    { target: { value } }
  )
}

// Select an operation using the new operation buttons.
function selectOperation(name: string) {
  fireEvent.click(
    screen.getByRole('button', { name })
  )
}

function clickCalculate() {
  fireEvent.click(
    screen.getByRole('button', { name: 'Calculate' })
  )
}

describe('PascalCalc - Precision Mode', () => {
  it('renders the calculator form', () => {
    renderPrecisionCalculator()

    expect(
      screen.getByRole('heading', { name: 'PascalCalc' })
    ).toBeInTheDocument()

    expect(
      screen.getByRole('heading', {
        name: 'Precision Calculator',
      })
    ).toBeInTheDocument()

    expect(
      screen.getByLabelText(/Operand 1/i)
    ).toBeInTheDocument()

    expect(
      screen.getByLabelText(/Operand 2/i)
    ).toBeInTheDocument()

    expect(
      screen.getByRole('button', { name: 'Addition' })
    ).toBeInTheDocument()

    expect(
      screen.getByRole('button', { name: 'Calculate' })
    ).toBeInTheDocument()
  })

  it('calculates addition using the Arithmetic API', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      status: 200,
      json: async () => ({ result: 25 }),
    })

    vi.stubGlobal('fetch', fetchMock)

    renderPrecisionCalculator()

    selectOperation('Addition')
    enterOperand(1, '20')
    enterOperand(2, '5')

    clickCalculate()

    expect(
      await screen.findByRole('status')
    ).toHaveTextContent('Result: 25')

    expect(fetchMock).toHaveBeenCalledWith(
      '/api/v1/arithmetic/add',
      expect.objectContaining({
        method: 'POST',
        headers: expect.objectContaining({
          'Content-Type': 'application/json',
        }),
        body: JSON.stringify({
          a: 20,
          b: 5,
        }),
      })
    )
  })

  it('calculates square root using the Math API', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      status: 200,
      json: async () => ({ result: 9 }),
    })

    vi.stubGlobal('fetch', fetchMock)

    renderPrecisionCalculator()

    selectOperation('Square Root')
    enterOperand(1, '81')

    // Square root requires only one operand.
    expect(
      screen.queryByLabelText(/Operand 2/i)
    ).not.toBeInTheDocument()

    clickCalculate()

    expect(
      await screen.findByRole('status')
    ).toHaveTextContent('Result: 9')

    expect(fetchMock).toHaveBeenCalledWith(
      '/api/v1/math/sqrt',
      expect.objectContaining({
        method: 'POST',
        body: JSON.stringify({
          a: 81,
        }),
      })
    )
  })

  it('rejects division by zero without calling the API', async () => {
    const fetchMock = vi.fn()

    vi.stubGlobal('fetch', fetchMock)

    renderPrecisionCalculator()

    selectOperation('Division')
    enterOperand(1, '20')
    enterOperand(2, '0')

    clickCalculate()

    expect(
      await screen.findByRole('alert')
    ).toHaveTextContent('Cannot divide by zero.')

    expect(fetchMock).not.toHaveBeenCalled()
  })

  it('displays an error returned by the backend', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: false,
      status: 400,
      json: async () => ({
        error: 'Calculation failed',
      }),
    })

    vi.stubGlobal('fetch', fetchMock)

    renderPrecisionCalculator()

    selectOperation('Addition')
    enterOperand(1, '20')
    enterOperand(2, '5')

    clickCalculate()

    expect(
      await screen.findByRole('alert')
    ).toHaveTextContent('Calculation failed')

    expect(fetchMock).toHaveBeenCalledOnce()
  })
})