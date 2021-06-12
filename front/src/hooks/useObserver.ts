import { useState, useRef, useEffect } from 'react'

export const useObserver = (): [boolean, React.RefObject<HTMLDivElement>] => {
  const target = useRef<HTMLDivElement>(null)
  const [intersect, setIntersect] = useState(false)

  useEffect(() => {
    if (!target.current) return

    let observer: IntersectionObserver
    if (target.current.parentElement) {
      observer = new IntersectionObserver(
        ([entry]) => setIntersect(entry.isIntersecting),
        { root: target.current.parentElement, rootMargin: '10px' },
      )
      observer.observe(target.current)
    } else {
      setIntersect(false)
    }

    return () => {
      if (observer) {
        observer.disconnect()
      }
    }
  }, [target])

  return [intersect, target]
}
