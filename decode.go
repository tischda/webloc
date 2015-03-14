package main
import (
    "path/filepath"
    "howett.net/plist"
    "os"
    "flag"
    "fmt"
)

type weblocHeader struct {
    URL string `plist:"URL"`
}

var delete bool
var noop bool

func init() {
    flag.BoolVar(&delete, "delete", false, "delete .webloc files after conversion")
    flag.BoolVar(&noop, "noop", false, "decode urls, but do not change file system")
    flag.Parse()
}

func main() {
    root := flag.Arg(0)
    filepath.Walk(root, walkpath)
}

func walkpath(path string, f os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    matched, err := filepath.Match("*.webloc", f.Name())
    if err != nil {
        return err
    }
    if matched {
        process(path)
    }
    return nil
}

func process(path string) {
    url := decode(path)
    fmt.Println(url)

    if ! noop {
        writeUrl(path[:len(path)-len(".webloc")] + ".url", url)
        if delete {
            err := os.Remove(path)
            check(err)
        }
    }
}

func decode(path string) string {
    var data weblocHeader

    f, err := os.Open(path)
    check(err)
    defer f.Close()

    decoder := plist.NewDecoder(f)
    check(decoder.Decode(&data))
    return data.URL
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeUrl(path string, url string) {
    f, err := os.Create(path)
    check(err)
    defer f.Close()

    f.WriteString("[InternetShortcut]\nURL=" + url)
}
