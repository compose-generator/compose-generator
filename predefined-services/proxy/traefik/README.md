## Traefik Reverse Proxy
Traefik is an open-source edge router that supports automatic service detection/exposure and fully-automated TLS certificate issuing/renewal. Thus, it can be used as single-service reverse proxy without a TLS helper service.

## Setup
**Warning**: Per default, the Traefik web UI is *not* password protected. Do not use it like that in production under no circumstances. For production use, please disable the web UI or setup password protection.