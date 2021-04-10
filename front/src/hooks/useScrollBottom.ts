import { useCallback, useEffect, useRef, useState } from 'react'

const useScrollBottom = (list: unknown[]) => {
  const mounted = useRef<boolean>(false)
  const scrollAreaRef = useRef<HTMLUListElement>(null)
  const endItemRef = useRef<HTMLLIElement>(null)
  const [isBottom, setIsBottom] = useState(true)

  const onScroll = useCallback(() => {
    if (scrollAreaRef && scrollAreaRef.current) {
      setIsBottom(
        scrollAreaRef.current.scrollHeight - scrollAreaRef.current.scrollTop <=
          scrollAreaRef.current.clientHeight,
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

  const scrollBottom = useCallback(() => {
    if (!(scrollAreaRef && scrollAreaRef.current)) return
    if (!(endItemRef && endItemRef.current)) return

    const bottomPos =
      scrollAreaRef.current.scrollHeight - endItemRef.current.scrollHeight
    scrollAreaRef.current.scroll({
      top: bottomPos,
      behavior: 'smooth',
    })
  }, [scrollAreaRef, endItemRef])

  useEffect(() => {
    if (!mounted.current) {
      scrollBottom()
      mounted.current = true
    }
    if (isBottom) {
      scrollBottom()
    }
  }, [list.length, scrollBottom, isBottom])

  return { scrollAreaRef, endItemRef }
}

export default useScrollBottom
