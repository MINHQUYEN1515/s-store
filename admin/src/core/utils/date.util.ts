export function formatDate(date: Date | string, locale = 'vi-VN'): string {
  const value = typeof date === 'string' ? new Date(date) : date
  return value.toLocaleDateString(locale)
}
