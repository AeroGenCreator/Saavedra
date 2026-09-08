// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

  async goHome() { await GoHome() },
  async logOut() { await LogOut() },

  async secureUsers() { await OnlyRedirect("/users") },
  async secureProduct() { await OnlyRedirect("/product/menu") },
  async secureCustomer() { await OnlyRedirect("/customer") },
  async secureQuote() {  await OnlyRedirect("/quote/menu")}

  }))
})
