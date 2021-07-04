import { useCallback, useEffect, useRef } from 'react'

export const useThrottleFn = <T>(
  callback: (args: T) => void,
  threshold: number,
) => {
  const wait = useRef(false)
  const timeout = useRef<NodeJS.Timeout>()

  useEffect(() => {
    return () => {
      if (timeout.current) {
        clearTimeout(timeout.current)
      }
    }
  }, []) // No need for deps here since 'timeout' is mutated

  return useCallback(
    (args: T) => {
      if (!wait.current) {
        callback(args)
        wait.current = true
        if (timeout.current) {
          clearTimeout(timeout.current)
        }

        timeout.current = setTimeout(() => {
          wait.current = false
        }, threshold)
      }
    },
    [callback, threshold],
  )
}
