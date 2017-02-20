package backend

import (
	"fmt"
	"io"
	"os"
	"path"
	"text/template"

	"github.com/digitalrebar/digitalrebar/go/common/store"
)

// RenderTemplate is the result of rendering a BootEnv template
type RenderedTemplate struct {
	// Path is the absolute path that the Template will be rendered to.
	Path string
	// Template is the template that will rendered
	Template *Template
	// Vars holds the variables that will be used during template expansion.
	Vars *RenderData
}

func (r *RenderedTemplate) MkdirAll() error {
	return os.MkdirAll(path.Dir(r.Path), 0755)
}

func (r *RenderedTemplate) Write() error {
	tmplDest, err := os.Create(r.Path)
	if err != nil {
		return fmt.Errorf("Unable to create file %s: %v", r.Path, err)
	}
	defer tmplDest.Close()
	if err := r.Template.Render(tmplDest, r.Vars); err != nil {
		r.Remove()
		return fmt.Errorf("Error rendering template %s: %v", r.Template.Key(), err)
	}
	tmplDest.Sync()
	return nil
}

func (r *RenderedTemplate) Remove() error {
	return os.Remove(r.Path)
}

// Template represents a template that will be associated with a boot environment.
// swwagger:model
type Template struct {
	// ID is a unique identifier for this template.
	// required: true
	ID string
	// A description of this template
	Description string
	// Contents is the raw template.
	// required: true
	Contents   string
	parsedTmpl *template.Template
	p          *DataTracker
}

func (t *Template) Prefix() string {
	return "templates"
}

func (t *Template) Backend() store.SimpleStore {
	return t.p.getBackend(t)
}

func (t *Template) Key() string {
	return t.ID
}

func (t *Template) New() store.KeySaver {
	res := &Template{ID: t.ID, p: t.p}
	return store.KeySaver(res)
}

func (t *Template) List() []*Template {
	return AsTemplates(t.p.FetchAll(t))
}

// Parse checks to make sure the template contents are valid according to text/template.
func (t *Template) Parse() (err error) {
	parsedTmpl, err := template.New(t.ID).Parse(t.Contents)
	if err != nil {
		return err
	}
	t.parsedTmpl = parsedTmpl.Option("missingkey=error")
	return nil
}

func (t *Template) BeforeSave() error {
	e := &Error{Code: 422, Type: ValidationError, o: t}
	if t.ID == "" {
		e.Errorf("Template must have an ID")
	}
	if err := t.Parse(); err != nil {
		e.Errorf("Parse error: %v", err)
	}
	return e.OrNil()
}

func (t *Template) OnChange(oldThing store.KeySaver) error {
	e := &Error{Code: 422, Type: ValidationError, o: t}
	old := AsTemplate(oldThing)
	if old.ID != t.ID {
		e.Errorf("Cannot change ID of %s", t.ID)
	}
	return e.OrNil()
}

func (t *Template) BeforeDelete() error {
	var err error
	bootenv := t.p.NewBootEnv()
	bootEnvs := bootenv.List()
	for _, bootEnv := range bootEnvs {
		for _, tmpl := range bootEnv.Templates {
			if tmpl.ID == t.ID {
				return fmt.Errorf("template: %s is in use by bootenv %s (template %s", t.ID, bootEnv.Name, tmpl.Name)
			}
		}
	}
	return err
}

// Render executes the template with params writing the results to dest
func (t *Template) Render(dest io.Writer, params interface{}) error {
	if t.parsedTmpl == nil {
		if err := t.Parse(); err != nil {
			return fmt.Errorf("template: %s does not compile: %v", t.ID, err)
		}
	}
	if err := t.parsedTmpl.Execute(dest, params); err != nil {
		return fmt.Errorf("template: cannot execute %s: %v", t.ID, err)
	}
	return nil
}

func (p *DataTracker) NewTemplate() *Template {
	return &Template{p: p}
}

func AsTemplate(o store.KeySaver) *Template {
	return o.(*Template)
}

func AsTemplates(o []store.KeySaver) []*Template {
	res := make([]*Template, len(o))
	for i := range o {
		res[i] = AsTemplate(o[i])
	}
	return res
}