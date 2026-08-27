// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

    // DOM BOUND VARIABLES && CONTENTS
    "headerContent": '',
    "footerContent": '',
    "initStatus": true,

    async init() {

      try {

        const headerRes = await fetch("/assets/src/html/header.html");
        const footerRes = await fetch("/assets/src/html/footer.html");

        this.headerContent = await headerRes.text();
        this.footerContent = await footerRes.text();

      } catch (error) {

        console.error("Layout loading erro:", error);

      } finally {

        this.initStatus = false;

      }

    }

  })
  )
})
