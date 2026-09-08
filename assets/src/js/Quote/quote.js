document.addEventListener('alpine:init', () => {
  Alpine.data('menuComponent', () => ({

    async goHome() { await GoHome() },
    async logOut() { await LogOut() },
    async goBack() { await GoBack("/welcome") },

    secureQuoteNew() { return },
    secureQuote() { return },

  }))
})
