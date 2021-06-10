export const merge = <T extends { id: string }>(x: T[], y: T[]): T[] => {
  const merged = [...x, ...y]

  return Array.from(
    merged
      .reduce((map, current) => map.set(current.id, current), new Map())
      .values(),
  )
}

export const sortById = <T extends { id: string }>(ary: T[]): T[] => {
  return ary.sort((a, b) => {
    if (a.id > b.id) {
      return -1
    }
    return 1
  })
}
