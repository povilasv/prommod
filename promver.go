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
	BuidInfo, ok = debug.ReadBuildInfo()
)

// NewCollector returns a collector which exports metrics about current dependency information.
func NewCollector(program string) *prometheus.GaugeVec {
	buildInfo := prometheus.NewGaugeVec(
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
		return buildInfo
	}

	for _, dep := range BuidInfo.Deps {
		d := dep
		if dep.Replace != nil {
			d = dep.Replace
		}
		buildInfo.WithLabelValues(d.Path, d.Version).Set(1)
	}
	return buildInfo
}

// versionInfoTmpl contains the template used by Info.
var versionInfoTmpl = `
{{.program}}
{{range $k,$v := .} $k: $v\n{{end}`

// Print returns version information.
func Print(program string) string {
	m := map[string]string{
		"program": program,
	}
	for _, dep := range BuidInfo.Deps {
		d := dep
		if dep.Replace != nil {
			d = dep.Replace
		}
		m[d.Path] = d.Version
	}

	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// Info returns dependency versions
func Info() string {
	var info []string
	for _, dep := range BuidInfo.Deps {
		d := dep
		if dep.Replace != nil {
			d = dep.Replace
		}
		info = append(info, d.Path+": ", d.Version)
	}

	return fmt.Sprintf("(%s)", strings.Join(info, ","))
}
