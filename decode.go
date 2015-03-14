package main
import (
    "howett.net/plist"
    "os"
    "fmt"
    "bytes"
    "io/ioutil"
)
func main() {
    fmt.Println("--------- Enconding:")
    encoder := plist.NewEncoder(os.Stdout)
    encoder.Encode(map[string]string{"hello": "world"})

    fmt.Println("\n--------- Decoding:")

    type sparseBundleHeader struct {
        InfoDictionaryVersion string `plist:"CFBundleInfoDictionaryVersion"`
        BandSize              uint64 `plist:"band-size"`
        BackingStoreVersion   int    `plist:"bundle-backingstore-version"`
        DiskImageBundleType   string `plist:"diskimage-bundle-type"`
        Size                  uint64 `plist:"size"`
    }

    buf := bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>CFBundleInfoDictionaryVersion</key>
		<string>6.0</string>
		<key>band-size</key>
		<integer>8388608</integer>
		<key>bundle-backingstore-version</key>
		<integer>1</integer>
		<key>diskimage-bundle-type</key>
		<string>com.apple.diskimage.sparsebundle</string>
		<key>size</key>
		<integer>4398046511104</integer>
	</dict>
</plist>`))

    var data sparseBundleHeader
    decoder := plist.NewDecoder(buf)
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data)


    fmt.Println("--------- Decoding:")
    
    type weblocHeader struct {
        URL string `plist:"URL"`
    }
    var data2 weblocHeader

    buf2, err2 := ioutil.ReadFile("../binary-plist.webloc")
    check(err2)

    decoder2 := plist.NewDecoder(bytes.NewReader([]byte(buf2)))
    err3 := decoder2.Decode(&data2)
    if err3 != nil {
        fmt.Println(err3)
    }
    fmt.Println(data2)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}