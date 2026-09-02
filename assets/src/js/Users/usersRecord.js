document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersRecordComponent', () => ({

    id: '',
    name: '',
    email: '',
    password: '',
    repeatPassword: '',
    newPassword: '',
    required: '',
    passMatch: '',

    // LOAD GO STATIC DATA INTO THIS ALPINE COMPONENT
    initData(id, name, email) {
      this.id = id
      this.name = name;
      this.email = email;
    },

    // CHECK REQUIRED FIELDS
    hasEmptyRequiredFields() { return (this.name === '' || this.email === '') },

    // Check passwordValidation
    passwordsMatch() { return (this.password === this.repeatPassword) },

    // UPDATE USER
    async updateUser() {
      try {

        this.newPassword = ''
        this.required = ''
        this.passMatch = ''

        if (this.hasEmptyRequiredFields()) {
          this.required = "Faltan campos."
          return
        }

        if (this.password || this.repeatPassword) {
            if (!this.passwordsMatch()) {
              this.passMatch = "Contraseñas no coinciden."
              return
            }

            this.newPassword = this.password
          }

        const res = await SecureFetching(
          "/users/record",
          {
            method: "PUT", body: JSON.stringify(
              { id: this.id, name: this.name, email: this.email, password: this.newPassword }
            ),
          }
        )
        if (!res.ok) {
          throw new Error(res.status)
        }
        location.reload()
      }
      catch (error) {
        throw error
      }
    },

    // DELETE USER
    async deleteUser() {
      try {

        console.log(`Attempting delete record: ${this.id}`)
        const res = await SecureFetching(
          "/users/record",
          {
            method: "DELETE", body: JSON.stringify(
              { id: this.id, name: this.name, email: this.email, password: this.password }
            ),
          }
        )

        if (!res.ok) {
          switch (res.status) {
            case 403:
              alert("Server has forbidden this record to be deleted.")
              location.reload()
            default:
              throw new Error(res.status)
          }
        }

        window.location.href = "/users"

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
