// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateDino() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html data-theme=\"lofi\"><head><link href=\"https://cdn.jsdelivr.net/npm/daisyui@4.11.1/dist/full.min.css\" rel=\"stylesheet\" type=\"text/css\"><script src=\"https://cdn.tailwindcss.com\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12\" integrity=\"sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2\" crossorigin=\"anonymous\"></script></head><main class=\"p-5\"><div class=\"p-5\"><div class=\"navbar bg-primary flex justify-between text-white rounded-lg\"><a href=\"/\" class=\"btn btn-ghost text-xl\">Gophosaurus</a></div><div id=\"create_form\" class=\"card\"><form class=\"card-body\"><p class=\"card-title\">Create a new Dino</p><label class=\"input input-bordered flex items-center gap-2\">Dino Name <input type=\"text\" name=\"name\" class=\"grow\" placeholder=\"Dondino\"></label> <label class=\"input input-bordered flex items-center gap-2\">Dino small description <input type=\"text\" name=\"description\" class=\"grow\" placeholder=\"Dondino is cool but you are cooler\"></label></form><button type=\"button\" class=\"btn btn-primary\">Create Dino</button></div></div></main></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
