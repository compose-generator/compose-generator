# Common Website with PHP support (insecure)
This is the manual document for the `6_website_php_insecure` stack.

## Description
This stack deploys a one-container website with PHP support over HTTP to the web. Please only use this template for test purposes. We recommend to use the template `5_website_php` stack instead.

## Usage instructions
In the most cases, you want the website to be visible to anyone on the global web. Just deploy it with Docker Compose by executing `$ docker-compose up`. After that, your website should be exposed to port 80 / your custom port.

## Credits
Template created & maintained by @marcauberer