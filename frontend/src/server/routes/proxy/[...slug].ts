import { defineEventHandler, proxyRequest } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const goApiBase = config.public.goBase
  const slug = event.context.params?.slug || ''
  const targetPath = `/dashboard/${slug}`

  // 获取原始查询参数
  const originalQuery = getQuery(event)
  const url = new URL(targetPath, goApiBase)

  // 把原始查询参数拼到目标URL
  for (const [key, value] of Object.entries(originalQuery)) {
    if (Array.isArray(value)) {
      value.forEach(v => url.searchParams.append(key, v))
    } else if (value !== undefined) {
      url.searchParams.append(key, value as string)
    }
  }

  const targetUrl = url.toString()

  return proxyRequest(event, targetUrl)
})