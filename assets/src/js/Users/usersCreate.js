document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersCreateComponent', () => ({

    name: "",
    email: "",
    password: "",
    repeatPassword: "",
    required: false,
    passMatch: false,

    checkFields() {
      if (!this.name || !this.email || !this.password || !this.repeatPassword) {
        return false
      } else {
        return true
      }
    },

    checkPassword() {
      if (this.password != this.repeatPassword) {
        return false
      } else {
        return true
      }
    },

    // CREATE NEW
    async createUser() {
      try {
        this.passMatch = false
        this.required = false
        const allowCreation = this.checkFields()
        if (allowCreation) {
          const passwordMatch = this.checkPassword()
          if (!passwordMatch) {
            this.passMatch = true
            return
          }
          console.log("Creating record in db...")
          const res = await SecureFetching(
            "/users/new",
            {
              method: "POST",
              body: JSON.stringify(
                {name: this.name, email: this.email, password: this.password}
              )
            }
          )
          if (!res.ok) {
            const errorText = await res.text();
            throw new Error(errorText)
          }
          const newUser = await response.json();
          const newId = newUser.id
          window.location.href = `/users/record?id=${newId}`
        } else {
          this.required = true
          return
        }
      } catch (error) {
        throw error
      }
    },

    // GO HOME
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

    // LOG OUT
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
