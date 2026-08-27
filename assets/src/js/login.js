// Alpine Init

document.addEventListener('alpine:init', () => {

  Alpine.data("LoginData", () => ({

    // DOM bound variables
    "email": "",
    "password": "",
    "statusMessage": false,

    // Submit function
    async SubmitLogin() {

      console.log(`Attempting login request for '${this.email}'`)

      try {

        // Login Request => POST {Credentials}
        const response = await fetch(
          "/login",
          {
            method: "POST",
            body: JSON.stringify({ "email": this.email, "password": this.password }),
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
          })

        // Handler of errors HTTP
        if (!response.ok) {

          // Unauthorized
          if (response.status === 401) {

            console.log("Invalid Credentials");

            this.statusMessage = true;

            return;
          }

          // Real errros
          throw new Error(`Server error: ${response.status}`);
        }

        // Decoding Data From Response
        const objectJSON = await response.json()

        console.log(objectJSON)

        this.statusMessage = false;

        // Redirection
        console.log("Correct Login Pattern")
        window.location.href = "/welcome";

      } catch (error) {

        console.log("Error", error)
      }

    }

  }))
})
