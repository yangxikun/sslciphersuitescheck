package main

import (
    "fmt"
    "strings"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
)

type clientInfo struct {
    Id uint
    Name string
    Platform string
    Version string
    SuiteIds []uint32
    SupportsSni bool
    SupportsStapling bool
    SupportsTickets bool
    NpnProtocols []string
    AlpnProtocols []string
}

func (c clientInfo) String() string {
    return fmt.Sprintf("%d\t%s\t%s\t%s\t%v\t%v\t%v\t%s\t%s", c.Id, c.Name, c.Platform, c.Version, c.SupportsSni,
        c.SupportsStapling, c.SupportsTickets, strings.Join(c.NpnProtocols, ","),
        strings.Join(c.AlpnProtocols, ","))
}

func fromApi() []clientInfo {
    resp, err := http.Get("https://api.ssllabs.com/api/v3/getClients")
    if err != nil {
        log.Println("http request error.", err)
        return nil
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Println("unable to read response body.", err)
        return nil
    }
    var clients []clientInfo
    err = json.Unmarshal(body, &clients)
    if err != nil {
        log.Println("https://api.ssllabs.com/api/v3/getClients response wrong json data.", err)
        return nil
    }

    return clients
}

func fromLocal() []clientInfo {
    var clients []clientInfo
    err := json.Unmarshal([]byte(ssllabsClientsInfo), &clients)
    if err != nil {
        log.Println(err)
        return nil
    }
    return clients
}

func getSSLClientInfo() []clientInfo {
    clients := fromApi()
    if len(clients) == 0 {
        fmt.Println("Warning: use local clients info data")
        clients = fromLocal()
    }

    return clients
}
