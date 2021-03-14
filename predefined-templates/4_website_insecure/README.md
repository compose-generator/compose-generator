# Common Website (insecure)
This is the manual document for the `4_website_insecure` stack.

## Description
This stack deploys a one-container website over HTTP to the web. Please only use this template for test purposes. We recommend to use the template `3_website` stack instead.

## Usage instructions
In the most cases, you want the website to be visible to anyone on the global web. Just deploy it with Docker Compose by executing `$ docker-compose up`. After that, your website should be exposed to port 80 / your custom port.

## Credits
Template created & maintained by @marcauberer