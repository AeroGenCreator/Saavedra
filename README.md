# Saavedra

_Así como Don Quijote, me hallo luchando contra molinos de viento. ¿Serán reales?_.

# About it

Saavedra was born as a response to a lack of opportunities in the traditional job market, but more importantly, as an initiative to turn that situation into an opportunity for building software. Despite having developed a solid understanding of software engineering, I recognized a growing demand for highly customized applications alongside a significant reduction in the cost of developing them, largely driven by generative AI.

![image](assets/img/GitHub/saavLogin.png)

This led me to propose an architecture focused on **code reuse, minimal metaprogramming, explicit contracts, and clear separation of responsibilities**, rather than relying heavily on generalized abstractions. The backend exposes REST APIs built around the **Repository Pattern**, while the system deliberately separates volatile application data from persistent historical data.

Instead of treating relational database modeling as something that must be redesigned for every client requirement, Saavedra uses **JSON as a source of volatile configuration and application data**, while the database remains the **historical source of truth**. This approach reduces the complexity involved in constantly adapting relational schemas to highly specific client requirements.

The architecture also accepts a certain degree of code repetition. In a traditional software development environment, excessive repetition is generally considered undesirable. However, when AI can assist with generating, refactoring, and maintaining repetitive functionality at a significantly lower cost, the trade-off changes. In this context, some duplication can be preferable to introducing additional abstraction and architectural complexity solely for the sake of reuse.

A central idea behind Saavedra is that generative AI may change the economics of software construction. If implementation becomes significantly cheaper, it may become economically reasonable to replace part of a highly generalized architecture with more specialized, client-specific implementations—as long as the system preserves explicit contracts, dependency boundaries, and the persistence of relevant facts.

This does not mean that Saavedra is designed to allow AI-generated code to modify the entire system without constraints. On the contrary, the architecture is intentionally divided into relatively independent layers so that specific components can be rewritten, adapted, or extended without unnecessarily breaking unrelated services. The objective is to create an environment where generative AI can assist implementation while the overall architecture continues to protect system boundaries.

The complete architecture is organized into loosely coupled layers:

`DB → Store → Service → App → Router → HTML/CSS → JS → Request`

Each layer has a specific responsibility and can evolve independently. The **Service layer** acts as the main coordination point: it can request analytical operations from Python or retrieve configuration and volatile data from JSON sources, without forcing the rest of the system to depend directly on those implementation details.

The backend is built primarily with **Go and Python**, while the frontend uses **HTML, CSS, and JavaScript**. Authentication and API protection are handled through **JWT**. The application follows a hybrid approach combining **REST APIs and Server-Side Rendering (SSR)**, while JavaScript is used to manage the dynamic behavior of the DOM and provide a more interactive client experience.

## Generative AI and My Development Process

There is an important distinction between the philosophy behind Saavedra and my personal approach to writing its code.

Although the system is intentionally designed to take advantage of generative AI-assisted development, I do not currently rely on generative AI to write most of my code. I primarily use AI as a **rapid source of technical consultation**, documentation, conceptual clarification, and architectural discussion.

The implementation itself has largely been written manually.

This is not a rejection of AI-assisted programming. It is a personal engineering preference based on how I currently work. In many situations, I find that writing a piece of code myself is more efficient than generating it with an AI system and then spending time reviewing, understanding, correcting, and integrating the generated output.

My reasoning is relatively simple: if I write the code myself, I have already gone through the process of reading and understanding it while producing it. The time saved by delegating implementation to an AI may sometimes be partially offset by the time required to inspect the generated code, verify its assumptions, identify incorrect details, and adapt it to the existing system.

Therefore, at least for the way I currently work, manual implementation can be more efficient because the act of writing the code is simultaneously part of the process of understanding and validating it.

However, I do not consider this a universal conclusion.

This is currently a **personal hypothesis rather than an established result**. I have not yet conducted formal studies, collected sufficient statistical data, or developed a complete thesis around this approach. One of my long-term goals is to investigate this question more rigorously.

In particular, I am interested in evaluating whether AI-assisted software development actually produces a measurable economic advantage when factors such as the following are considered:

* Time required to generate code.
* Time required to review and understand generated code.
* Time spent correcting implementation errors.
* Integration costs.
* Maintenance costs.
* Token and computational costs.
* Defect rates.
* Development speed.
* The degree of architectural coupling introduced by generated code.

The central question is not simply whether AI can write code faster than a developer. The more interesting question is whether **the total cost of producing, understanding, validating, integrating, and maintaining that code is actually lower**.

Saavedra could eventually become a practical environment for investigating this question. Its architecture is partly based on the assumption that AI can make specialized and partially redundant implementations economically viable, provided that architectural boundaries remain stable and explicit.

However, this remains an idea that requires empirical validation.

Before pursuing formal research around these questions, my immediate objective is more practical: **to commercialize the system, work with real software requirements, and establish a sustainable source of income through its development**. Real-world use would also provide a more meaningful foundation for future research by allowing architectural decisions to be evaluated against actual development costs, maintenance requirements, and client-specific implementations.

The objective of Saavedra is therefore not simply to demonstrate that I can write code. It is intended to demonstrate my progression toward **full-stack software engineering**, including database design, backend architecture, API design, authentication, frontend development, system decomposition, and the trade-offs involved in building maintainable software.

I would not describe myself as a senior engineer. However, I have developed an understanding that goes beyond the implementation of individual functions or features. I am increasingly interested in the underlying computational and architectural principles that make software systems reliable, adaptable, economically viable, and maintainable.

In that sense, Saavedra is both a software project and a practical engineering hypothesis. Its architecture assumes that the role of the engineer is increasingly centered on defining **boundaries, contracts, constraints, and system structure**, while implementation can be accelerated—and in some cases partially replaced—by generative AI.

At the same time, my own development process demonstrates another side of that hypothesis: even in a future where AI can generate increasingly large portions of software, understanding how and when a human developer should delegate implementation remains an open engineering and economic question.

**The engineer defines the architecture and its constraints. AI can assist the implementation. The challenge is determining when that assistance actually reduces the total cost of software development.**

## Current Implementation Notice

Saavedra is currently implemented for a specific technology business, where it is being used primarily to manage quotation-related workflows. This implementation should not be interpreted as the architectural limit of the system.

The purpose of the project is precisely to allow business-specific functionality to be removed, replaced, or rewritten while preserving the services and architectural components that remain relevant to a different type of business. In other words, the current implementation represents one concrete application of the architecture rather than a restriction on its potential use.

When adapting Saavedra to another business, technology-specific components can be removed while reusable services, contracts, and infrastructure can remain in place wherever they are appropriate. New business-specific functionality can then be implemented without requiring the entire system to be rebuilt from scratch.

Most importantly, the architecture is not purely theoretical. **There is already a real business using the system**, providing an initial practical validation of the approach and a foundation for its continued evolution.

![image](assets/img/GitHub/saavMenu.png)
![image](assets/img/GitHub/saavList.png)
![image](assets/img/GitHub/saavQuote.png)

## Quick Start

To configure Saavedraa you must provide the following environment variables.

```env
# Special token to validate request from the client.
SESSION_TOKEN=
# Specific name for your Sqlite3 Database
DATABASE_FILE_NAME=
# Administrator credentials
ADMIN_NAME=
ADMIN_EMAIL=
ADMIN_PASSWORD=
# SET TO FALSE IN PRODUCTION
IS_PRODUCTION=
# QUANTITY OF RECORDS SHOWN/FETCH BY LISTS VIEW
RECORDS_PER_SLICE=
```

## Dependencies

Either to develop or run Saavedra it is important to add the following header.

```html
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/assets/web/lucide-font/lucide.css">
<link rel="stylesheet" href="/assets/css/savedraaCSS.css">
<link rel="stylesheet" href="/assets/css/bulma/css/bulma.min.css">
<script defer src="/assets/src/js/utils/utils.js"></script>
<script defer src=""></script> <!-- JavaScript For current HTML File -->
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.16.3/dist/cdn.min.js"></script>
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

## Service Dependencies

Some services relay on others. To keep a map of them you can point them as follows. Furthermore, keep migration order when loading services... For example: First you load `Login - Routes And Migrations` then `Users - Router and Migration`.

```txt
All -> 'depends on' -> Welcome and ServeAssets
Users ->'depends on' -> Login
Quote ->'depends on' -> Product
Quote ->'depends on' -> Proveedor
Quote ->'depends on' -> Customer
```
