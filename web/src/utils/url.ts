const urlRegex = /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)/gi;

function extractUrls(text: string, filter: (u: string) => boolean): string[] {
  const urls: string[] = [];
  let match;

  while ((match = urlRegex.exec(text)) !== null) {
    // urls.push(match[0]);
    // 移除末尾的 / 符号
    const url = match[0].replace(/\/$/, '');
    if (filter(url)) {
      urls.push(url);
    }
  }

  return removeDuplicates(urls);
}

function removeDuplicates<T>(array: T[]): T[] {
  return Array.from(new Set(array));
}

export { extractUrls };
