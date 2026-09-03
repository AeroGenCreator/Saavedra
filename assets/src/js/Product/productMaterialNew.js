// HANDLES MATERIAL VIEW
document.addEventListener('alpine:init', () => {
  Alpine.data('productMateriaNewComponent', () => ({
    material: '',

    async create() {
      try {
        const res = await SecureFetching("/product/material/new", { method: "POST", body: JSON.stringify({name: this.material}) })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product/material"
      } catch (error) {
        throw error
      }
    },

    async goBack() {
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
