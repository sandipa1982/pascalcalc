import { useState } from 'react'
import { evaluateExpression } from '../api/calculatorApi'
import './ClassicCalculator.css'

export const MAX_CLASSIC_EXPRESSION_LENGTH = 256
const LENGTH_ERROR = `Maximum expression length is ${MAX_CLASSIC_EXPRESSION_LENGTH} characters.`

const keys = [
  'C', '(', ')', '÷',
  '7', '8', '9', '×',
  '4', '5', '6', '−',
  '1', '2', '3', '+',
  '√', '0', '.', '=',
  'xʸ', '%', '±',
]
const operators: Record<string, string> = {
  '÷': '/', '×': '*', '−': '-', '+': '+', 'xʸ': '^',
}
function ClassicCalculator() {
  const [input, setInput] = useState('')
  const [result, setResult] = useState('0')
  const [completed, setCompleted] = useState('')
  const [justEvaluated, setJustEvaluated] = useState(false)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  function press(key: string) {
    if (loading) return
    setError('')
    if (key === 'C') {
      setInput(''); setResult('0'); setCompleted(''); setJustEvaluated(false)
      return
    }
    if (key === '=') {
      const expression = input || (justEvaluated ? result : '')
      if (!expression) return
      if (expression.length > MAX_CLASSIC_EXPRESSION_LENGTH) {
        setError(LENGTH_ERROR)
        return
      }
      setLoading(true)
      void evaluateExpression(expression).then(response => {
        setResult(String(response.result))
        setCompleted(expression)
        setInput('')
        setJustEvaluated(true)
      }).catch((cause: unknown) => {
        setError(cause instanceof Error ? cause.message : 'Calculation failed')
      }).finally(() => setLoading(false))
      return
    }
    if (key === '±') {
      const base = justEvaluated ? result : input
      const next = base ? `-(${base})` : '-'
      if (next.length > MAX_CLASSIC_EXPRESSION_LENGTH) {
        setError(LENGTH_ERROR)
        return
      }
      setInput(next)
      setJustEvaluated(false); setCompleted('')
      return
    }
    const text = key === '√' ? 'sqrt(' : (operators[key] ?? key)
    const isOperator = key in operators
    const startsNew = justEvaluated && !isOperator && key !== '%' && key !== ')'
    const base = justEvaluated ? (startsNew ? '' : result) : input
    const next = isOperator && /[+*/^\-]$/.test(base)
      ? base.slice(0, -1) + text
      : base + text
    if (next.length > MAX_CLASSIC_EXPRESSION_LENGTH) {
      setError(LENGTH_ERROR)
      return
    }
    setInput(next)
    setJustEvaluated(false); setCompleted('')
  }
  return (
    <section className="classic-calculator">
      <div className="classic-display">
        <p className="classic-operation">
          {loading ? 'Calculating...' : (input || (completed ? `${completed} =` : '0'))}
        </p>
        <output className="classic-result" role="status">{input || result}</output>
      </div>
      {error && <p role="alert" className="error">{error}</p>}
      <div className="classic-keypad">
        {keys.map(key => (
          <button key={key} type="button" disabled={loading}
            className={`classic-key ${key === '=' ? 'classic-equals' : ''}`}
            onClick={() => press(key)}>{key}</button>
        ))}
      </div>
    </section>
  )
}
export default ClassicCalculator
