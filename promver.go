// +build go1.12

package promver

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime/debug"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Build depedency information. Populated at build-time.
var (
	buildInfo, ok = debug.ReadBuildInfo()
)

// NewCollector returns a collector which exports metrics about current dependency information.
func NewCollector(program string) *prometheus.GaugeVec {
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: program,
			Name:      "go_mod_info",
			Help: fmt.Sprintf(
				"A metric with a constant '1' value labeled by dependency name, version, from which %s was built.",
				program,
			),
		},
		[]string{"name", "version"},
	)
	if !ok {
		return gauge
	}

	for _, dep := range buildInfo.Deps {
		d := dep
		if dep.Replace != nil {
			d = dep.Replace
		}
		gauge.WithLabelValues(d.Path, d.Version).Set(1)
	}
	return gauge
}

// versionInfoTmpl contains the template used by Info.
var versionInfoTmpl = `
{{.Program}}
{{range $k,$v := .Deps}} {{$k}}: {{$v}}
{{end}}`

type versionPrint struct {
	Program string
	Deps    map[string]string
}

// Print returns module version information.
func Print(program string) string {
	m := make(map[string]string)
	if ok {
		for _, dep := range buildInfo.Deps {
			d := dep
			if dep.Replace != nil {
				d = dep.Replace
			}
			m[d.Path] = d.Version
		}
	}

	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", versionPrint{
		Program: program,
		Deps:    m,
	}); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// Info returns dependency versions
func Info() string {
	var info []string
	for _, dep := range buildInfo.Deps {
		d := dep
		if dep.Replace != nil {
			d = dep.Replace
		}
		info = append(info, d.Path+": ", d.Version)
	}

	return fmt.Sprintf("(%s)", strings.Join(info, ","))
}
