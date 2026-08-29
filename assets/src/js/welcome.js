// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

    async imgRefresh() {

      const res = await SecureFetching("/welcome")

      if (!res.ok) {

        if (res.status === 401) {

          alert("Session expires")
          window.location.href = "/"

        } else {

          alert(res.status)
          window.location.href = "/"

        }

      } else {

        window.location.href = "/welcome"

      }

    },

  async closeSession() {

    try {
      console.log("Attempting user logout...")
      const res = await fetch("/login", { method: "PATCH", credentials: "include" })
      if (res.ok) {

        window.location.href = "/login", { method: "GET" }

      }

    } catch (error) {

      console.log(error)
      return

    }

  }

  })

  )

})
