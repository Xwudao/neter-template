import dayjs from 'dayjs'
import duration from 'dayjs/plugin/duration'
import timezone from 'dayjs/plugin/timezone'
import utc from 'dayjs/plugin/utc'

dayjs.extend(duration)
dayjs.extend(utc)
dayjs.extend(timezone)

const shanghaiTimeZone = 'Asia/Shanghai'

const getShanghaiNow = () => dayjs().tz(shanghaiTimeZone)

const formatShanghaiDate = (date: Date | string | number = new Date()) =>
  dayjs(date).tz(shanghaiTimeZone).format('YYYY-MM-DD')

const formatDayTime = (day: string) => {
  if (day.match(/^\d+$/)) {
    return formatUnix(parseInt(day))
  }
  return dayjs(day).format('YYYY-MM-DD HH:mm')
}

const formatDate = (day: string) => {
  if (day.match(/^\d+$/)) {
    return formatUnix(parseInt(day))
  }
  return dayjs(day).format('YYYY-MM-DD')
}

const formatUnix = (unix: number) => {
  if (String(unix).length === 10) {
    unix *= 1000
  }
  return dayjs(unix).format('YYYY-MM-DD HH:mm')
}

/**
 * Format a duration between two ISO date strings into a human-readable form.
 * Returns `--` when neither date is available.
 * Returns a relative label when only started_at is set (task still running).
 */
function formatDuration(startedAt?: string | null, finishedAt?: string | null): string {
  if (!startedAt) return '--'

  const start = dayjs(startedAt)
  const end = finishedAt ? dayjs(finishedAt) : dayjs()
  const diffMs = end.diff(start)
  if (diffMs < 0) return '--'

  const d = dayjs.duration(diffMs)

  const hours = Math.floor(d.asHours())
  const minutes = d.minutes()
  const seconds = d.seconds()

  const parts: string[] = []
  if (hours > 0) parts.push(`${hours}h`)
  if (minutes > 0) parts.push(`${minutes}m`)
  parts.push(`${seconds}s`)

  return parts.join(' ')
}

export {
  formatDayTime,
  formatUnix,
  formatDate,
  formatDuration,
  formatShanghaiDate,
  getShanghaiNow,
}
