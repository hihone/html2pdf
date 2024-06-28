package pdfGenerator

import (
	"bytes"
	"fmt"
	wkPdf "github.com/hihone/go-wkhtmltopdf"
	werrors "github.com/pkg/errors"
	"html/template"
	"os"
	"strings"
	"time"
)

const (
	tplTypeParseTplContent = "TplContent"
	tplTypeParseTplFile    = "TplFile"
)

type RequestPDF struct {
	html    string
	workDir string
	pdfPath string
	tplType string
}

func NewRequestPDF(pdfPath string) *RequestPDF {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &RequestPDF{
		pdfPath: pdfPath,
		workDir: workDir,
	}
}

func (r *RequestPDF) ParseTemplate(htmlTplContent string, data map[string]any) error {
	r.tplType = tplTypeParseTplContent
	t, err := template.New(fmt.Sprintf("%d", time.Now().UnixMicro())).
		Funcs(template.FuncMap{
			"showHtml": func(text string) (template.HTML, error) {
				return template.HTML(text), nil
			},
		}).
		Delims("{{", "}}").
		Parse(htmlTplContent)
	if err != nil {
		return werrors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return werrors.WithStack(err)
	}

	r.html = buf.String()

	return nil
}

func (r *RequestPDF) ParseTemplateFile(tplFile string, data map[string]any) error {
	r.tplType = tplTypeParseTplFile
	tplPath := strings.Split(tplFile, "/")
	t, err := template.New(tplPath[len(tplPath)-1]).
		Funcs(template.FuncMap{
			"showHtml": func(text string) (template.HTML, error) {
				return template.HTML(text), nil
			},
		}).
		Delims("{{", "}}").
		ParseFiles(tplFile)
	if err != nil {
		return werrors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return werrors.WithStack(err)
	}

	r.html = buf.String()

	return nil
}

func (r *RequestPDF) Build() (bool, error) {
	pdf, err := wkPdf.NewPDFGenerator()
	if err != nil {
		return false, werrors.WithStack(err)
	}
	switch r.tplType {
	case tplTypeParseTplContent, tplTypeParseTplFile:
		pdf.AddPage(wkPdf.NewPageReader(strings.NewReader(r.html)))
	}

	if err := pdf.Create(); err != nil {
		return false, werrors.WithStack(err)
	}
	if err := pdf.WriteFile(r.pdfPath); err != nil {
		return false, werrors.WithStack(err)
	}

	return true, nil
}
