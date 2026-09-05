document.addEventListener('alpine:init', () => {
  Alpine.data('proveedorNewComponent', () => ({

  //FIELDS
    name: '',
    phone: '',

  // METHODS

  async create() {
    try {
      const res = await SecureFetching("/proveedor/new", {
        method: "POST", body: JSON.stringify(
          { name: this.name, phone: this.phone }
        )
      })
      if (!res.ok) {
        throw new Error(res.status)
      }
      window.location.href = "/proveedor"
    } catch (error) {
      throw error
    }
  },

  async goBack() {
    try {
      const res = await SecureFetching("/proveedor", { method: "HEAD" })
      if (!res.ok) {
        throw new Error(res.status)
      }
      window.location.href = "/proveedor"
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
