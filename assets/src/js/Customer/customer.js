document.addEventListener('alpine:init', () => {
  Alpine.data('customerComponent', () => ({
    records: [],
    page: 1,
    hasNextPage: false,
    loading: false,

    async init() {
      this.loadRecords()
    },

    async goHome() { await GoHome() },
    async goBack() { await GoBack() },
    async logOut() { await LogOut() },

    async newRecord() {
      try {
        const res = await SecureFetching("/customer", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = `/customer/new`
      } catch (error) {
        throw error
      }
    },

    async nextPage() { if (this.hasNextPage) { this.page += 1; await this.loadRecords() } },
    async previousPage() { if (this.page > 1) { this.page -= 1; await this.loadRecords() } },

    async openRecord(id) {
      try {
        const res = await SecureFetching("/customer", { method: "HEAD" })
        if (!res.ok) {
          throw new Error(res.status)
        }
        window.location.href = `/customer/record?id=${id}`
      } catch (error) {
        throw error
      }
    },

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

document.addEventListener('alpine:init', () => {
  Alpine.data('customerRecordComponent', (record) => ({
    id: record.id || null,
    name: record.name || '',
    fullName: record.fullName || '',
    address: record.address || '',
    technicianPhone: record.technicianPhone || '',
    buyerPhone: record.buyerPhone || '',
    customerEmail: record.customerEmail || '',

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/customer") },

    required() { },

    async updateRecord() { },
    async deleteRecord() { },

  }))
})
