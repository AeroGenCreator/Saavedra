// STORES MULTI REQUEST
let refreshSubscribers = [];

// FLAG - ALERTS OF REFRESH IN COURSE
let isRefreshing = false;

// ONCE TOKEN IS REFRESHED, REQUEST CAN BE SEND
function onTokenRefreshed() {

  refreshSubscribers.forEach((callback) => callback());

  // REFRESH STORING ARRAY
  refreshSubscribers = [];

}

async function SecureFetching(route, requestContent = {}, customHeaders = {'X-Requested-With': 'jsFrontendComponent'}) {

  console.log(`Attempting secure fetch for ${route}...`)

  // CREDENTIALS, HEADERS && OPTIONS
  const options = {
    ...requestContent,
    credentials: 'include',
    headers: {
      'X-Requested-With': 'jsFrontendComponent',
      ...(requestContent.headers || {}),
      ...customHeaders
    }
  };

  try {

    let response = await fetch(route, options);

    // IF ORIGINAL FETCH REQUIRES REFRESHING|
    if (response.status === 401) {

      // FIRST: IS THERE OTHER REFRESHING IN COURSE?
      if (isRefreshing) {

        // RETURN PROMISE
        return new Promise((resolve) => {

          // REQUETS IS SAVED AS A FUNCTION OF A NEW FETCHING
          refreshSubscribers.push(async () => {

            resolve(await fetch(route, options));

          });

        });

      }

      // IF NO REFRESH THEN REFRESHING PETITION CAN TAKE PLACE
      isRefreshing = true;
      console.log("First attempt of refreshing...")
      const refreshResponse = await fetch("/refresh", {
        method: "POST",
        credentials: 'include'
      });

      // ONCE REFRESH IS DONE, FLAG CAN RETURN TO FALSE
      isRefreshing = false;

      // POSITIVE REFRESH? 1. ALL REQUEST ARE EXECUTED 2. NEW FETCH WITH NEW TOKEN
      if (refreshResponse.ok) {

        onTokenRefreshed();

        return await fetch(route, options);

      } else {

        // DB TOKEN AND COOKIE DON'T MATCH, SESSION EXPIRED.
        return refreshResponse;

      }

    }

    // RETURNS ORIGINAL RESPONSE AS LONG AS IT IS NOT 401
    return response;

  } catch (error) {

    console.error("Reusable fetching component error:", error);

    throw error;

  }

}

async function GoHome() {
  const res = await this.SecureFetching('/welcome', { method: 'HEAD' })
  if (res.ok) {
    window.location.href = '/welcome'
    return
  }
  await this.LogOut()
  alert(res.status === 401 ? 'Sesión expirada' : `Error ${res.status}`)
}

async function LogOut() {
  try {
    const res = await fetch('/login', {
      method: 'PATCH',
      credentials: 'include',
      headers: { 'X-Requested-With': 'jsFrontendComponent' },
    })
    if (res.ok) window.location.href = '/login'
  } catch (error) {
    console.error('No fue posible cerrar sesión:', error)
  }
}

async function GoBack(redirect = "/welcome") {
  try {
    const res = await SecureFetching("/welcome", { method: "HEAD" })
    if (!res.ok) {
      throw new Error(res.status)
    }
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function GoNew(redirect) {
  try {
    const res = await SecureFetching("/welcome", { method: "HEAD" })
    if (!res.ok) {
      throw new Error(res.status)
    }
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function OpenRecord(id, redirect) {
  try {
    const res = await SecureFetching("/welcome", { method: "HEAD" })
    if (!res.ok) {
      throw new Error(res.status)
    }
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function CreateRecord(path, redirect, options = {}) {
  try {
    const res = await SecureFetching(path, options)
    if (!res.ok) {
      throw new Error(res.status)
    }
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function UpdateRecord(path, redirect, options = {}) {
  try {
    const res = await SecureFetching(path, options)
    if (!res.ok) throw new Error(await res.text())
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function DeleteRecord(path, redirect, options = {}) {
  try {
    const res = await SecureFetching(path, options)
    if (!res.ok) throw new Error(await res.text())
    window.location.href = redirect
  } catch (error) {
    throw error
  }
}

async function OnlyRedirect(path) {
  try {
    const res = await SecureFetching("/welcome", { method: "HEAD" })
    if (!res.ok) {
      throw new Error(res.status)
    }
    window.location.href = path
  } catch (error) {
    throw error
  }
}

async function FetchDataFromResponse(path, options = {}) {
  try {
    const res = await SecureFetching(path, options)
    if (!res.ok) throw new Error(`Error ${res.status}`)
    const data = await res.json()
    return data
  } catch (error) { throw error }
}
