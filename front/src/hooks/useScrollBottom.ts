import { useCallback, useEffect, useRef, useState } from 'react'

export const useScrollBottom = (list: unknown[]) => {
  const mounted = useRef<boolean>(false)
  const scrollAreaRef = useRef<HTMLUListElement>(null)
  const endItemRef = useRef<HTMLLIElement>(null)
  const [isBottom, setIsBottom] = useState(true)
  const [initialized, setInitialized] = useState(false)

  const onScroll = useCallback(() => {
    if (scrollAreaRef && scrollAreaRef.current) {
      setIsBottom(
        scrollAreaRef.current.scrollHeight - scrollAreaRef.current.scrollTop <=
          // bottom付近にスクロール位置があったらbottom判定したいので+numしている
          scrollAreaRef.current.clientHeight + 10,
      )
    }
  }, [scrollAreaRef, setIsBottom])

  useEffect(() => {
    if (scrollAreaRef && scrollAreaRef.current) {
      scrollAreaRef.current.addEventListener('scroll', onScroll)
      const copyScrollAreaRef = scrollAreaRef
      return () => {
        copyScrollAreaRef &&
          copyScrollAreaRef.current &&
          copyScrollAreaRef.current.removeEventListener('scroll', onScroll)
      }
    }
  }, [scrollAreaRef, onScroll])

  const scrollBottom = useCallback(
    (behavior: ScrollBehavior = 'smooth') => {
      if (!(scrollAreaRef && scrollAreaRef.current)) return
      if (!(endItemRef && endItemRef.current)) return

      const bottomPos =
        scrollAreaRef.current.scrollHeight - endItemRef.current.scrollHeight
      scrollAreaRef.current.scroll({
        top: bottomPos,
        behavior,
      })
    },
    [scrollAreaRef, endItemRef],
  )

  /* eslint react-hooks/exhaustive-deps: 0 */
  useEffect(() => {
    if (!mounted.current) {
      scrollBottom('auto')
      mounted.current = true
      setInitialized(true)
      return
    }

    if (isBottom) {
      scrollBottom()
    }
  }, [list.length, scrollBottom])

  return { scrollAreaRef, endItemRef, isBottom, initialized }
}
