package obfs

import (
    "log"
    utls "github.com/refraction-networking/utls"
)

func EnableUTLS(profile string) {
    var id utls.ClientHelloID
    switch profile {
    case "chrome_124": id = utls.HelloChrome_124
    case "firefox_125": id = utls.HelloFirefox_125
    default: id = utls.HelloRandomized
    }
    log.Printf("UTLS fingerprint set to %s", id.Str())
}
