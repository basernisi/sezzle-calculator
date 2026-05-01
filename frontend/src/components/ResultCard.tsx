interface ResultCardProps {
  result: number | null;
}

export function ResultCard({ result }: ResultCardProps) {
  return (
    <section className="result-card" aria-live="polite">
      <p className="result-card__label">Result</p>
      <p className="result-card__value">{result === null ? 'Awaiting calculation' : result}</p>
    </section>
  );
}
