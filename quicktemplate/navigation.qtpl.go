// This file is automatically generated by qtc from "navigation.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line quicktemplate/navigation.qtpl:1
package quicktemplate

//line quicktemplate/navigation.qtpl:1
import (
	"io"

	"github.com/valyala/quicktemplate"
)

//line quicktemplate/navigation.qtpl:1
import "github.com/SlinSo/goTemplateBenchmark/model"

//line quicktemplate/navigation.qtpl:2
var (
	_ = io.Copy
	_ = quicktemplate.AcquireByteBuffer
)

//line quicktemplate/navigation.qtpl:2
func StreamNavigation(qw *quicktemplate.Writer, nav []*model.Navigation) {
	//line quicktemplate/navigation.qtpl:2
	qw.N().S(`
<ul class="navigation">
`)
	//line quicktemplate/navigation.qtpl:4
	for _, item := range nav {
		//line quicktemplate/navigation.qtpl:4
		qw.N().S(`
	<li><a href="`)
		//line quicktemplate/navigation.qtpl:5
		qw.E().S(item.Link)
		//line quicktemplate/navigation.qtpl:5
		qw.N().S(`">`)
		//line quicktemplate/navigation.qtpl:5
		qw.E().S(item.Item)
		//line quicktemplate/navigation.qtpl:5
		qw.N().S(`</a></li>
`)
		//line quicktemplate/navigation.qtpl:6
	}
	//line quicktemplate/navigation.qtpl:6
	qw.N().S(`
</ul>
`)
//line quicktemplate/navigation.qtpl:8
}

//line quicktemplate/navigation.qtpl:8
func WriteNavigation(qww io.Writer, nav []*model.Navigation) {
	//line quicktemplate/navigation.qtpl:8
	qw := quicktemplate.AcquireWriter(qww)
	//line quicktemplate/navigation.qtpl:8
	StreamNavigation(qw, nav)
	//line quicktemplate/navigation.qtpl:8
	quicktemplate.ReleaseWriter(qw)
//line quicktemplate/navigation.qtpl:8
}

//line quicktemplate/navigation.qtpl:8
func Navigation(nav []*model.Navigation) string {
	//line quicktemplate/navigation.qtpl:8
	qb := quicktemplate.AcquireByteBuffer()
	//line quicktemplate/navigation.qtpl:8
	WriteNavigation(qb, nav)
	//line quicktemplate/navigation.qtpl:8
	qs := string(qb.B)
	//line quicktemplate/navigation.qtpl:8
	quicktemplate.ReleaseByteBuffer(qb)
	//line quicktemplate/navigation.qtpl:8
	return qs
//line quicktemplate/navigation.qtpl:8
}