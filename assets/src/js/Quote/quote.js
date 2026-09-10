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

    records: [], userArray: [], customerArray: [], productArray: [], user: '', customer: '',
    hrMen: 0, hrCost: 0, transportCost: 0, product: '', quantity: 0, loading: false, productTotal: 0, sendMulti: false,

    async init() {
      this.loading = true
      try {
        const data = await FetchDataFromResponse("/quote/many2one", { method: "GET" })
        this.userArray = data.userArray
        this.customerArray = data.customerArray
        this.productArray = data.productArray
      } finally { this.loading = false }
    },

    appendItem() {
      const exists = this.records.find(item => item.name.toLowerCase() === this.product.toLowerCase())
      if (exists) {
        alert(`El articulo ${this.product} ya existe en la lista, eliminar primero.`)
        return
      }
      var record = this.productArray.find(item => item.name.toLowerCase() === this.product.toLowerCase());
      record.quantity = this.quantity;
      record.subtotal = record.quantity * record.price;
      this.records.push(record);
      this.productTotal = this.records.reduce((sum, item) => sum + item.subtotal, 0)
      this.product = ''
      this.quantity = 0
    },

    popItem(name) {
      var newRecords = this.records.filter(item => item.name !== name)
      this.productTotal = newRecords.reduce((sum, item) => sum + item.subtotal, 0)
      this.records = newRecords
    },

    hrMenTotal() { return this.hrMen * this.hrCost},
    total() { return parseFloat(this.transportCost) + parseFloat(this.hrMenTotal()) + this.productTotal},

    printQuote() { return },
    createRecord() { return },
    send() { return },

    appendRequired() { return (this.product === '' || this.quantity <= 0) },
    toggleMulti() { return (!this.sendMulti) },
    unlockProduct() {
      return (this.user === '' || this.customer === '' || this.hrMen <= 0 || this.hrCost <= 0 || this.transportCost <= 0)
    },
    unlockActions() {
      return (this.unlockProduct() || this.records.length === 0)
    },

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/quote/menu") },

  }))
})
