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

  console.log("Attempting secure fetch...")

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
    console.log(route, options)
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
