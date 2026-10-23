import { afterEach, expect, it, vi } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import { createRef } from 'react'
import ShowSpeechWaveform from './Show-SpeechWaveform'
afterEach(() => vi.unstubAllGlobals())
it('draws a circular waveform using the streamed speech amplitude', async () => {
  const stream = new ReadableStream({ start(controller) { controller.enqueue(new TextEncoder().encode('data: {"voice":"speaking","level":0.64}\n\n')) } })
  vi.stubGlobal('fetch', vi.fn(async () => ({ ok: true, body: stream })))
  const { container } = render(<ShowSpeechWaveform voice="idle" theme={createRef()} />)
  expect(await screen.findByRole('img', { name: /Circular waveform/ })).toBeInTheDocument()
  const lines = container.querySelectorAll('line')
  expect(lines).toHaveLength(64)
  const last = lines[63], first = lines[0]
  const length = (line: Element) => Math.hypot(Number(line.getAttribute('x2')) - Number(line.getAttribute('x1')), Number(line.getAttribute('y2')) - Number(line.getAttribute('y1')))
  expect(length(last)).toBeGreaterThan(length(first))
  await waitFor(() => expect(fetch).toHaveBeenCalledWith('/api/voice-stream', expect.objectContaining({ headers: { 'X-Jarvis-Display': '1' } })))
})
