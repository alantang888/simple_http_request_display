# simple_http_request_display
[![Docker Build](https://img.shields.io/docker/build/alantang888/simple_http_request_display.svg)][docker-hub]
[![Docker Pulls](https://img.shields.io/docker/pulls/alantang888/simple_http_request_display.svg?maxAge=604800)][docker-hub]

This is a very simple http echo program to listen a TCP port 8080 and return requester IP:Port, request path, headers and cookies. Also log those information in log.
But not work behide proxy, it will return proxy address.

There has a docker image [`alantang888/simple_http_request_display`][docker-hub]

### Install by Helm
You can install this by `helm install -f helm_chart/sample_vaule.yaml -n http-echo-test helm_chart/simple_http_request_display`.

Please remember change the ingress hosts value.

[docker-hub]: https://hub.docker.com/r/alantang888/simple_http_request_display/
