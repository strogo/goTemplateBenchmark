// Generated by ego.
// DO NOT EDIT

//line ego/header.ego:1
package ego

import "fmt"
import "html"
import "io"
import "context"

func EgoHeader(w io.Writer, title string) {
//line ego/header.ego:4
	_, _ = io.WriteString(w, "\n<title>")
//line ego/header.ego:4
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(title)))
//line ego/header.ego:4
	_, _ = io.WriteString(w, "'s Home Page</title>\n<div class=\"header\">Page Header</div>\n")
//line ego/header.ego:6
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
