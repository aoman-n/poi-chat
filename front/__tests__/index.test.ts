import { merge, sortById } from '../src/utils/array'

type Item = {
  id: string
  name: string
}

describe('utils', () => {
  describe('array', () => {
    describe('#merge', () => {
      it('結合されること', () => {
        const x: Item[] = [
          { id: '1', name: '1' },
          { id: '2', name: '2' },
        ]

        const y: Item[] = [
          { id: '3', name: '3' },
          { id: '4', name: '4' },
        ]

        const actual = merge(x, y)
        const expected: Item[] = [
          { id: '1', name: '1' },
          { id: '2', name: '2' },
          { id: '3', name: '3' },
          { id: '4', name: '4' },
        ]

        expect(actual).toEqual(expected)
      })
      it('重複が削除されること', () => {
        const x: Item[] = [
          { id: '1', name: '1' },
          { id: '2', name: '2' },
        ]

        const y: Item[] = [
          { id: '1', name: '1' },
          { id: '2', name: '2' },
        ]

        const actual = merge(x, y)
        const expected: Item[] = [
          { id: '1', name: '1' },
          { id: '2', name: '2' },
        ]

        expect(actual).toEqual(expected)
      })
    })
    describe('#sortById', () => {
      it('idが新しい順にソートされること', () => {
        const ary: Item[] = [
          { id: '3', name: '3' },
          { id: '2', name: '2' },
          { id: '1', name: '1' },
          { id: '4', name: '4' },
        ]

        const actual = sortById(ary)
        const expected: Item[] = [
          { id: '4', name: '4' },
          { id: '3', name: '3' },
          { id: '2', name: '2' },
          { id: '1', name: '1' },
        ]

        expect(actual).toEqual(expected)
      })
    })
  })
})
