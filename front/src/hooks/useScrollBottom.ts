import { useEffect, useLayoutEffect } from 'react'
import { useObserver } from './useObserver'

/* eslint react-hooks/exhaustive-deps: 0 */
export const useScrollBottom = (list: unknown[]) => {
  const [isBottom, scrollBottomRef] = useObserver()

  // 初回レンダリング時にスクロールバーを最下部に設定する
  useLayoutEffect(() => {
    scrollBottomRef?.current?.scrollIntoView()
  }, [])

  // 新規アイテムが追加されたら最下部へ自動スクロール
  useEffect(() => {
    if (isBottom) {
      scrollBottomRef?.current?.scrollIntoView({ behavior: 'smooth' })
    }
  }, [list.length])

  return { scrollBottomRef }
}
