document.addEventListener('alpine:init', () => {
  Alpine.data('productComponent', () => ({

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
  Alpine.data('productListComponent', () => ({

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
        const res = await SecureFetching(`/product/slice?page=${this.page}`)
        if (!res.ok) throw new Error(`Error ${res.status}`)
        const data = await res.json()
        this.records = data.records
        this.hasNextPage = data.hasNextPage
      } finally { this.loading = false }
    },

    async newRecord() {
      try {
        const res = await SecureFetching("/product/new", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product/new"
      }
      catch (error) {
        throw error
      }
    },

    async nextPage() { if (this.hasNextPage) { this.page += 1; await this.loadRecords() } },

    async previousPage() { if (this.page > 1) { this.page -= 1; await this.loadRecords() } },

    async openRecord(id) {
      try {
        const res = await SecureFetching("/proveedor/record", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = `/proveedor/record?id=${id}`
      } catch (error) {
        throw error
      }
    },

  }))
})
