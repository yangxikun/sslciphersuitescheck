package main

import (
    "fmt"
    "strings"
    "encoding/hex"
    "errors"
    "bufio"
    "os"
)

type cipherSuite struct {
    code uint32
    id string
    name string
    sslVer string
    kx string
    au string
    enc string
    mac string
    export bool
}

func (cs cipherSuite) String() string {
    return fmt.Sprintf("%s %s %s %s %s %s %s %v", cs.id, cs.name, cs.sslVer, cs.kx, cs.au,
        cs.enc, cs.mac, cs.export)
}

func parseCode(text string) (uint32, error) {
    text = strings.Replace(text, "0x", "", -1)
    text = strings.Replace(text, ",", "", -1)
    hc, err := hex.DecodeString(text)
    if err != nil {
        return 0, err
    }
    var code uint32
    switch len(hc) {
    case 2:
        code = uint32(hc[1]) | uint32(hc[0])<<8
    case 3:
        code = uint32(hc[2]) | uint32(hc[1])<<8 | uint32(hc[0])<<16
    default:
        return 0, errors.New("parse code fail")
    }
    return code, nil
}

func getCipherSuites() map[uint32]cipherSuite {
    cs := make(map[uint32]cipherSuite)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    for len(text) > 0 {
        cipher := strings.Fields(text)
        if len(cipher) < 8 {
            fmt.Println("wrong ciper:", text)
            text, _ = reader.ReadString('\n')
            continue
        }
        code, err := parseCode(cipher[0])
        if err != nil {
            fmt.Println("cannot parse code:", text)
            text, _ = reader.ReadString('\n')
            continue
        }
        c := cipherSuite{}
        c.id = cipher[0]
        c.code = code
        c.name = cipher[2]
        c.sslVer = cipher[3]
        c.kx = cipher[4]
        c.au = cipher[5]
        c.enc = cipher[6]
        c.mac = cipher[7]
        if len(cipher) == 9 && cipher[8] == "export" {
            c.export = true
        }
        cs[code] = c
        text, _ = reader.ReadString('\n')
    }
    return cs
}
