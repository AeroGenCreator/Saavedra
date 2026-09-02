document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersRecordComponent', () => ({

    id: '',
    name: '',
    email: '',
    password: '',
    repeatPassword: '',
    required: false,

    // LOAD GO STATIC DATA INTO THIS ALPINE COMPONENT
    initData(id, name, email) {
      this.id = id
      this.name = name;
      this.email = email;
    },

    // Check required fields.
    noEmptys() {
      if (this.name === '' || this.email === '') {
        return false
      } else {
        return true
      }
    },

    // Check passwordValidation
    checkPasword() {
      if (this.password != this.repeatPassword) {
        return false
      } else {
        return true
      }
    },

    // UPDATE USER
    async updateUser() {
      try {
        this.required = false
        const requiredFields = this.noEmptys()
        if (!requiredFields) {
          this.required = true
          return
        }
        var newPassword = ""
        if (this.password && this.repeatPassword) {
          const approvedPass = this.checkPasword()
          console.log("Si actualizar las password...")
          newPassword = this.password
        }
        const res = await SecureFetching(
          "/users/record",
          {
            method: "PUT", body: JSON.stringify(
              { id: this.id, name: this.name, email: this.email, password: this.password }
            ),
          }
        )
        if (!res.ok) {
          throw new Error(res.status)
        }
      }
      catch (error) {
        throw error
      }
      finally {
        location.reload();
      }
    },

    // DELETE USER
    async deleteUser() {
      return
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
