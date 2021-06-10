import config from '@/config'

type LoginByGuestParams = {
  name: string
  image: Blob
}

export const loginByGuest = async ({ name, image }: LoginByGuestParams) => {
  const formData = new FormData()

  formData.append('name', name)
  formData.append('image', image)

  try {
    const res = await fetch(`${config.apiBaseUrl}/guest-login`, {
      // @see: https://developer.mozilla.org/ja/docs/Web/API/Fetch_API/Using_Fetch#sending_a_request_with_credentials_included
      // set-cookieを機能させるために必要なパラメータ指定
      credentials: 'include',
      method: 'POST',
      body: formData,
    })

    if (res.status >= 500) {
      alert('サーバーエラー')
      return
    } else if (res.status >= 400) {
      const resData = await res.json()
      alert('リクエストエラー:' + resData.message)
      console.log('OK!')
      return
    }
  } catch (err) {
    console.log(err)
  }
}
