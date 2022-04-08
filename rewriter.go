package htmlrewriter

import (
    "golang.org/x/net/html"
    "strings"
    "bytes"
)

func RewriteHTML(htmlString, url string) string { // rewrite HTML. JavaScript handled separately
    doc, _ := html.Parse(strings.NewReader(htmlString))
    var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
                    newUrl := fixLink(a.Val, url)
                    n.Attr[i].Val = newUrl
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
    buf := new(bytes.Buffer)
    html.Render(buf, doc)
    return buf.String()
}
