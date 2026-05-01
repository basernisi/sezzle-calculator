import { CalculatorForm } from './components/CalculatorForm';
import { ResultCard } from './components/ResultCard';
import { useCalculator } from './hooks/useCalculator';

export default function App() {
  const { result, error, isLoading, submit, reset } = useCalculator();

  return (
    <main className="app-shell">
      <section className="hero-panel">
        <p className="eyebrow">Senior Engineer Take-Home</p>
        <h1>Secure calculator with a clean service boundary.</h1>
        <p className="hero-copy">
          A focused full-stack calculator that exercises validation, structured API errors,
          token-protected requests, and maintainable UI composition.
        </p>
      </section>
      <section className="workspace-card">
        <CalculatorForm
          isLoading={isLoading}
          error={error}
          onSubmit={submit}
          onResetRemoteState={reset}
        />
        <ResultCard result={result} />
      </section>
    </main>
  );
}
