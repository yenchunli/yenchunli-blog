---
title: "[Nginx] Install nginx-module-vts for monitoring"
date: 2020-04-14T10:23:09+08:00
draft: false

summary: "Nginx-module-vts is designed to monitor Internet traffic on Nginx."

tags: ["Nginx"]
---

## Environment

- Ubuntu 16.04

## 1. Download Nginx Source Code

[Nginx source code download](https://nginx.org/en/download.html)

We download the stable version.

```
cd /opt/local
wget https://nginx.org/download/nginx-1.16.1.tar.gz
tar xvf nginx-1.16.1.tar.gz
sudo chown -R <user>:<group> nginx-1.16.1
```

## 2. Download nginx-module-vts

[Nginx-module-vts release](https://github.com/vozlt/nginx-module-vts/releases)

```
cd /opt/local
wget https://github.com/vozlt/nginx-module-vts/archive/v0.1.18.tar.gz
tar xvf v0.1.18.tar.gz
```

## 3. Build Nginx source code

Check the dependencies and the settings
```
./configure
```

Usually you need to install PCRE, zlib, Openssl to support `nginx rewrite` , `gzip` and `https`.

```
sudo apt-get install libpcre3 libpcre3-dev
sudo apt-get install zlib1g-dev
sudo apt-get install openssl libssl-dev
```

```
./configure \
--prefix=/etc/nginx \
--sbin-path=/usr/sbin/nginx \
--modules-path=/usr/lib64/nginx/modules \
--conf-path=/etc/nginx/nginx.conf \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--pid-path=/var/run/nginx.pid \
--lock-path=/var/run/nginx.lock \
--http-client-body-temp-path=/var/cache/nginx/client_temp \
--http-proxy-temp-path=/var/cache/nginx/proxy_temp \
--http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp \
--http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp \
--http-scgi-temp-path=/var/cache/nginx/scgi_temp \
--user=www-data \
--group=www-data \
--with-compat \
--with-file-aio \
--with-threads \
--with-http_addition_module \
--with-http_auth_request_module \
--with-http_dav_module \
--with-http_flv_module \
--with-http_gunzip_module \
--with-http_gzip_static_module \
--with-http_mp4_module \
--with-http_random_index_module \
--with-http_realip_module \
--with-http_secure_link_module \
--with-http_slice_module \
--with-http_ssl_module \
--with-http_stub_status_module \
--with-http_sub_module \
--with-http_v2_module \
--with-mail \
--with-mail_ssl_module \
--with-stream \
--with-stream_realip_module \
--with-stream_ssl_module \
--with-stream_ssl_preread_module \
--with-http_geoip_module \
--with-cc-opt='-O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector-strong --param=ssp-buffer-size=4 -grecord-gcc-switches -m64 -mtune=generic -fPIC' \
--with-ld-opt='-Wl,-z,relro -Wl,-z,now -pie' \
--add-dynamic-module=../nginx-module-vts-0.1.18/
```

it will show the configs we are about to use in `Makefile`

```
Configuration summary
  + using threads
  + using system PCRE library
  + using system OpenSSL library
  + using system zlib library

  nginx path prefix: "/etc/nginx"
  nginx binary file: "/usr/sbin/nginx"
  nginx modules path: "/usr/lib64/nginx/modules"
  nginx configuration prefix: "/etc/nginx"
  nginx configuration file: "/etc/nginx/nginx.conf"
  nginx pid file: "/var/run/nginx.pid"
  nginx error log file: "/var/log/nginx/error.log"
  nginx http access log file: "/var/log/nginx/access.log"
  nginx http client request body temporary files: "/var/cache/nginx/client_temp"
  nginx http proxy temporary files: "/var/cache/nginx/proxy_temp"
  nginx http fastcgi temporary files: "/var/cache/nginx/fastcgi_temp"
  nginx http uwsgi temporary files: "/var/cache/nginx/uwsgi_temp"
  nginx http scgi temporary files: "/var/cache/nginx/scgi_temp"
```

## 4. Makefile and move compiled module to right place

```bash
# Compile the nginx source code
make
# Move compiled files to the location where we configure
make install
# Move compiled *.so file to place to let /etc/nginx/nginx.conf use this module
sudo cp objs/ngx_http_vhost_traffic_status_module.so /etc/nginx/modules/
```

## Reference

[install-nginx-virtual-host-traffic-status-module](https://ahelpme.com/software/nginx/install-nginx-virtual-host-traffic-status-module-traffic-information-in-nginx-and-more-per-server-block-and-upstreams/)


