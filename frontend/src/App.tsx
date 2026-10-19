import { useState } from 'react'
import ClassicCalculator from './components/ClassicCalculator'
import PrecisionCalculator from './components/PrecisionCalculator'

type CalculatorMode = 'classic' | 'precision'

function App() {
  const [mode, setMode] = useState<CalculatorMode>('classic')

  return (
    <main>
      <header>
        <h1>PascalCalc</h1>
        <p>Classic inspiration. Modern calculation.</p>
      </header>

      <nav aria-label="Calculator modes">
        <button
          type="button"
          aria-pressed={mode === 'classic'}
          onClick={() => setMode('classic')}
        >
          Classic Mode
        </button>

        <button
          type="button"
          aria-pressed={mode === 'precision'}
          onClick={() => setMode('precision')}
        >
          Precision Mode
        </button>
      </nav>

      {mode === 'classic' ? (
        <ClassicCalculator />
      ) : (
        <PrecisionCalculator />
      )}
    </main>
  )
}

export default App