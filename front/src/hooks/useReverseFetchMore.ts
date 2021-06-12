import { useState, useEffect, useRef } from 'react'
import { useObserver } from './useObserver'

/* eslint react-hooks/exhaustive-deps: 0 */
export const useReverseFetchMore = <P extends { id: string }>(
  list: P[],
  fetchMoreFn: () => void,
  enabled: boolean,
) => {
  const [isTop, scrollTopRef] = useObserver()
  const firstItemRef = useRef<HTMLLIElement>(null)
  const [prevFirstItem, setPrevFirstComment] = useState(list[0])

  // 最上部まで来たら読み込む
  useEffect(() => {
    if (isTop && enabled) {
      fetchMoreFn()
    }
  }, [isTop])

  // 取得後にスクロールを元の位置に戻す
  useEffect(() => {
    if (isTop) {
      firstItemRef?.current?.scrollIntoView()
      setPrevFirstComment(list[0])
    }
  }, [list.length])

  return { scrollTopRef, prevFirstItem, firstItemRef }
}
