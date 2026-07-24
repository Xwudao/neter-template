function openUrl(url: string, target: boolean = true) {
  const a = document.createElement('a')
  a.href = url
  a.setAttribute('rel', 'noreferrer')
  a.setAttribute('target', target ? '_blank' : '_self')
  a.click()
}

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export { openUrl, sleep }
