document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersComponent', () => ({

    // Users model variables
    records: [],
    page: 1,
    totalPages: 1,
    totalRecords: 1,
    loading: false,

    // When Component initialize
    async init() {

      try {

        const res = await SecureFetching(`/users/records?page=${this.page}`)

        if (!res.ok) {

          throw new Error(`HTTP error! Status: ${res.status}`);

        }

        var data = await res.json()

        this.records = data.rows
        this.page = data.current_page
        this.totalPages = data.count_pages
        this.totalRecords = data.count_records

      } catch (error) {
        debugger;
        throw new Error(error)

      }

    },

    // RENDER TABLE
    async loadPlusTable() {

      this.loading = true

      this.page = this.page + 1

      if (this.page > this.totalPages) {
        console.log("No more pages...")
        this.page = this.page - 1
        return
      }

      const res = await SecureFetching(`/users/records?page=${this.page}`)

      if (!res.ok) {

        throw new Error(res.status)

      }

      data = res.json()

      this.records = data.rows
      this.page = data.current_page
      this.totalPages = data.count_pages
      this.totalRecords = data.count_records

      this.loading = false

    },

    // GO HOME
    async goHome() {

      console.log("Attempting go Home...")

      const res = await SecureFetching("/welcome")

      if (!res.ok) {

        if (res.status === 401) {

          this.logOut()

          alert("Session expires")

        } else {

          this.logOut()

          alert(res.status)

        }

      } else {

        window.location.href = "/welcome"

      }

    },

    // LOG OUT
    async logOut() {

      try {

        console.log("Attempting logout...")

        const res = await fetch(
          "/login", {
          method: "PATCH",
          credentials: "include",
          headers: { "X-Requested-With": "jsFrontendComponent" }
        })

        if (res.ok) {

          window.location.href = "/login", { method: "GET" }

        }

      } catch (error) {

        console.log(error)

        return

      }

    }

  }))
})
