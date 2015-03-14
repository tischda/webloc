package main
import (
    "path/filepath"
    "howett.net/plist"
    "fmt"
    "bytes"
    "io/ioutil"
    "os"
    "flag"
)

type weblocHeader struct {
    URL string `plist:"URL"`
    URLN string `plist:"urln"`
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

func decode(filename string) {
    var data weblocHeader

    buf, err := ioutil.ReadFile(filename)
    check(err)

    decoder := plist.NewDecoder(bytes.NewReader([]byte(buf)))
    check(decoder.Decode(&data))
    fmt.Println(data)
}