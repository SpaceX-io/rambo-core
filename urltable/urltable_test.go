package urltable

import (
	"net/http"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	pattern, err := Format(" view / a / b / c ")
	if err != nil {
		t.Fatal(err)
	}

	if pattern != methodView+"/a/b/c" {
		t.Fatal("format failed")
	}

	pattern, err = Format("get/a/*")
	if err != nil {
		t.Fatal(err)
	}

	if pattern != http.MethodGet+"/a/"+fuzzy {
		t.Fatal("format failed")
	}
}

func TestParse(t *testing.T) {
	for i, pattern := range []string{
		"get/ a / b / c ",
		"get/ a / b / c / ** ",
		"post/ a / b / * / c / ** ",
		"delete/ a / b / * / c / **",
		"put/ a / b / * / c / ** ",
	} {
		paths, err := parse(pattern)
		if err != nil {
			t.Fatal(pattern, "should be illegal err: ", err.Error())
		}
		t.Log(i, strings.Join(paths, delimiter))
	}

	for _, pattern := range []string{
		" ",
		" / ",
		" x / ",
		"get/ ",
		"get/ * ",
		"get/ ** ",
		"get/ * / ** ",
		"post/ * / ** ",
		"delete/ * / a / **",
	} {
		if _, err := parse(pattern); err != nil {
			t.Fatal(pattern, "should be illegal")
		}
	}
}
