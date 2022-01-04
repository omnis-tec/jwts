# GmsTemp

### Env variables:

```
DEBUG: true
LOG_LEVEL: "debug"
HTTP_LISTEN: ":80"
KID: "key-id"
PRIVATE_PEM: "path"
PUBLIC_PEM: "path"
```

<br/>

### Install `swagger-cli`:

```
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
rm -rf "$dir"
```
