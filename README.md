# msws
msws stands for **M**ini **S**tatic **W**eb **S**erver, a static HTTP server for serving static content with some simple configuration. Written in Go.

## Build and Run
Use [go get](https://golang.org/pkg/cmd/go/internal/get/) or `git clone` to download the repository. Then,
```
$ go build -o server.go && ./server
```

## Configuration
**Note**: `settings.json` should be placed under the same location with the `server` executable.

**Note**: All paths in `settings.json` can be absolute or relative, and **relative paths are relative to the `server` executable's located directory**.

(e.g., if executable is under `/home/user/msws`, path `./somewhere/404.html` will refer to `/home/user/msws/somewhere/404.html`)

### settings.json
* `root_dir`: `string` Root directory of your website.

* `log_path`: `string` Log file path.

* `port`: `string` The port number for the server to listen on.

* `allowed_any_extensions`: `boolean` Serve any extension files or not.

* `allowed_extensions`: `array` Only these extensions can be served, otherwise, the not found page will be responsed.

* `use_custom_not_found_page`: `boolean` Use custom not found page instead of default not found page.

* `custom_not_found_page_path`: `string` The path to your custom not found page file.