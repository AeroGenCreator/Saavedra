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
    constant: 0,
    product: '',
    quantity: 0,
    loading: false,
    hrsHombreTotal: 0,
    productosTotal: 0,
    total: 0,

    init() { this.fetchData() },

    fetchData() { return },
    popItem() { return },
    appendItem() { return },
    openRecord(name) { return },

    saveRequired() { return },
    appendRequired() { return },

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/quote/menu") },

  }))
})
