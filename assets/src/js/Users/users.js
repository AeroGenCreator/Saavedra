document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersComponent', () => ({

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

  }))
})


document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersSliceComponent', () => ({

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
        const res = await SecureFetching(`/users/records?page=${this.page}`)
        if (!res.ok) throw new Error(`Error ${res.status}`)
        const data = await res.json()
        this.records = data.records
        this.hasNextPage = data.hasNextPage
      } finally { this.loading = false }
    },

    async nextPage() { if (this.hasNextPage) { this.page += 1; await this.loadRecords() } },

    async previousPage() { if (this.page > 1) { this.page -= 1; await this.loadRecords() } },

    async openRecord(id) {
      try {
        const res = await SecureFetching("/users/record", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = `/users/record?id=${id}`
      } catch (error) {
        throw error
      }
    },

    async newRecord() {
      try {
        const res = await SecureFetching("/users/new", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = "/users/new"
      } catch (error) {
        throw error
      }
    },

  }))
})
