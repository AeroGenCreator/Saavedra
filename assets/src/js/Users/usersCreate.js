document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersCreateComponent', () => ({

    name: '',
    email: '',
    password: '',
    repeatPassword: '',
    required: '',
    passMatch: '',

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

        this.required = ''
        this.passMatch = ''

        const allowCreation = this.checkFields()
        if (!allowCreation) {
          this.required = "Todos los campos son obligatorios."
          return
        }

        const passwordMatch = this.checkPassword()
        if (!passwordMatch) {
          this.passMatch = "La validación de contraseña no coincide."
          return
        }

        console.log("Attempting request at /users/new...")
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

        const newUser = await res.json();
        const newId = newUser.id

        window.location.href = `/users/record?id=${newId}`

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
