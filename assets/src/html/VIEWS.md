# Views

**This section discusses `HTML` and `CSS Styles` only.**

Every framework is compound of reusable components that simplify software development however `Saavedra` follows DSEE principles where redundant code is not a problem.

The following examples provide code snipets that are a must for specific purposes.

## Header

**None of the followig libraries are an obligation however `Saavedra Vanilla` is based on them**.

Since `Saavedra Vanilla` modularity has been taken to its limits, every single aspect can be changed but at the same time older 
services will require same changes to fit styles properly.

`Saavedra Vanilla` relays on the listed libraries below. 

- Alpine.js 3.16.3
- Bulma CSS
- Lucide

Copy and paste the following header in every `HTML` view before you start.

```html
<!DOCTYPE html>
<html lang="es">

<head>
    <meta charset="UTF-8">
    <!-- Load Frontend Dependencies & configuration-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/web/lucide-font/lucide.css">
    <link rel="stylesheet" href="/assets/css/savedraaCSS.css">
    <link rel="stylesheet" href="/assets/css/bulma/css/bulma.min.css">
    <script defer src="/assets/src/js/utils/utils.js"></script>

    <!-- Specific JS script for the current service -->
    <script defer src=""></script>

    <!-- defer 'Alpine' should load after your own JS to map every custom component. -->
    <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.16.3/dist/cdn.min.js"></script>

	<title>Current Service</title>
</head>

</html>
```

## Naavigation Bar

This section along with `Alpine.js` is the most important.

The navigation bar is not only HTML and Styles but an exposed line to understand `Session State` and `Alpine reusable snipets`.

The `HTML` structure is simply.

- Brang Image
- Logout Button

The first one teaches `session state` process while the second one `complements the first one` but most important, it `prevents you to program "logout" function again`.

### Alpine

When adding `x-data` to a `HTML` tags, developers can start interactions with the `DOM`, specificly with the section where the `x-data` was declared.

`Saavedra Vanilla` navbar was declared as it follows `class="container" x-data=""`.

The component name can change for every view but the code snipets stay the same. For example `class="container" x-data="tutorialComponent"` or `class="container" x-data="firstViewComponent"`.

Then you can create a `JavaScript` in `/assets/src/js/` called `tutorial.js` and initialize Alpine like this.

```js
document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('tutorialComponent', () => ({

  }))
})
```

Do not forget adding the new JavaScript file as an import before `Alpine` import.

```html
<!-- Specific JS script for the current service -->
<script defer src="/assets/src/js/tutorial.js"></script>
```

### Brand Image

Brand image is basically a `home` button. It takes user to `/welcome` page which is basically `home` page however important patterns must be take in count.

Alpine prevents default button behavior and instead we trigger a specific event before redirecting to /welcome page.

This is a security event, the code is the same, we only need to relate the action with a method in the `Alpine` component.

```html
<!-- The image tries redirection while alpine prevents and then calls for a JS function. -->
<a class="navbar-item" href="/welcome" @click.prevent="imgRefresh">
<!-- Changing the name is valid but not a need. -->
<a class="navbar-item" href="/welcome" @click.prevent="goHome">
```

So now your JS should look like this:

```js
document.addEventListener('alpine:init', () => {
  // Register this alpine component to a section in your HTML.
  Alpine.data('tutorialComponent', () => ({

    // Relates an event with a function
    async goHome() {
      
    }

  }))
})
```

Imagine we have follow every step, adding `goHome()` requires a secure fetching where state session is validated for every user request. Do not worry, the code for this already exists in `/assets/src/js/utils/utils.js` and since you have already imported it, you can only copy and paste the required logic for `tutorial.js` to ensure every user session state for your `home` action.

```js
async goHome() {
  console.log("Attempting refresh...", document.cookie)
  // SecureFetching is an obligation whenever you try a fetch(), this ensures the server to accept only user requests.
  // Otherwise, every request will be processed by the server.
  const res = await SecureFetching("/welcome")
  if (!res.ok) {
    if (res.status === 401) {
      this.closeSession()
      alert("Session expires")
    } else {
      this.closeSession()
      alert(res.status)
    }
  } else {
    window.location.href = "/welcome"
  }
},
```

The most important is `SecureFetching(route, requestContent = {}, customHeaders = {'X-Requested-With': 'jsFrontendComponent'})`. It ensures requets but also refreshes user token.

For every redirection or server request you shlould use it, unless you want to program your own session state handler.

### Logout

Same situation exists when it comes to `logout`. A button, a redirection, alpine prevents, some session state logic is triggered behind.

```html
<a href="/login" @click.prevent="logOut">
```

So your JS should include this event.

```js
async closeSession() {
  try {
    console.log("Attempting user logout...")
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
```

Notice that `SecureFetching` was not used since `/login` requires no protection.

## Navbar Template

Now that you understand the whole content of the navigation bar. You can safely copy and paste the following code snipet into every new `HTML` view.

```html
<!-- NAVBAR-->
<body class="has-navbar-fixed-top">

    <main class="hero">
        <div class="container" x-data="welcomeComponent">
            <nav class="navbar is-fixed-top has-background-black-ter">
                <div class="navbar-brand">
                    // Don't forget registering the event in your JS in an Alpine Component.
                    <a class="navbar-item" href="/welcome" @click.prevent="imgRefresh">
                        <figure class="image mx-6">
                            <img src="/assets/img/saavedraVanilla.svg"/>
                        </figure>
                    </a>
                </div>
                <div class="navbar-end">
                    <div class="navbar-item">
                        <div class="field is-grouped">
                            <p class="control">
                                // Don't forget registering the event in your JS in an Alpine Component.
                                <a href="/login" @click.prevent="closeSession">
                                    <button class="button is-primary px-3 is-small">
                                        Cerrar Sesión
                                        <span class="px-2">
                                            <i class="icon-log-in"></i>
                                        </span>
                                    </button>
                                </a>
                            </p>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </main>

</body>
```

## Footer

The footer handles no important logic. You cant copy and paste it, even edited if needed. Just remember it should be placed inside a `body` tag.

```html
<!-- FOOTER -->
<footer class="footer has-background-black-ter p-0">
    <div class="content has-text-left">

        <section class="section">
            <div class="container">
                <div class="columns">

                    <div class="column is-6">
                        <p class="has-text-white is-size-7">
                            <span class="is-bold">Saavedra Vanilla</span> esta construido sobre una arquitectura
                            desacomplada basada en DSEE.
                            Cualquier servicio extra puede ser desarrollado según la necesidad del negocio.
                            El cobro depende de la cantidad de
                            capas necesitadas por servicio lo cual reduce el costo drasticamente.
                            Para cotizar un servicio extra
                            en tu app
                            <a href="mailto:gallo85floyd@gmail.com">
                                <strong class="has-text-primary">contactanos</strong>.
                            </a>
                        </p>
                    </div>

                    <div class="column is-6">
                        <p class="is-size-7 has-text-white">
                            Numero de Contacto / WhatsApp:
                            <span class="has-text-primary">
                                <i class="icon icon-phone"></i>
                                + 52 246 207 29 85
                            </span>
                        </p>
                    </div>

                </div>
            </div>
        </section>

    </div>
</footer>
```

## Menu Cards

Basic configuration for a new menuitem:

```html
<div class="column is-3">
      <a href="/users">
          <div class="notification is-danger card-fixed-height is-clickable-card">
              <article class="has-text-left">
                  <i class="icon-users"></i>
                  <div class="content is-small">
                      <h3>Servicio Usuarios</h3>
                      <p class="has-text-small truncate-text">
                          Crear, editar o eliminar usuarios.
                      </p>
                  </div>
              </article>
          </div>
      </a>
</div>
```
