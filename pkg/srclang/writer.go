package srclang

import (
	"fmt"
	"io"
	"strings"
)

const (
	xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>`
	namespace = "https://srclang.dev/ns/core/0"
)

func WriteDocument(w io.Writer, doc *Document) error {
	b := &xmlBuilder{w: w}
	b.line(xmlHeader)
	b.openf(`<srclang version="%s" xmlns="%s">`, doc.Version, namespace)
	writeHead(b, &doc.Head)
	writeBody(b, &doc.Body)
	b.close("srclang")
	return b.err
}

func writeHead(b *xmlBuilder, h *Head) {
	b.open("head")
	if h.Producer != "" {
		b.element("producer", h.Producer)
	}
	if h.Repository != nil {
		b.openf(`<repository uri="%s">`, escAttr(h.Repository.URI))
		if h.Repository.Commit != "" {
			b.linef(`<commit sha="%s"/>`, escAttr(h.Repository.Commit))
		}
		if h.Repository.Branch != "" {
			b.element("branch", h.Repository.Branch)
		}
		b.close("repository")
	}
	if h.Component != "" {
		b.element("component", h.Component)
	}
	if h.Extracted != "" {
		b.element("extracted", h.Extracted)
	}
	b.linef(`<layer name="%s"/>`, escAttr(h.Layer))
	if len(h.Languages) > 0 {
		b.open("languages")
		for _, lang := range h.Languages {
			if lang.Version != "" {
				b.linef(`<language name="%s" version="%s"/>`, escAttr(lang.Name), escAttr(lang.Version))
			} else {
				b.linef(`<language name="%s"/>`, escAttr(lang.Name))
			}
		}
		b.close("languages")
	}
	if len(h.Diagnostics) > 0 {
		b.open("diagnostics")
		for _, d := range h.Diagnostics {
			if d.File != "" {
				b.linef(`<warning file="%s">%s</warning>`, escAttr(d.File), escText(d.Message))
			} else {
				b.element("warning", d.Message)
			}
		}
		b.close("diagnostics")
	}
	b.close("head")
}

func writeBody(b *xmlBuilder, body *Body) {
	b.open("body")
	writeLayer(b, &body.Layer)
	b.close("body")
}

func writeLayer(b *xmlBuilder, l *Layer) {
	b.openf(`<layer name="%s">`, escAttr(l.Name))
	if l.Summary != "" {
		b.element("summary", l.Summary)
	}
	for i := range l.Findings {
		writeFinding(b, &l.Findings[i])
	}
	for i := range l.Files {
		writeFile(b, &l.Files[i])
	}
	for i := range l.Types {
		writeType(b, &l.Types[i], "")
	}
	for i := range l.Resources {
		writeResource(b, &l.Resources[i])
	}
	for i := range l.Relationships {
		writeRelationship(b, &l.Relationships[i])
	}
	for i := range l.Configs {
		writeConfig(b, &l.Configs[i])
	}
	for i := range l.Imports {
		writeImport(b, &l.Imports[i])
	}
	b.close("layer")
}

func writeFile(b *xmlBuilder, f *File) {
	attrs := fmt.Sprintf(`path="%s"`, escAttr(f.Path))
	if f.Language != "" {
		attrs += fmt.Sprintf(` language="%s"`, escAttr(f.Language))
	}
	if f.Lines > 0 {
		attrs += fmt.Sprintf(` lines="%d"`, f.Lines)
	}
	if f.ParseError {
		attrs += ` parse-error="true"`
	}
	b.openf("<file %s>", attrs)
	if f.Summary != "" {
		b.element("summary", f.Summary)
	}
	for i := range f.Functions {
		writeFunction(b, &f.Functions[i], f.Path)
	}
	for i := range f.Types {
		writeType(b, &f.Types[i], f.Path)
	}
	for i := range f.Imports {
		writeImport(b, &f.Imports[i])
	}
	b.close("file")
}

func writeFunction(b *xmlBuilder, fn *Function, parentFile string) {
	attrs := fmt.Sprintf(`name="%s"`, escAttr(fn.Name))
	if fn.Kind != "" {
		attrs += fmt.Sprintf(` kind="%s"`, escAttr(fn.Kind))
	}
	if fn.Trust != "" {
		attrs += fmt.Sprintf(` trust="%s"`, escAttr(fn.Trust))
	}
	if fn.TaintRole != "" {
		attrs += fmt.Sprintf(` taint-role="%s"`, escAttr(fn.TaintRole))
	}
	if fn.AuthRequired != "" {
		attrs += fmt.Sprintf(` auth-required="%s"`, escAttr(fn.AuthRequired))
	}
	b.openf("<function %s>", attrs)

	if parentFile != "" && fn.SourceLine > 0 {
		b.linef(`<source line="%d"/>`, fn.SourceLine)
	} else if fn.SourceFile != "" {
		b.linef(`<source file="%s" line="%d"/>`, escAttr(fn.SourceFile), fn.SourceLine)
	}

	if fn.ReceiverType != "" {
		if fn.ReceiverName != "" {
			b.linef(`<receiver name="%s" type="%s"/>`, escAttr(fn.ReceiverName), escAttr(fn.ReceiverType))
		} else {
			b.linef(`<receiver type="%s"/>`, escAttr(fn.ReceiverType))
		}
	}

	if len(fn.Params) > 0 {
		b.open("params")
		for _, p := range fn.Params {
			if p.Type != "" {
				b.linef(`<param name="%s" type="%s"/>`, escAttr(p.Name), escAttr(p.Type))
			} else {
				b.linef(`<param name="%s"/>`, escAttr(p.Name))
			}
		}
		b.close("params")
	}

	if len(fn.Returns) > 0 {
		b.open("returns")
		for _, r := range fn.Returns {
			b.linef(`<return type="%s"/>`, escAttr(r.Type))
		}
		b.close("returns")
	}

	if fn.Complexity > 0 || fn.BodyLines > 0 {
		attrs := ""
		if fn.Complexity > 0 {
			attrs += fmt.Sprintf(` complexity="%d"`, fn.Complexity)
		}
		if fn.BodyLines > 0 {
			attrs += fmt.Sprintf(` lines="%d"`, fn.BodyLines)
		}
		b.linef("<metrics%s/>", attrs)
	}

	if fn.Code != "" {
		b.linef("<code><![CDATA[%s]]></code>", escapeCDATA(fn.Code))
	}

	for _, m := range fn.Metas {
		b.linef(`<meta domain="%s" key="%s" value="%s"/>`, escAttr(m.Domain), escAttr(m.Key), escAttr(m.Value))
	}

	b.close("function")
}

func writeType(b *xmlBuilder, t *Type, parentFile string) {
	attrs := fmt.Sprintf(`name="%s"`, escAttr(t.Name))
	if t.Kind != "" {
		attrs += fmt.Sprintf(` kind="%s"`, escAttr(t.Kind))
	}
	b.openf("<type %s>", attrs)

	if parentFile != "" && t.SourceLine > 0 {
		b.linef(`<source line="%d"/>`, t.SourceLine)
	} else if t.SourceFile != "" {
		b.linef(`<source file="%s" line="%d"/>`, escAttr(t.SourceFile), t.SourceLine)
	}

	if len(t.Fields) > 0 {
		b.open("fields")
		for _, f := range t.Fields {
			attrs := fmt.Sprintf(`name="%s"`, escAttr(f.Name))
			if f.Type != "" {
				attrs += fmt.Sprintf(` type="%s"`, escAttr(f.Type))
			}
			if f.Visibility != "" {
				attrs += fmt.Sprintf(` visibility="%s"`, escAttr(f.Visibility))
			}
			b.linef("<field %s/>", attrs)
		}
		b.close("fields")
	}

	if len(t.Implements) > 0 {
		b.open("implements")
		for _, iface := range t.Implements {
			b.linef(`<interface name="%s"/>`, escAttr(iface))
		}
		b.close("implements")
	}

	if t.Summary != "" {
		b.element("summary", t.Summary)
	}

	b.close("type")
}

func writeResource(b *xmlBuilder, r *Resource) {
	attrs := fmt.Sprintf(`kind="%s" name="%s"`, escAttr(r.Kind), escAttr(r.Name))
	b.openf("<resource %s>", attrs)
	if r.SourceFile != "" {
		b.linef(`<source file="%s" line="%d"/>`, escAttr(r.SourceFile), r.SourceLine)
	}
	if r.APIGroup != "" {
		b.element("api-group", r.APIGroup)
	}
	if r.APIVersion != "" {
		b.element("api-version", r.APIVersion)
	}
	if r.Scope != "" {
		b.element("scope", r.Scope)
	}
	if r.Origin != "" {
		b.element("origin", r.Origin)
	}
	if r.FieldCount > 0 {
		b.linef("<field-count>%d</field-count>", r.FieldCount)
	}
	if r.Summary != "" {
		b.element("summary", r.Summary)
	}
	for _, c := range r.Children {
		b.raw(c.XMLContent)
	}
	b.close("resource")
}

func writeRelationship(b *xmlBuilder, r *Relationship) {
	b.openf(`<relationship kind="%s">`, escAttr(r.Kind))
	writeEndpoint(b, "from", &r.From)
	writeEndpoint(b, "to", &r.To)
	if r.Confidence > 0 {
		b.linef(`<confidence value="%.2f"/>`, r.Confidence)
	}
	b.close("relationship")
}

func writeEndpoint(b *xmlBuilder, tag string, e *Endpoint) {
	attrs := ""
	if e.Function != "" {
		attrs += fmt.Sprintf(` function="%s"`, escAttr(e.Function))
	}
	if e.Resource != "" {
		attrs += fmt.Sprintf(` resource="%s"`, escAttr(e.Resource))
	}
	if e.TypeName != "" {
		attrs += fmt.Sprintf(` type="%s"`, escAttr(e.TypeName))
	}
	if e.File != "" {
		attrs += fmt.Sprintf(` file="%s"`, escAttr(e.File))
	}
	if e.Line > 0 {
		attrs += fmt.Sprintf(` line="%d"`, e.Line)
	}
	if e.TaintRole != "" {
		attrs += fmt.Sprintf(` taint-role="%s"`, escAttr(e.TaintRole))
	}
	if e.APIGroup != "" {
		attrs += fmt.Sprintf(` api-group="%s"`, escAttr(e.APIGroup))
	}
	if e.Resolved != nil && !*e.Resolved {
		attrs += ` resolved="false"`
	}
	b.linef("<%s%s/>", tag, attrs)
}

func writeFinding(b *xmlBuilder, f *Finding) {
	b.openf(`<finding id="%s" domain="%s" severity="%s" rule="%s">`,
		escAttr(f.ID), escAttr(f.Domain), escAttr(f.Severity), escAttr(f.Rule))
	if f.SourceFile != "" {
		b.linef(`<source file="%s" line="%d"/>`, escAttr(f.SourceFile), f.SourceLine)
	}
	if f.Title != "" {
		b.element("title", f.Title)
	}
	if f.Description != "" {
		b.element("description", f.Description)
	}
	if len(f.Evidence) > 0 {
		b.open("evidence")
		for _, r := range f.Evidence {
			attrs := fmt.Sprintf(`type="%s" name="%s"`, escAttr(r.Type), escAttr(r.Name))
			if r.File != "" {
				attrs += fmt.Sprintf(` file="%s"`, escAttr(r.File))
			}
			if r.Line > 0 {
				attrs += fmt.Sprintf(` line="%d"`, r.Line)
			}
			b.linef("<ref %s/>", attrs)
		}
		b.close("evidence")
	}
	b.close("finding")
}

func writeConfig(b *xmlBuilder, c *Config) {
	b.openf(`<config kind="%s" path="%s">`, escAttr(c.Kind), escAttr(c.Path))
	for _, ch := range c.Children {
		b.raw(ch.XMLContent)
	}
	b.close("config")
}

func writeImport(b *xmlBuilder, imp *Import) {
	attrs := fmt.Sprintf(`module="%s"`, escAttr(imp.Module))
	if imp.Kind != "" {
		attrs += fmt.Sprintf(` kind="%s"`, escAttr(imp.Kind))
	}
	if imp.Path != "" {
		attrs += fmt.Sprintf(` path="%s"`, escAttr(imp.Path))
	}
	if imp.Version != "" {
		attrs += fmt.Sprintf(` version="%s"`, escAttr(imp.Version))
	}
	b.linef("<import %s/>", attrs)
}

func escapeCDATA(s string) string {
	return strings.ReplaceAll(s, "]]>", "]]]]><![CDATA[>")
}

func escAttr(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, `"`, "&quot;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\n", "&#10;")
	s = strings.ReplaceAll(s, "\r", "&#13;")
	s = strings.ReplaceAll(s, "\t", "&#9;")
	return s
}

func escText(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

type xmlBuilder struct {
	w      io.Writer
	err    error
	indent int
}

func (b *xmlBuilder) line(s string) {
	if b.err != nil {
		return
	}
	_, b.err = fmt.Fprintf(b.w, "%s%s\n", strings.Repeat("  ", b.indent), s)
}

func (b *xmlBuilder) linef(format string, args ...interface{}) {
	b.line(fmt.Sprintf(format, args...))
}

func (b *xmlBuilder) open(tag string) {
	b.linef("<%s>", tag)
	b.indent++
}

func (b *xmlBuilder) openf(format string, args ...interface{}) {
	b.line(fmt.Sprintf(format, args...))
	b.indent++
}

func (b *xmlBuilder) close(tag string) {
	b.indent--
	b.linef("</%s>", tag)
}

func (b *xmlBuilder) element(tag, content string) {
	b.linef("<%s>%s</%s>", tag, escText(content), tag)
}

func (b *xmlBuilder) raw(content string) {
	if b.err != nil || content == "" {
		return
	}
	for _, line := range strings.Split(content, "\n") {
		b.line(line)
	}
}
