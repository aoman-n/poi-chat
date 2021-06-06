import { useEffect, useState } from 'react'

export const usePrevScroll = (
  el: React.RefObject<HTMLUListElement>,
  list: unknown[],
  enabled: boolean,
) => {
  const [oldScrollHeight, setOldScrollHeight] = useState(0)

  useEffect(() => {
    if (!enabled) return
    if (el && el.current) {
      el.current.scrollTop = el.current.scrollHeight - oldScrollHeight
    }
  }, [list.length, el, oldScrollHeight, enabled])

  return {
    setOldScrollHeight,
  }
}
