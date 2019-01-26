import wepy from 'wepy'

const apiToken = 'api-token'

export function saveToken(token) {
  try {
    wepy.setStorageSync(apiToken, token)
    return true
  } catch (err) {
    return false
  }
}

export function getToken() {
  try {
    const token = wepy.getStorageSync(apiToken)
    return token
  } catch (err) {
    return false
  }
}

export function removeToken() {
  try {
    wepy.removeStorageSync(apiToken)
    return true
  } catch (err) {
    return false
  }
}
