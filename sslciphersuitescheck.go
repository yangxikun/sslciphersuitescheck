package main

import (
    "os"
    "strings"
    "fmt"
    "text/tabwriter"
    "flag"
    "log"
)

var verbose = flag.Bool("v", false, "show client support ciphersuites")
var testclient = flag.String("client", "", "test specified client")
var nosupport = flag.Bool("nosupport", false, "show no support client instead")

func main() {
    log.SetFlags(log.Lshortfile)
    flag.Parse()
    fmt.Println("get clients info...")
    clientsInfo := getSSLClientInfo()
    fmt.Println("parse ciphersuites...")
    cipherSuites := getCipherSuites()
    if len(cipherSuites) == 0 {
        log.Fatal("empty cipher suites")
    }
    fmt.Println("checking...")
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
    if *verbose == false || *nosupport {
        fmt.Fprintln(w, "CLient Id\tClient Name\tPlatform\tClient Version\tSNI\tStapling\tTickets\tNPN\tALPN\t")
    }
    for _, client := range clientsInfo {
        if *testclient != "" && client.Name != *testclient {
            continue
        }
        supportCS := []string{}
        for _, code := range client.SuiteIds {
            cs, ok := cipherSuites[code]
            if ok {
                supportCS = append(supportCS, cs.String())
            }
        }
        if *nosupport == false && len(supportCS) > 0 {
            if *verbose {
                fmt.Fprintln(w, "CLient Id\tClient Name\tPlatform\tClient Version\tSNI\tStapling\tTickets\tNPN\tALPN\t")
            }
            fmt.Fprintln(w, client)
            if *verbose {
                fmt.Fprintln(w, "-\t" + strings.Join(supportCS, "\n-\t"))
            }
        }
        if *nosupport && len(supportCS) == 0 {
            fmt.Fprintln(w, client)
        }
    }
    w.Flush()
}
