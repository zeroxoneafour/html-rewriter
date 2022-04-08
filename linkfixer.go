package htmlrewriter

import (
    "net/url"
    "strings"
)

func fixLink(linkStr, referenceUrlStr string) string {
    link, _ := url.Parse(linkStr)
    referenceUrl, _ := url.Parse(referenceUrlStr)

    if len(link.Path) == 0 { // root url
        return "/" + linkStr
    } else if link.Path[0] != '/' { // relative paths
        lastSlash := strings.LastIndex(referenceUrl.Path, "/")
        fixedPath := referenceUrl.Path[:lastSlash] + link.Path
        return "/" + referenceUrl.Scheme + "://" + referenceUrl.Host + "/" + fixedPath
    } else if len(link.Host) == 0 { // absolute paths
        return "/" + referenceUrl.Scheme + "://" + referenceUrl.Host + link.Path
    } else { // full urls
        return "/" + linkStr
    }
}
