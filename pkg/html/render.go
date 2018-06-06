package html

import (
	"bytes"
	"html/template"

	"github.com/ssalvetti/driving-license-inspector/pkg/patenti"
)

var tabletemplate = `
<!DOCTYPE html>
<body>
<table>
	<tr>
		<th>Anno di Nascita</th>
		<th>Provincia</th>
		<th>Sesso</th>
		<th>Data Rilascio</th>
		<th>Punti</th>		 
	</tr>
	{{range .}}
	<tr>
		<td>{{.AnnoNascita}}</td>
		<td>{{.Provincia}}</td>
		<td>{{.Sesso}}</td>
		<td>{{.DataRilascio}}</td>
		<td>{{.PuntiPatente}}</td>
	</tr>
	{{end}}
</table>	
</body>
</html>
`

func RenderRecords(records []patenti.RecordPatente) ([]byte, error) {
	t, err := template.New("table").Parse(tabletemplate)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer(make([]byte, 0))
	err = t.Execute(w, records)

	return w.Bytes(), err
}
