## Live-Poll
Live-Poll is an open-source platform for taking surveys or polls. There is an [free hosted version online](https://www.live-poll.de) for the common use, but can also be self-hosted. This Compose Generator configuration is a good starting point for the usage of Live-Poll.

### Setup
Live-Poll comes as two parts: A frontend (this service) and an API (categorized as backend service).

This frontend service, has to be built on the server where it's going to be live, because due to it's single-page-application behavior, the API url has to be baked into the application at build time. This is why the "demo app creation" step takes so long.