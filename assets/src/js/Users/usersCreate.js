document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersCreateComponent', () => ({

    name: '',
    email: '',
    password: '',
    repeatPassword: '',
    newPassword: '',
    passMatch: '',

    passwordsMatch() { return (this.password === this.repeatPassword) },

    async create() {
      try {

        this.passMatch = ''
        if (this.password && this.repeatPassword) {
            if (!this.passwordsMatch()) {
              this.passMatch = "Contraseñas no coinciden."
              return
            }
            this.newPassword = this.password
          }

        console.log("Attempting request at /users/new...")
        const res = await SecureFetching(
          "/users/new",
          {
            method: "POST",
            body: JSON.stringify(
              {name: this.name, email: this.email, password: this.newPassword}
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
