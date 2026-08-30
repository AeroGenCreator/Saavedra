## Header

```html
<!DOCTYPE html>
<html lang="es">

<head>
    <meta charset="UTF-8">
    <!-- Load Frontend Dependencies & configuration-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/web/lucide-font/lucide.css">
    <link rel="stylesheet" href="/assets/css/bulma/css/bulma.min.css">
    <script defer src="/assets/src/js/utils/utils.js"></script>
    <script defer src=""><!-- Specific JS script for the current service --></script>
    <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.16.3/dist/cdn.min.js"></script>

	<title>Current Service</title>
</head>

<!-- NAVBAR & FOOTER -->
<body class="has-navbar-fixed-top">

    <main class="hero">
        <div class="container" x-data="welcomeComponent">
            <nav class="navbar is-fixed-top has-background-black-ter">
                <div class="navbar-brand">
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

</body>
</html>

```

Basic configuration for a new menuitem:

```html
<div class="column is-3">
<a href="/users">
    <div class="notification is-danger">
        <article class="has-text-left">
            <i class="icon-users"></i>
            <div class="content is-small">
                <h3>Servicio Usuarios</h3>
                <p class="has-text-small">Crear, editar o eliminar usuarios.</p>
            </div>
        </article>
    </div>
</a>
</div>
```
