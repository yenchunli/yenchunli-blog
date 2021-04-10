---
title: "Install ngx_http_geoip_module for ip analysis"
date: 2020-04-14T12:18:32+08:00
draft: true

summary: "Ngx-http-geoip-module can enable Internet traffic analysis for indivisual conuntry."

tags: ["Nginx"]

---

We want to calculate traffic for indivisual country using GeoIP. But it requires [ngx_http_geoip_module](http://nginx.org/en/docs/http/ngx_http_geoip_module.html). This module is not built by default, it should be enabled with the --with-http_geoip_module configuration parameter. Moreover, this module depends on `MaxMind GeoIP` library. Unfortuantely, less tutorials talk about this.

I will try to install it using `apt-get`

## Install steps

1. Search for geoip package
```bash
apt-cache search geoip
```

2. Install libgeoip-dev
```bash
sudo apt-get install libgeoip-dev
```

3. Add configure settings when you build Nginx
```bash
./configure \
--with-http_geoip_module \
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

