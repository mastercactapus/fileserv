package main

//go:generate go-bindata assets/

import (
	"fmt"
	"github.com/bradfitz/http2"
	"github.com/spf13/cobra"
	"html/template"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var bindAddr string
var useMdns bool
var dirName string
var fileHandler http.Handler
var mainCmd = &cobra.Command{
	Use:   "fileserv",
	Short: "Quick and easy fileserver",
	Run:   serve,
}

type fileservHandler struct {
	http.Handler
	dir  http.Dir
	tmpl *template.Template
}

type Page struct {
	Path      string
	ParentUrl string
	Contents  FileList
}
type File struct {
	Icon         string
	IsVideo      bool
	IsAudio      bool
	IsDir        bool
	Name         string
	Modified     string
	ModifiedTime time.Time
	Size         string
	SizeNum      int64
	Url          string
}

type FileList struct {
	Files          []File
	SortByName     bool
	SortByModified bool
	SortByType     bool
	SortBySize     bool
	Reverse        bool
	ReverseFlip    string
	SortMethod     string
}

const (
	SORT_NAME     = "name"
	SORT_MODIFIED = "modified"
	SORT_SIZE     = "size"
	SORT_TYPE     = "type"

	KiB_Limit = 1000
	KiB       = 1024
	MiB_Limit = KiB * 1000
	MiB       = KiB * 1024
	GiB_Limit = MiB * 1000
	GiB       = MiB * 1024
	TiB_Limit = GiB * 1000
	TiB       = GiB * 1024
)

func (f FileList) Swap(a, b int) {
	f.Files[a], f.Files[b] = f.Files[b], f.Files[a]
}
func (f FileList) Len() int {
	return len(f.Files)
}
func (f FileList) Less(a, b int) bool {
	if f.Files[a].IsDir != f.Files[b].IsDir {
		return f.Files[a].IsDir
	}
	var result bool
	if f.SortByModified {
		result = f.Files[a].ModifiedTime.Before(f.Files[b].ModifiedTime)
	} else if f.SortBySize {
		result = f.Files[a].SizeNum < f.Files[b].SizeNum
	} else if f.SortByType {
		result = strings.ToLower(filepath.Ext(f.Files[a].Name)) < strings.ToLower(filepath.Ext(f.Files[b].Name))
	} else {
		result = strings.ToLower(f.Files[a].Name) < strings.ToLower(f.Files[b].Name)
	}
	if f.Reverse {
		result = !result
	}
	return result
}

func serveAsset(w http.ResponseWriter, name string) {
	data, err := Asset("assets/" + name)
	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(404)
		io.WriteString(w, err.Error())
	} else {
		w.Header().Set("content-type", mime.TypeByExtension(filepath.Ext(name)))
		w.Write(data)
	}
}

func prettySize(size int64) string {
	var unit string
	var value float64
	if size >= TiB_Limit {
		unit = "TiB"
		value = float64(size) / float64(TiB)
	} else if size >= GiB_Limit {
		unit = "GiB"
		value = float64(size) / float64(GiB)
	} else if size >= MiB_Limit {
		unit = "MiB"
		value = float64(size) / float64(MiB)
	} else if size >= KiB_Limit {
		unit = "KiB"
		value = float64(size) / float64(KiB)
	} else {
		return fmt.Sprintf("%d Bytes", size)
	}
	return fmt.Sprintf("%.2f %s", value, unit)
}

func (h fileservHandler) serveDir(w http.ResponseWriter, req *http.Request, contents []os.FileInfo) {
	page := Page{Path: req.URL.Path}
	if req.URL.Path != "/" {
		parts := strings.Split(req.URL.Path, "/")
		page.ParentUrl = strings.Join(parts[:len(parts)-2], "/") + "/"
	}
	files := make([]File, 0, len(contents))
	for _, v := range contents {
		if v.Name()[0] == '.' {
			//don't list hidden
			continue
		}
		file := File{
			IsDir: v.IsDir(),
			Name:  v.Name(),
		}

		file.Modified = v.ModTime().Format("Mon, Jan 2 2006 3:04pm")
		if v.IsDir() {
			file.Size = "-"
			file.Url = req.URL.Path + v.Name() + "?" + req.URL.RawQuery
		} else {
			file.Size = prettySize(v.Size())
			file.Url = req.URL.Path + v.Name()
		}

		files = append(files, file)
	}
	switch req.URL.Query().Get("sort") {
	case SORT_MODIFIED:
		page.Contents.SortByModified = true
	case SORT_SIZE:
		page.Contents.SortBySize = true
	case SORT_TYPE:
		page.Contents.SortByType = true
	default:
		page.Contents.SortByName = true
	}
	if req.URL.Query().Get("reverse") == "true" {
		page.Contents.Reverse = true
		page.Contents.ReverseFlip = "false"
	} else {
		page.Contents.ReverseFlip = "true"
	}

	page.Contents.Files = files
	sort.Sort(page.Contents)
	w.Header().Set("content-type", "text/html")
	err := h.tmpl.Execute(w, page)
	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(500)
		io.WriteString(w, err.Error())
	}
}

func (h fileservHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" && req.URL.Query().Get("asset") != "" {
		serveAsset(w, req.URL.Query().Get("asset"))
	} else if strings.HasSuffix(req.URL.Path, "/") {
		info, err := h.dir.Open(req.URL.Path)
		if err != nil {
			w.WriteHeader(404)
			io.WriteString(w, err.Error())
			return
		}
		contents, err := info.Readdir(-1)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, err.Error())
			return
		}
		h.serveDir(w, req, contents)
	} else {
		h.Handler.ServeHTTP(w, req)
	}
}

func serve(cmd *cobra.Command, args []string) {
	s := new(http.Server)
	s.Addr = bindAddr
	s.SetKeepAlivesEnabled(true)
	dir := http.Dir(dirName)
	fs := fileservHandler{http.FileServer(dir), dir, nil}

	indexData, err := Asset("assets/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	fs.tmpl, err = template.New("index").Parse(string(indexData))
	if err != nil {
		log.Fatalln(err)
	}

	s.Handler = fs
	http2.ConfigureServer(s, nil)
	log.Fatalln(s.ListenAndServe())
}

func main() {
	mainCmd.Flags().StringVarP(&bindAddr, "bind", "b", ":8000", "Bind address to serve from.")
	mainCmd.Flags().BoolVarP(&useMdns, "mdns", "m", true, "Advertise via MDNS.")
	mainCmd.Flags().StringVarP(&dirName, "dir", "d", ".", "Directory to serve.")
	mainCmd.Execute()
}
