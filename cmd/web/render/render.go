package render

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type TemplateData struct {
	ApiUrl string
	Data   map[string]interface{}
}

//go:embed templates
var templateFS embed.FS

// Execute and render the template with the writer
func RenderTemplate(w http.ResponseWriter, r *http.Request, pageName string, templateData *TemplateData, partials ...string) error {
	var t *template.Template

	templateName := fmt.Sprintf("%s/%s.page.gohtml", "templates", pageName)

	t, _ = parseTemplate(partials, pageName, templateName)

	if templateData == nil {
		templateData = &TemplateData{}
	}

	templateData = addDefaultData(templateData, r)
	if err := t.Execute(w, templateData); err != nil {
		return err
	}

	return nil
}

// Add global vars to templates
func addDefaultData(templateData *TemplateData, r *http.Request) *TemplateData {
	templateData.ApiUrl = os.Getenv("API_URL")

	return templateData
}

// Build partials for each template
func buildPartials(partials []string) []string {
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("%s/%s.partial.gohtml", "templates", x)
		}
	}

	return partials
}

// Parse each template
func parseTemplate(partials []string, page, templateName string) (*template.Template, error) {
	var t *template.Template

	partials = buildPartials(partials)
	templateFile := fmt.Sprintf("%s.page.gohtml", page)
	templateBaseLayout := fmt.Sprintf("%s/base.layout.gohtml", "templates")

	if len(partials) > 0 {
		t, _ = template.New(templateFile).ParseFS(templateFS, templateBaseLayout, strings.Join(partials, ","), templateName)
	} else {
		t, _ = template.New(templateFile).ParseFS(templateFS, templateBaseLayout, templateName)
	}

	return t, nil
}
