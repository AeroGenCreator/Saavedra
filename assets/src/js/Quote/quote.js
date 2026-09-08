document.addEventListener('alpine:init', () => {
  Alpine.data('menuComponent', () => ({

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/welcome") },

    async secureQuoteNew() { await OnlyRedirect("/quote/new") },
    secureQuote() { return },

  }))
})

document.addEventListener('alpine:init', () => {
  Alpine.data('quoteNewComponent', () => ({

    records: [],
    userArray: [],
    customerArray: [],
    productArray: [],
    user: '',
    customer: '',
    hrsHombre: 0,
    hrCost: 0,
    transportCost: 0,
    product: '',
    quantity: 0,
    loading: false,
    hrsHombreTotal: 0,
    productosTotal: 0,
    total: 0,
    sendMulti: false,

    async init() {
      this.loading = true
      try {
        const data = await FetchDataFromResponse("/quote/many2one", { method: "GET" })
        this.userArray = data.userArray
        this.customerArray = data.customerArray
        this.productArray = data.productArray
      } finally { this.loading = false }
    },

    toggleMulti() { return (!this.sendMulti) },
    popItem() { return },
    appendItem() {
      this.product = '',
      this.quantity = 0
    },
    openRecord(name) { return },

    printQuote() { return },
    createRecord() { return },

    saveRequired() { return },
    appendRequired() { return (this.product === '' || this.quantity <= 0) },

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/quote/menu") },

  }))
})
