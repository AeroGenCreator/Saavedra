document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersComponent', () => ({

    // Relates an event with a function
    async goHome() {

      console.log("Attempting go Home...")

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

        console.log("Attempting logout...")

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

    }

  }))
})
