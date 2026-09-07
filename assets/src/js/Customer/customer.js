document.addEventListener('alpine:init', () => {
  Alpine.data('customerComponent', () => ({
    records: [],
    page: 1,
    hasNextPage: false,
    loading: false,

    async init() {
      this.loadRecords()
    },

    async goBack() {
      try {
        const res = await SecureFetching("/welcome", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/welcome"
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

    newRecord() { },

    previousPage() { },
    nextPage() { },
    openRecord(id) { },

    async loadRecords() {
      this.loading = true
      try {
        const res = await SecureFetching(`/customer/slice?page=${this.page}`)
        if (!res.ok) throw new Error(`Error ${res.status}`)
        const data = await res.json()
        this.records = data.records
        this.hasNextPage = data.hasNextPage
      } finally { this.loading = false }
    },

  }))
})
