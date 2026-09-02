document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('usersComponent', () => ({

    // Users model variables
    records: [],
    page: 1,
    totalPages: 1,
    totalRecords: 0,
    loading: false,

    // INITIALIZE -> FETCH PAGE 1
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
        this.totalRecords = data.count_pages

      } catch (error) {
        debugger;
        throw new Error(error)

      }

    },

    // NEW USER BUTTON
    async newUser() {

      try {

        // METHOD: HEAD -> SECURE FETCHING RETURNS STATUS -> REDIRECTION IF STATUS OK. OPTIMIZED PROCESS
        const res = await SecureFetching("/users/new", { method: "HEAD" })

        if (!res.ok) {

          throw new Error(`Bad response from server...${res.status}`)

        }

        window.location.href = "/users/new"

      } catch (error) {

        throw error

      }

    },

    // OPEN AN EXISTING RECORD
    async openRecord(id) {

      try {

        const res = await SecureFetching("/users/record", { method: "HEAD" })

        if (!res.ok) {

          throw new Error(res.status)

        }

        window.location.href = `/users/record?id=${id}`

      } catch (error) {

        throw (error)

      }

    },

    // FETCH 1 PLUS PAGE
    async loadPlusTable() {

      try {

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

        data = await res.json()

        this.records = data.rows
        this.page = data.current_page
        this.totalPages = data.count_pages
        this.totalRecords = data.count_pages

      } catch (error) {

        throw new Error(error)

      } finally {

        this.loading = false

      }

    },

    // FETCH 1 LESS PAGE
    async loadMinusTable() {

      try {

        this.loading = true

        this.page = this.page - 1

        if (this.page <= 0) {

          console.log("No less pages...")
          this.page = 1
          return
        }

        const res = await SecureFetching(`/users/records?page=${this.page}`)

        if (!res.ok) {

          throw new Error(res.status)

        }

        data = await res.json()

        this.records = data.rows
        this.page = data.current_page
        this.totalPages = data.count_pages
        this.totalRecords = data.count_pages

      } catch (error) {

        throw new Error(error)

      } finally {

        this.loading = false

      }

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
