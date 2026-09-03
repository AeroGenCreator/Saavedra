document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersRecordComponent', () => ({

    id: '',
    name: '',
    email: '',
    password: '',
    repeatPassword: '',
    newPassword: '',
    passMatch: '',

    initData(id, name, email) {
      this.id = id
      this.name = name;
      this.email = email;
    },

    passwordsMatch() { return (this.password === this.repeatPassword) },

    async updateRecord() {
      try {

        this.newPassword = ''
        this.passMatch = ''

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

    async deleteRecord() {
      try {
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

    async goBack() {
      try {
        const res = await SecureFetching("/users", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/users"
      } catch (error) {
        throw error
      }
    },

    async goHome() {
      const res = await SecureFetching('/welcome', { method: 'HEAD' })
      if (res.ok) {
        window.location.href = '/welcome'
        return
      }
      await this.logOut()
      alert(res.status === 401 ? 'Sesión expirada' : `Error ${res.status}`)
    },

    async logOut() {
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
    },

  }))
})
