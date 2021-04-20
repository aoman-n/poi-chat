export const getRoomIdParam = (id: string) => {
  const idParts = id.split(':')
  return idParts[1]
}
