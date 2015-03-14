package main
import (
    "howett.net/plist"
    "fmt"
    "bytes"
    "io/ioutil"
)
func main() {

    type weblocHeader struct {
        URL string `plist:"URL"`
        URLN string `plist:"urln"`
    }
    var data weblocHeader

    buf, err := ioutil.ReadFile("../binary-plist.webloc")
    // buf, err := ioutil.ReadFile("../xml-content.webloc")
    check(err)

    decoder := plist.NewDecoder(bytes.NewReader([]byte(buf)))
    check(decoder.Decode(&data))
    fmt.Println(data)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}