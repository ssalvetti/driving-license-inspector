package html

import (
	"fmt"
	"log"
	"testing"

	"github.com/ssalvetti/driving-license-inspector/pkg/patenti"
	"github.com/stretchr/testify/assert"
)

func Test_StaticRendering(t *testing.T) {
	webpage, err := RenderRecords(nil)
	assert.Nil(t, err)
	assert.Contains(t, string(webpage), "<th>Anno di Nascita</th>")
	assert.Contains(t, string(webpage), "<th>Provincia</th>")
}

func Test_multipleRecordsAreRendered(t *testing.T) {
	var records = patenti.CreateTestRecords(2)
	webpage, err := RenderRecords(records)
	assert.Nil(t, err, fmt.Sprintf("error running template: %v", err))
	assert.Contains(t, string(webpage), "<td>1990</td>")
	assert.Contains(t, string(webpage), "<td>t</td>")
	assert.Contains(t, string(webpage), "<td>30</td>")
	log.Println(string(webpage))
}
