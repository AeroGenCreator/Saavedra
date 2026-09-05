document.addEventListener('alpine:init', () => {
  Alpine.data('proveedorRecordComponent', () => ({

    id: '',
    name: '',
    phone: '',

    init(id, name, phone) {
      this.id = id
      this.name = name
      this.phone = phone
    },

    async updateRecord() {
      try {
        const res = await SecureFetching('/proveedor/record', {
          method: 'PUT', body: JSON.stringify({ id: this.id, name: this.name, phone: this.phone }),
        })
        if (!res.ok) throw new Error(await res.text())
        window.location.reload()
      } catch (error) {
        throw error
      }
    },

    async deleteRecord() {
      try {
        const res = await SecureFetching(`/proveedor/record?id=${this.id}`, {
          method: 'DELETE', body: JSON.stringify({ id: this.id })
        })
        if (!res.ok) throw new Error(await res.text())
        window.location.href = '/proveedor'
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
