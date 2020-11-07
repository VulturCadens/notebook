# Nginx

Sending precompressed files with the “.gz” filename extension instead of regular files.

### Module

```console
nginx -V
```

Should be something like **--with-http_gzip_static_module**.

### Configuration file

```console
gzip_static  on;
```

### Gzip

```console
gzip -k file.ext
```

**-k** Keep input files during compression.