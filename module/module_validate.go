package module

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/yuansudong/protoc-gen-validate/templates"
)

const (
	validatorName     = "validate"
	langParam         = "lang"
	importPrefixParam = "import_prefix"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Validator() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return validatorName }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	lang := m.Parameters().Str(langParam)
	m.Assert(lang != "", "`lang` parameter must be set")
	tpls := templates.Template(m.Parameters())[lang]
	m.Assert(tpls != nil, "could not find templates for `lang`: ", lang)
	importPrefix := m.Parameters().Str(importPrefixParam)
	m.Assert(tpls != nil, "could not find templates for `import_prefix`: ", importPrefix)
	for _, f := range targets {
		for _, msg := range f.AllMessages() {
			m.CheckRules(msg)
		}
		for _, tpl := range tpls {
			out := templates.FilePathFor(importPrefix, tpl)(f, m.ctx, tpl)
			if out != nil {
				m.AddGeneratorTemplateFile(out.String(), tpl, f)
			}
		}
		m.Pop()
	}

	return m.Artifacts()
}

var _ pgs.Module = (*Module)(nil)
