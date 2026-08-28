// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

    async validateRefresh() {

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

    }

  })

  )

})
