import { useEffect, useState, type RefObject } from 'react'

export default function ShowSpeechWaveform({ voice, theme }: { voice: string; theme: RefObject<HTMLIFrameElement | null> }) {
  const [audio, setAudio] = useState({ voice: '', samples: Array<number>(64).fill(0) })
  useEffect(() => {
    const controller = new AbortController()
    const preference = window.matchMedia?.('(prefers-reduced-motion: reduce)')
    const listen = async () => {
      while (!controller.signal.aborted) {
        try {
          const response = await fetch('/api/voice-stream', { headers: { 'X-Jarvis-Display': '1' }, signal: controller.signal })
          if (!response.ok || !response.body) throw new Error('Voice visualiser stream unavailable')
          const reader = response.body.getReader(), decoder = new TextDecoder()
          let pending = ''
          try {
            while (!controller.signal.aborted) {
              const { value, done } = await reader.read()
              if (done) break
              pending += decoder.decode(value, { stream: true })
              const events = pending.split('\n\n'); pending = events.pop() || ''
              for (const event of events) {
                if (!event.startsWith('data: ')) continue
                const frame = JSON.parse(event.slice(6)) as { voice: string; level: number }
                const level = Math.min(1, Math.max(0, Number(frame.level) || 0))
                setAudio(old => ({ voice: frame.voice, samples: preference?.matches ? Array<number>(64).fill(0.08) : [...old.samples.slice(1), level] }))
                theme.current?.contentWindow?.postMessage({ api_version: 1, type: 'voice', voice: frame.voice, level }, '*')
              }
            }
          } finally { await reader.cancel().catch(() => {}) }
        } catch { /* The static voice label remains available while reconnecting. */ }
        if (!controller.signal.aborted) await new Promise(resolve => setTimeout(resolve, 1000))
      }
    }
    void listen()
    return () => controller.abort()
  }, [theme])
  if ((audio.voice || voice) !== 'speaking') return null
  return <aside className="speech-waveform" aria-label="Jarvis speaking">
    <svg viewBox="0 0 240 240" role="img" aria-label="Circular waveform reacting to Jarvis’s voice">
      <defs><linearGradient id="voice-colour" x1="0" y1="0" x2="1" y2="1"><stop stopColor="#64d8ff" /><stop offset="1" stopColor="#bc97ff" /></linearGradient></defs>
      <circle cx="120" cy="120" r="72" fill="var(--display-bg)" stroke="var(--display-border)" />
      <g stroke="url(#voice-colour)" strokeWidth="3" strokeLinecap="round">{audio.samples.map((sample, i) => {
        const length = 4 + Math.min(1, Math.sqrt(sample) * 2.5) * 32
        const angle = i * Math.PI / 32
        return <line key={i} x1={120 + Math.sin(angle) * 80} y1={120 - Math.cos(angle) * 80} x2={120 + Math.sin(angle) * (80 + length)} y2={120 - Math.cos(angle) * (80 + length)} />
      })}</g>
      <text x="120" y="117" textAnchor="middle" fill="var(--display-text)" fontSize="20">Jarvis</text><text x="120" y="139" textAnchor="middle" fill="var(--display-muted)" fontSize="13">Speaking</text>
    </svg>
  </aside>
}
