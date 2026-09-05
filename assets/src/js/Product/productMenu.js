document.addEventListener('alpine:init', () => {
  Alpine.data('menuComponent', () => ({

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

    async secureMaterial() {
      try {
        const res = await SecureFetching("/product/material", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product/material"
      } catch (error) {
        throw error
      }
    },

    async secureProveedor() {
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

    async secureProduct() {
      try {
        const res = await SecureFetching("/product", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product"
      } catch (error) {
        throw error
      }
    },

  }))
})
