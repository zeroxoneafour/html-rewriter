package htmlrewriter

import (
    "testing"
    "fmt"
)

// html from my website hotcorn.ga
const testHtml = `<!DOCTYPE html>
<html>
        <head>
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <link rel="stylesheet" href="/css/index.css">
                <link rel="icon" type="image/png" href="/img/cornicon.png">
                <title>Corn Unblocked</title>
        </head>
        <body>
                <h1>Welcome to Corn Unblocked!</h1>
                <img class="mainimg" src="/img/corn.svg" alt="Corn Unblocked" class="center">
                <p>This site features unblocked games and a pr0xy!</p>
                <a class="button" href="/games/index.html">Games</a>
                <a class="button" href="/iframe/index.html">IFrame HTML Viewer</a>
                <a class="button" href="/pr0xy/index.html">Pr0xy</a>
                <a class="button" href="/other/index.html">More Games</a>
                <a class="button" href="/about/index.html">About</a>
                <p class="bottom">Licensed under the <a href="license.txt">MIT License</a>. Image under the <a href="https://creativecommons.org/licenses/by/4.0/legalcode">CC BY 4.0</a>.</p>
        </body>
</html>
`

func TestHTMLFix(t *testing.T) {
    fmt.Print(RewriteHTML(testHtml, "https://hotcorn.ga/"))
}
