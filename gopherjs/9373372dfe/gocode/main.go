package main

import (
	"cloud.google.com/go/compute/metadata"
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("gocode", map[string]interface{}{
		"setElement": setElement,
	})
}

func setElement()  {
	d := dom.GetWindow().Document()
	el := d.GetElementByID("topelement")
	go func() {
		if metadata.OnGCE() {
			el.SetInnerHTML("On GCE")
			return 
		}
		el.SetInnerHTML("Not on GCE")
	}()
}
