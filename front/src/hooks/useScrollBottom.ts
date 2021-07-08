import { useRef, useEffect, useLayoutEffect } from 'react'
import { isExistsRef } from '@/utils/elements'
import { useObserver } from './useObserver'

/* eslint react-hooks/exhaustive-deps: 0 */
export const useScrollBottom = (list: unknown[]) => {
  const parentRef = useRef<HTMLUListElement>(null)
  const [isBottom, scrollBottomRef] = useObserver()

  // 初回レンダリング時にスクロールバーを最下部に設定する
  useLayoutEffect(() => {
    if (isExistsRef(parentRef)) {
      const bottom =
        parentRef.current.scrollHeight - parentRef.current.clientHeight
      parentRef.current.scroll({ top: bottom })
    }
  }, [])

  // 新規アイテムが追加されたら最下部へ自動スクロール
  useEffect(() => {
    if (isBottom && isExistsRef(parentRef)) {
      const bottom =
        parentRef.current.scrollHeight - parentRef.current.clientHeight
      parentRef.current.scroll({
        top: bottom,
        behavior: 'smooth',
      })
    }
  }, [list.length])

  return { scrollBottomRef, parentRef }
}
