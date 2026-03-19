import { useEffect, useState } from 'react';

type HealthResponse = { status: string; service: string }
type HitResponse = { count: number; hostname: string }

const apiBase = 'http://localhost:4545';

function App() {
  const [health, setHealth] = useState<HealthResponse | null>(null)
  const [hit, setHit] = useState<HitResponse | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetch(`${apiBase}/health`)
      .then(res => res.json())
      .then(data => {
        setHealth(data)
      })
      .catch(err => setError(`Error: ${err.message}`))
  }, [])

  const fetchHit = () => {
    fetch(`${apiBase}/hit`)
      .then(res => res.json())
      .then((data: HitResponse) => {
        setHit(data)
        setError(null)
      })
      .catch(err => {
        setError(`Error: ${err.message}`)
        setHit(null)
      })
  }

  return (
    <div>
      <h1>CloudNativeApp v7- Hello World</h1>
      <p>This is a minimal React app for learning cloud-native workflows.</p>

      <div style={{ marginTop: '2rem', padding: '1rem', background: '#1a1a1a', borderRadius: '8px' }}>
        <h2>Backend Health Check</h2>
        {health && (
          <div>
            <p>Status: <strong>{health.status}</strong></p>
            <p>Service: <strong>{health.service}</strong></p>
          </div>
        )}
        {error && <p style={{ color: '#ff6b6b' }}>{error}</p>}
        {!health && !error && <p>Checking...</p>}
      </div>

      <div style={{ marginTop: '2rem', padding: '1rem', background: '#1a1a1a', borderRadius: '8px' }}>
        <h2>Hit Counter</h2>
        <button onClick={fetchHit}>Hit</button>
        {hit && (
          <div>
            <p>Count: <strong>{hit.count}</strong></p>
            <p>Hostname: <strong>{hit.hostname}</strong></p>
          </div>
        )}
      </div>

      <div style={{ marginTop: '2rem', fontSize: '0.9rem', color: '#888' }}>
        <p>Backend API: {apiBase}</p>
        <p>Frontend: http://localhost:3000</p>
      </div>
    </div>
  )
}

export default App