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
    async goNew() { await GoNew("/customer/new") },
    async openRecord(id) { await OpenRecord(id, `/customer/record?id=${id}`)},

    async nextPage() { if (this.hasNextPage) { this.page += 1; await this.loadRecords() } },
    async previousPage() { if (this.page > 1) { this.page -= 1; await this.loadRecords() } },

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

    required() {
      return (
        [
          this.name,
          this.fullName,
          this.address,
          this.technicianPhone,
          this.buyerPhone,
          this.customerEmail
        ].includes("")
      )
    },

    async updateRecord() {
      const values = JSON.stringify(
        {
          id: this.id,
          name: this.name,
          fullName: this.fullName,
          address: this.address,
          technicianPhone: this.technicianPhone,
          buyerPhone: this.buyerPhone,
          customerEmail: this.customerEmail,
        }
      )
      await UpdateRecord("/customer/record", "/customer", { method: "PUT", body: values })
    },

    async deleteRecord() {
      await DeleteRecord(`/customer/record?id=${this.id}`, "/customer", { method: "DELETE" })
    },
  }))
})

document.addEventListener('alpine:init', () => {
  Alpine.data('customerNewComponent', () => ({

    name: '',
    fullName: '',
    address: '',
    technicianPhone: '',
    buyerPhone: '',
    customerEmail: '',

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/customer") },

    required() {
      return (
        [
          this.name,
          this.fullName,
          this.address,
          this.technicianPhone,
          this.buyerPhone,
          this.customerEmail
        ].includes("")
      )
    },

    async createRecord() {
      const values = JSON.stringify({
        name: this.name,
        fullName: this.fullName,
        address: this.address,
        technicianPhone: this.technicianPhone,
        buyerPhone: this.buyerPhone,
        customerEmail: this.customerEmail
      })
      await CreateRecord("/customer/new", "/customer", { method: "POST", body: values })
    },

  }))
})
