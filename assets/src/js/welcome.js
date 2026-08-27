// ALPINE INIT

document.addEventListener('alpine:init', () => {

  Alpine.data("welcomeComponent", () => ({

    // DOM BOUND VARIABLES && CONTENTS
    "headerContent": '',
    "footerContent": '',
    "initStatus": true,

    async init() {
      try {

        const headerRes = await fetch("/assets/layout/header.html");
        const footerRes = await fetch("/assets/layout/footer.html");

        this.headerContent = headerRes.text();
        this.footerContent = footerRes.text();

      } catch (error) {
        console.error("Layout loading erro:", error);
      } finally {
        this.initStatus = false;
      }
    }

  })
  )
})
