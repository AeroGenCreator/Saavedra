document.addEventListener('alpine:init', () => {
  Alpine.data('proveedorComponent', () => ({

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

document.addEventListener('alpine:init', () => {
  Alpine.data('proveedorListComponent', () => ({

    records: [],
    page: 1,
    hasNextPage: false,
    loading: false,

    async init() {
      this.loadRecords()
    },

    async goBack() {
      try {
        const res = await SecureFetching("/product/menu", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product/menu"
      } catch (error) {
        throw error
      }
    },

    async loadRecords() {
      this.loading = true
      try {
        const res = await SecureFetching(`/proveedor/slice?page=${this.page}`)
        if (!res.ok) throw new Error(`Error ${res.status}`)
        const data = await res.json()
        this.records = data.records
        this.hasNextPage = data.hasNextPage
      } finally { this.loading = false }
    },

    async newRecord() { },

    async previousPage() { },

    async nextPage() { },

    async openRecord(id) { },

  }))
})
