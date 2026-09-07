// HANDLES CATEGORY VIEW
document.addEventListener('alpine:init', () => {
  Alpine.data('productCategoryComponent', () => ({

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

// HANDLES CATEGORY LIST
document.addEventListener('alpine:init', () => {
  Alpine.data('categoryListComponent', () => ({

    records: [],
    page: 1,
    hasNextPage: false,
    loading: false,

    async init() {
      await this.loadRecords()
    },

    async loadRecords() {
      this.loading = true
      try {
        const res = await SecureFetching(`/product/category/list?page=${this.page}`)
        if (!res.ok) throw new Error(`Error ${res.status}`)
        const data = await res.json()
        this.records = data.records // cambia si el contrato JSON usa otro nombre
        this.hasNextPage = data.hasNextPage // generar validacion en service
      } finally { this.loading = false }
    },

    async nextPage() { if (this.hasNextPage) { this.page += 1; await this.loadRecords() } },

    async previousPage() { if (this.page > 1) { this.page -= 1; await this.loadRecords() } },

    async openRecord(id) {
      try {
        const res = await SecureFetching("/product/category/record", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = `/product/category/record?id=${id}`
      } catch (error) {
        throw error
      }
    },

    async newRecord() {
      try {
        const res = await SecureFetching("/product/category/new", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/product/category/new"
      } catch (error) {
        throw error
      }
    },

  }))
})
