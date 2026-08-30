![image](assets/img/saavedraVanilla.png)

# About it

## Quick Start

To configure Saavedraa you must provide the following environment variables.

```env
# THIS TOKEN SIGNS EVERY JWT SENDED TO FRONTEND GENERATING SESSION STATES.
SESSION_TOKEN=
# SPECIFIC NAME FOR YOU DATABASE
DATABASE_FILE_NAME=
# ADMINISTRATOR CREDENTIALS
ADMIN_NAME=
ADMIN_EMAIL=
ADMIN_PASSWORD=
# SET TO TRUE IN PRODUCTION
IS_PRODUCTION=0
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

Finally to import CSS styles I decided to go with `Bulma CSS`. It is also served by `ServeAssets` directory.

```html
<!-- It's important to add dynamic window resizing. Suggested by bulma -->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<script defer src="/assets/css/bulma/css/bulma.min.css"></script>
```

## Debugging

Saavedra suggest debugging by CLI. The following is the current debugging tool used by Saavedra `dlv`.

```bash
# Initializing debug mode.
dlv debug main.go

# Adding a breakpoint on line 30. (Must have content or must not be a commented line).
# # Otherwise dlv won't create the breakpoint
(dlv) break utils/utils.go:30
```

```bash
# Breakpoint by calling a function.
(dlv) break utils.FunctionName
```
