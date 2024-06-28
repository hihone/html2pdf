package pdfGenerator

import (
	"fmt"
	"testing"
)

func TestParseHtml(t *testing.T) {
	tpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.title}}</title>
</head>
<body>
    <div>Name: {{.name}}</div>
    <div>Age: {{.age}}</div>
    {{if .items}}
    <ul>
        {{range $item := .items}}
        <li>{{$item.title}}</li>
        {{end}}
    </ul>
    {{end}}
    <div>{{showHtml .html}}</div>
</body>
</html>`
	r := NewRequestPDF("a.pdf")
	err := r.ParseTemplate(tpl, map[string]any{
		"name":  "YYang",
		"age":   30,
		"title": "模板测试",
		"items": []map[string]any{
			{
				"title": "列表 1",
			},
			{
				"title": "列表 2",
			},
			{
				"title": "列表 3",
			},
			{
				"title": "列表 4",
			},
		},
		"html": "<b style=\"color: red;\"> 加粗 </b>",
	})
	if err != nil {
		t.Error(err)
		return
	}
	b, err := r.Build()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%t\n", b)
}
func TestParseHtmlFile(t *testing.T) {
	r := NewRequestPDF("b.pdf")
	err := r.ParseTemplateFile("../tpl/tpl.html", map[string]any{
		"name":  "YYang",
		"age":   30,
		"title": "模板测试",
		"items": []map[string]any{
			{
				"title": "列表 1",
			},
			{
				"title": "列表 2",
			},
			{
				"title": "列表 3",
			},
			{
				"title": "列表 4",
			},
		},
		"html": "<b> 加粗 </b>",
	})
	if err != nil {
		t.Error(err)
		return
	}
	b, err := r.Build()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%t\n", b)
}
