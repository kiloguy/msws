# msws
msws stands for **M**ini **S**tatic **W**eb **S**erver, a static HTTP server for serving static content with some simple configuration.

## Build and Run
```
$ go build -o server.go && ./server
```

## Configuration
**Note**: `settings.json` should be placed under the same location with the `server` executable.

**Note**: All path in `settings.json` can be absolute or relative, and **relative paths are relative to the `server` executable's located directory**.
(e.g., if executable is under `/home/user/msws`, path `./somewhere/404.html` will refer to `/home/user/msws/somewhere/404.html`)