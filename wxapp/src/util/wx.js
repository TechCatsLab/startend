import wepy from 'wepy'

export function checkSetting({ success, fail }) {
  wepy.showModal({
    title: '微信授权提示',
    content: '',
    showCancel: true,
    success: () => {
      wepy.getSetting({
        success: res => {
          console.log(res.authSetting)
        },
        fail: () => {
        }
      })
    },
    fail: () => {
    }
  })
}

export function wxCheckSession() {
  return new Promise((resolve, reject) => {
    wepy.checkSession({
      success: res => resolve(res),
      fail: err => reject(err)
    })
  })
}

export function wxLogin() {
  return new Promise((resolve, reject) => {
    wepy.login({
      success: res => resolve(res),
      fail: err => reject(err)
    })
  })
}

export async function login() {
  let loginRequired = false

  try {
    await wxCheckSession()
  } catch (err) {
    loginRequired = true
  }

  if (loginRequired) {
    const resp = await wxLogin()
    return resp
  }

  return true
}
