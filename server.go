package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"encoding/json"
	"log"
	"context"
	"syscall"
)

// main mux, not for reuse
type mux struct {}

type settings struct {
	RootDir string `json:"root_dir"`
	LogPath string `json:"log_path"`
	Port string `json:"port"`
	AllowedAnyExtensions bool `json:"allowed_any_extensions"`
	AllowedExtensions []string `json:"allowed_extensions"`
	UseCustomNotFoundPage bool `json:"use_custom_not_found_page"`
	CustomNotFoundPagePath string `json:"custom_not_found_page_path"`
}

var execDir string
var s settings
var customNotFoundPageRaw string

func (m mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.EscapedPath()
	log.Println("Request from " + r.RemoteAddr + " " + "[" + URL + "]")

	notFound := false
	data := []byte{}
	if s.AllowedAnyExtensions {
		var err error
		data, err = ioutil.ReadFile(filepath.Join(s.RootDir, URL))
		if err != nil {
			notFound = true
		}
	} else {
		extension := filepath.Ext(URL)
		if extension == "" {
			notFound = true
		} else {
			allowed := false	
			for _, e := range s.AllowedExtensions {
				if e == extension {
					allowed = true
					break
				}
			}
			if !allowed {
				notFound = true
			} else {
				var err error
				data, err = ioutil.ReadFile(filepath.Join(s.RootDir, URL))
				if err != nil {
					notFound = true
				}
			}
		}
	}

	if notFound {
		if s.UseCustomNotFoundPage {
			fmt.Fprint(w, customNotFoundPageRaw)
		} else {
			http.NotFound(w, r)
		}
		log.Println("Response to " + r.RemoteAddr + " " + "[" + URL + "]: Not found.")
		return
	}
	fmt.Fprint(w, string(data))
	log.Println("Response to " + r.RemoteAddr + " " + "[" + URL + "]: Success.")
}

func main() {
	m := mux{}

	setup()
	log.Println("Server initialized.")
	server := http.Server{Addr: ":" + s.Port, Handler: m}

	go waitSignal(&server)
	
	fmt.Println("Listen on port " + s.Port + "...")
	fmt.Println("Press Ctrl-C (or send SIGINT, SIGTERM or SIGHUP to this process) to shutdown...")
	err := server.ListenAndServe()
	if err != nil {
		// maybe exit normally
		log.Println(err)
	}
}

// shutdown server when receive certain signal
func waitSignal(s *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	_ = <- c
	err := s.Shutdown(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

func setup() {
	// read settings.json
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	execDir = filepath.Dir(executable)
	settingsFilePath := filepath.Join(execDir, "settings.json")
	jsonBytes, err := ioutil.ReadFile(settingsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBytes, &s)
	if err != nil {
		log.Fatal(err)
	}
	s.RootDir = absPath(execDir, s.RootDir)
	s.LogPath = absPath(execDir, s.LogPath)

	// load custom not found page
	if s.UseCustomNotFoundPage {
		s.CustomNotFoundPagePath = absPath(execDir, s.CustomNotFoundPagePath)
		data, err := ioutil.ReadFile(s.CustomNotFoundPagePath)
		if err != nil {
			log.Fatal(err)
		}
		customNotFoundPageRaw = string(data)
	}
	
	// set logger
	f, err := os.OpenFile(s.LogPath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

// if p is relative path, return "base + p"
func absPath(base string, p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(base, p)
}