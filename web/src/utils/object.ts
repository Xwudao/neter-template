/**
 * Delete specified props from an object (shallow clone).
 */
const deleteProps = <T extends object>(obj: T, props: (keyof T)[]): Partial<T> => {
  const newObj = { ...obj }
  props.forEach((prop) => {
    delete newObj[prop]
  })
  return newObj
}

export { deleteProps }
