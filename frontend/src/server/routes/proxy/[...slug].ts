import { defineEventHandler, proxyRequest } from 'h3'

export default defineEventHandler(async (event) => {
  // 从运行时配置中获取 Go 后端的基础 URL
  const config = useRuntimeConfig()
  const goApiBase = config.public.goBase

  // `event.context.params.slug` 会捕获 /proxy/ 之后的所有路径部分
  // 例如，如果请求是 /proxy/dashboard/api/renew_token
  // `slug` 将会是 'dashboard/api/renew_token'
  const slug = event.context.params?.slug || ''

  // 构造目标路径，确保它以 '/' 开头
  const targetPath = `/dashboard/${slug}`

  // 构造完整的后端目标 URL
  const targetUrl = new URL(targetPath, goApiBase).toString()

  // console.log(`Proxying request from /proxy${targetPath} to ${targetUrl}`)
  // console.log('Request Method:', event.method) 

  // 使用 proxyRequest 将请求转发到 Go 后端
  // proxyRequest 会自动处理请求方法、头部（包括 Cookie）、请求体等
  // 并且会将后端响应（包括 Set-Cookie 头部）转发回客户端浏览器
  // console.log('event:', event)
  return proxyRequest(event, targetUrl, {
    // 这里可以传递一些额外的 fetch 选项给底层的 fetch 调用，如果需要的话
    // 例如：
    // fetchOptions: {
    //   headers: {
    //     'X-My-Custom-Proxy-Header': 'value'
    //   }
    // }
  })
})