# Common Website
This is the manual document for the `3_website` stack.

## Description
This stack deploys a one-container website with a secure wrapper. This means, that it has support for communication over HTTPS.

## Usage instructions
In the most cases, you want the website to be visible to anyone on the global web. See below for the deployment in home or company networks without access to the web.
Just deploy it with Docker Compose by executing `$ docker-compose up`. After that, your website should be exposed to port 443. The generation of the TLS certificate might take a while for the first start. Please watch the log output of `$ docker-compose up` for possible error messages.

### Custom routing
You can customize your routing configuration by adding a vhost configuration file to the `proxy-vhosts` volume. This file must named like your domain (If your domain is `example.com`, the file has to be `example.com`). There you can define custom path configurations like this:

```
location /test/ {
    proxy_pass http://website-api:5000/;
}

location /db/ {
    proxy_pass http://phpmyadmin:80/;
}
```

### Usage in home or company networks
If you want to deploy a website in your home network or company network, please note, that this stack works with certificate generation by LetsEncrypt. LetsEncrypt will perform challenges to check the domain ownership. This challenges can not work, if your server is behind a firewall and not reachable by LetsEncrypt. In this case, you have to issue your TLS certificates yourself and paste them into the `proxy-certs` volume. The certificates must have the following naming schema: <br>
If your domain is `example.com`, your certificate files have to be named `example.com.crt`, `example.com.csr`, `example.com.key`.

## Volumes
-   `proxy-certs`: This is where you have to paste your certificates or where LetsEncrypt will put the certificates
-   `proxy-config`: 
-   `proxy-html`: This directory is used for LetsEncrypt ownership challenges. If you do not use LetsEncrypt, you can remove this volume
-   `proxy-vhosts`: 

## Credits
Template created & maintained by @marcauberer