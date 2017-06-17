const setCookie = (name, value) => {
  let Minutes = 20
  let exp = new Date()
  exp.setTime(exp.getTime() + (Minutes * 60 * 60 * 1000))
  document.cookie = `${name}=${escape(value)};expires=${exp.toGMTString()}`
}

const getCookie = (name) => {
  let arr
  let reg = new RegExp(`(^| )${name}=([^;]*)(;|$)`)
  arr = document.cookie.match(reg)
  if (arr) { return unescape(arr[2]) }
  return null
}

const delCookie = (name) => {
  let exp = new Date()
  exp.setTime(exp.getTime() - 1)
  let cval = getCookie(name)
  if (cval != null) { document.cookie = `${name}=${cval};expires=${exp.toGMTString()}` }
}

export default {
  setCookie,
  getCookie,
  delCookie,
}
