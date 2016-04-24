```
fastcgi mode
http mode
```

nginx -> go-import-server:

###nginx.conf:

```
worker_processes  auto;

events {
    worker_connections  10240;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    access_log off;

    sendfile        on;

    keepalive_timeout  65;

    upstream importfastcgi {
        server 127.0.0.1:15001;
        keepalive 300;
    }

    upstream importhttp {
        server 127.0.0.1:15002;
        keepalive 300;
    }

    server {
        listen       80;
        server_name  import.fastcgi.ubuntu;

        location / {
            fastcgi_pass    importfastcgi;
            fastcgi_keep_conn on;
            include         fastcgi_params;
        }
    }

    server {
        listen       80;
        server_name  import.http.ubuntu;

        location / {
            proxy_pass    http://importhttp;
            proxy_http_version 1.1;
            proxy_set_header Connection "";
        }
    }
}
```

###/etc/hosts
```
127.0.0.1	import.fastcgi.ubuntu import.http.ubuntu
```

###go-import-server start:
```
GODEBUG=gctrace=1 GIN_MODE=release ./go-import-server 2>&1 > /dev/null
```

###Benchmark tests:

```
./wrk -t100 -c5000 -d30s http://import.fastcgi.ubuntu/go/lib/gin-startup
```

```
./wrk -t100 -c5000 -d30s http://import.http.ubuntu/go/lib/gin-startup
```

```
./wrk -t100 -c5000 -d30s http://127.0.0.1:15002/go/lib/gin-startup
```

###Conclusion

better to use http proxy.
