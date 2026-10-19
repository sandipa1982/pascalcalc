import { useState } from 'react'
import {
  calculateArithmetic,
  calculateMath,
  type ArithmeticOperation,
  type MathOperation,
} from '../api/calculatorApi'
import '../App.css'

type Operation = ArithmeticOperation | MathOperation

const arithmeticOperations: readonly Operation[] = [
  'add',
  'subtract',
  'multiply',
  'divide',
]
const operations = [
  { value: 'add', label: '+', title: 'Addition' },
  { value: 'subtract', label: '−', title: 'Subtraction' },
  { value: 'multiply', label: '×', title: 'Multiplication' },
  { value: 'divide', label: '÷', title: 'Division' },
  { value: 'power', label: 'xʸ', title: 'Power' },
  { value: 'sqrt', label: '√x', title: 'Square Root' },
  { value: 'percentage', label: '%', title: 'Percentage' },
] as const

function isArithmeticOperation(
  operation: Operation
): operation is ArithmeticOperation {
  return arithmeticOperations.includes(operation)
}

function PrecisionCalculator() {
  const [a, setA] = useState('20')
  const [b, setB] = useState('5')
  const [operation, setOperation] = useState<Operation>('add')
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const requiresSecondOperand = operation !== 'sqrt'

const MAX_DIGITS = 15

function handleNumberInput(
  value: string,
  setter: (value: string) => void
) {
  // Allow optional minus, digits, and only one decimal point.
  if (!/^-?\d*(\.\d*)?$/.test(value)) {
    return
  }

  // Count digits only, excluding '-' and '.'.
  const digitCount = (value.match(/\d/g) || []).length

  if (digitCount > MAX_DIGITS) {
    return
  }

  setter(value)
}

  async function handleCalculate() {
    setError('')
    setResult(null)

    if (
      a.trim() === '' ||
      (requiresSecondOperand && b.trim() === '')
    ) {
      setError(
        requiresSecondOperand
          ? 'Please enter both numbers.'
          : 'Please enter the first number.'
      )
      return
    }

    const first = Number(a)
    const second = Number(b)

    if (
      !Number.isFinite(first) ||
      (requiresSecondOperand && !Number.isFinite(second))
    ) {
      setError('Please enter valid numbers.')
      return
    }

    if (operation === 'divide' && second === 0) {
      setError('Cannot divide by zero.')
      return
    }

    if (operation === 'sqrt' && first < 0) {
      setError('Cannot calculate square root of a negative number.')
      return
    }

    setLoading(true)

    try {
      let value: number

      if (isArithmeticOperation(operation)) {
        const response = await calculateArithmetic(
          operation,
          first,
          second
        )
        value = response.result
      } else {
        const response = await calculateMath(
          operation,
          first,
          requiresSecondOperand ? second : undefined
        )
        value = response.result
      }

      setResult(value)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'Something went wrong.'
      )
    } finally {
      setLoading(false)
    }
  }

  return (
    <section className="calculator">
      <h2>Precision Calculator</h2>
      <p className="subtitle">Calculate with Precision</p>

      <label htmlFor="number-a">
        Operand 1 (Maximum 15 digits)
      </label>

      <input
        id="number-a"
        type="text"
        inputMode="decimal"
        value={a}
        onChange={event =>
          handleNumberInput(event.target.value, setA)
        }
        placeholder="Enter first number"
      />

<div className="operation-section">
  <label className="operation-label">Select operation</label>

  <div className="operation-buttons">
    {operations.map(item => (
      <button
        key={item.value}
        type="button"
        className={`operation-btn ${
          operation === item.value ? 'active' : ''
        }`}
        aria-label={item.title}
        aria-pressed={operation === item.value}
        title={item.title}
        onClick={() => {
          setOperation(item.value)
          setResult(null)
          setError('')
        }}
      >
        {item.label}
      </button>
    ))}
  </div>
</div>

    {requiresSecondOperand && (
      <>
        <label htmlFor="number-b">
          Operand 2 (Maximum 15 digits)
        </label>

        <input
          id="number-b"
          type="text"
          inputMode="decimal"
          value={b}
          onChange={event =>
            handleNumberInput(event.target.value, setB)
          }
          placeholder="Enter second number"
        />
      </>
    )}

      <button
        type="button"
        onClick={handleCalculate}
        disabled={loading}
      >
        {loading ? 'Calculating...' : 'Calculate'}
      </button>

      {result !== null && (
        <p className="result" role="status">
          Result: {result}
        </p>
      )}

      {error && (
        <p className="error" role="alert">
          {error}
        </p>
      )}
    </section>
  )
}

export default PrecisionCalculator