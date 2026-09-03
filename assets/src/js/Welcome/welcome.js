// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

    async goHome() {

      console.log("Attempting refresh...", document.cookie)

      const res = await SecureFetching("/welcome")

      if (!res.ok) {

        if (res.status === 401) {
          this.logOut()
          alert("Session expires")


        } else {
          this.logOut()
          alert(res.status)

        }

      } else {

        window.location.href = "/welcome"

      }

    },

  async logOut() {

    try {
      console.log("Attempting user logout...")
      const res = await fetch(
        "/login", {
        method: "PATCH",
        credentials: "include",
        headers: { "X-Requested-With": "jsFrontendComponent" }
      })
      if (res.ok) {

        window.location.href = "/login", { method: "GET" }

      }

    } catch (error) {

      console.log(error)
      return

    }

  },

  async secureUsers() {
    try {
      const res = await SecureFetching("/users")
      if (!res.ok) {
      window.location.href = "/welcome"
      }
      window.location.href = "/users"
    } catch (error) {
      console.log(error)
    }
  },

  // REDIRECT: PRODUCT SERVICE
  async secureProduct() {
    try {
      const res = await SecureFetching("/product/menu", { method: "HEAD" })
      if (!res.ok) {
        throw new Error(res.status)
      }
      window.location.href = "/product/menu"
    } catch (error) {
      throw error
    }
  }

  })

  )

})
