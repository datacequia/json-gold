package ld_test

import (
	"fmt"
	"testing"

	"path/filepath"

	"github.com/piprate/json-gold/ld"
)

func TestProcessorListener(t *testing.T) {

	const testDir = "testdata"

	processor := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	tl := &testListener{
		t: t,
	}

	options.Listener = tl

	jsonld := filepath.Join("testdata", "schema.org", "docs", "jsonldcontext.json")

	dl := ld.NewDefaultDocumentLoader(nil)

	rd, err1 := dl.LoadDocument(jsonld)
	if err1 != nil {
		t.Fatal("LoadDocument failed", err1)
	}
	//ld.PrintDocument("defaultdocloader", dl)

	_, err := processor.Expand(rd.Document, options)
	if err != nil {
		t.Fatal("ld.Expand failed:", err)
		return
	}
	//ld.PrintDocument("testdoc", x)

	if tl.count < 1 {
		t.Fatalf("expected OnTermDefinition count to be > 0, found %d", tl.count)
	}

}

type testListener struct {
	t     *testing.T
	count int
}

func (tl *testListener) OnTermDefinition(termDef ld.TermDefinition) {

	switch termDef.Term() {
	case "schema":
		if termDef.Value().IdKeyword != "http://schema.org/" {
			tl.t.FailNow()
		}
		if !termDef.Value().IsPrefix {
			tl.t.FailNow()
		}
	case "rdfs":
		if termDef.Value().IdKeyword != "http://www.w3.org/2000/01/rdf-schema#" {
			tl.t.FailNow()
		}
		if !termDef.Value().IsPrefix {
			tl.t.FailNow()
		}
		//"AboutPage": {"@id": "schema:AboutPage"}
	case "AboutPage":
		if termDef.Value().IdKeyword != "http://schema.org/AboutPage" {
			fmt.Println("AboutPage", termDef.Value().IdKeyword)
			tl.t.FailNow()
		}
	//"applicationCategory": { "@id": "schema:applicationCategory", "@type": "@id"},
	case "applicationCategory":
		if termDef.Value().IdKeyword != "http://schema.org/applicationCategory" {
			fmt.Println("AboutPage", termDef.Value().IdKeyword)
			tl.t.FailNow()
		}
		if termDef.Value().TypeKeyword != "@id" {
			fmt.Println("@type: ", termDef.Value().TypeKeyword)
			tl.t.FailNow()
		}

	}
	//fmt.Println("new term:", termDef.Term(), termDef.Value())

	tl.count++
}
