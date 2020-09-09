# msws
msws stands for **M**ini **S**tatic **W**eb **S**erver, a static HTTP server for serving static content with some simple configuration. Written in Go.

## Build and Run
```
$ go build -o server.go && ./server
```

## Configuration
**Note**: `settings.json` should be placed under the same location with the `server` executable.

**Note**: All path in `settings.json` can be absolute or relative, and **relative paths are relative to the `server` executable's located directory**.

(e.g., if executable is under `/home/user/msws`, path `./somewhere/404.html` will refer to `/home/user/msws/somewhere/404.html`)

* `root_dir`: Root directory of your website.

* `log_path`: Log file path.

* `port`: Port number for the server to listen on. Should be string type.

* `allowed_any_extensions`: Boolean type. Set to true won't check and block any kind of file extensions.

* `allowed_extensions`: A list of strings. If the request URL are not end with one of these extensions, response not found.

* `use_custom_not_found_page`: Boolean type. If true, server will response content of "custom not found page file".

* `custom_not_found_page_path`: Path to your "custom not found page file".