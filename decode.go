package main
import (
    "path/filepath"
    "howett.net/plist"
    "fmt"
    "os"
    "flag"
)

type weblocHeader struct {
    URL string `plist:"URL"`
}

func main() {
    flag.Parse()
    root := flag.Arg(0)
    filepath.Walk(root, walkpath)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func walkpath(path string, f os.FileInfo, err error) error {
    if err != nil {
        return nil
    }
    matched, err := filepath.Match("*.webloc", f.Name())
    if err != nil {
        return err
    }
    if matched {
        decode(path)
    }
    return nil
}

func decode(path string) {
    var data weblocHeader

    f, err := os.Open(path)
    check(err)
    defer f.Close()

    decoder := plist.NewDecoder(f)
    check(decoder.Decode(&data))
    fmt.Println(data)
}