# Saavedra Vanilla

## Quick Start

To configure Saavedraa you must provide the following environment variables.

```env
# Special token to validate request from the client.
JWT_KEY=
# Specific name for your Sqlite3 Database
DATABASE_FILE_NAME=
# Administrator credentials
ADMIN_NAME=
ADMIN_EMAIL=
ADMIN_PASSWORD=
```

## Dependencies

Either to develop or run Saavedra it is important to add the following alpine.js import to every HTML header. `v3.16.3`

```html
<!-- Alpine suggests to add defer when importing -->
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.16.3/dist/cdn.min.js"></script>
```

Another important dependency is a set of functions and alpine components writen beforehand as reusables.

```html
<script defer src="/assets/src/js/utils/utils.js"></script>
```

To import lucide-icons you must add the following import path to the HTML header.

```html
<!-- IMPORTANT: This route is served by de service 'ServeAssets' directory. --> 
<!-- Do not remove it from the project. -->
<script defer src="/assets/web/lucide-font/lucide.css"></script>
```

Finally to import CSS styles I decided to go with `Bulma CSS`. It is also served by `ServeAssets`.

```html
<!-- -->
<script defer src="/assets/css/bulma/css/bulma.min.css"></script>
```
