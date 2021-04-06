package zlog

import (
	"github.com/zytekaron/zgo/zlog"
	"testing"
)

func TestFormat(t *testing.T) {
	// Ensure the formatter does not get stuck anywhere (check this manually)
	format1 := "{{}{}}{}{}}}{{}{}{{{}}{{}}}{}{{{{{}}}}}}{}{}}{}{}}}}}"
	res1 := zlog.Format(format1, 1, 2, 4, 5, 6, 7, 8)
	if res1 != "{}1}2{}}{}4{{}{}}5{{{}}}6{}7{}}} 8" {
		t.Error("invalid result string:", res1)
	}

	// Ensure the formatter supports proper escaping
	format2 := "{{ {} }} {{{{}}}}"
	res2 := zlog.Format(format2, 123, 456)
	if res2 != "{ 123 } {{}} 456" {
		t.Error("invalid result string:", res2)
	}
}
