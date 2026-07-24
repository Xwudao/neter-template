function truncateStr(str: string, length: number, suffix = '...'): string {
  if (str.length <= length) {
    return str
  }

  let byteLength = 0
  for (let i = 0; i < str.length; i++) {
    const charCode = str.charCodeAt(i)
    if (charCode >= 0x4e00 && charCode <= 0x9fa5) {
      byteLength += 2
    } else {
      byteLength += 1
    }
  }

  if (byteLength <= length) {
    return str
  }

  let newStr = ''
  let i = 0
  while (byteLength > length) {
    const charCode = str.charCodeAt(i)
    if (charCode >= 0x4e00 && charCode <= 0x9fa5) {
      byteLength -= 2
    } else {
      byteLength -= 1
    }
    newStr += str.charAt(i)
    i++
  }
  return newStr + suffix
}

const randomStr = (l: number) => {
  let str = ''
  while (str.length < l) {
    str += Math.random().toString(36).substring(2)
  }
  return str.substring(0, l)
}

const removeStrDupLines = (str: string) => {
  const lines = str.split('\n')
  const newLines = Array.from(new Set(lines))
  return newLines.join('\n')
}

const removeHtml = (text: string) => {
  return text.replace(/<[^>]+>/g, '')
}

export { truncateStr, randomStr, removeStrDupLines, removeHtml }
