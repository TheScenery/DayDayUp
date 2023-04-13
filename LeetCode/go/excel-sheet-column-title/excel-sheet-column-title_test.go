package excelsheetcolumntitle

import "testing"

func TestConvertToTitle(t *testing.T) {
	ret := convertToTitle(1)
	if ret != "A" {
		t.Fatal(ret)
	}
}
