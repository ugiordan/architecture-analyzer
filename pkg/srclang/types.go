package srclang

type Document struct {
	Version string
	Head    Head
	Body    Body
}

type Head struct {
	Producer    string
	Repository  *Repository
	Component   string
	Extracted   string
	Layer       string
	Languages   []Language
	Diagnostics []Warning
}

type Repository struct {
	URI    string
	Commit string
	Branch string
}

type Language struct {
	Name    string
	Version string
}

type Warning struct {
	File    string
	Message string
}

type Body struct {
	Layer Layer
}

type Layer struct {
	Name          string
	Summary       string
	Files         []File
	Types         []Type
	Resources     []Resource
	Relationships []Relationship
	Findings      []Finding
	Configs       []Config
	Imports       []Import
}

type File struct {
	Path       string
	Language   string
	Lines      int
	Summary    string
	ParseError bool
	Functions  []Function
	Types      []Type
	Imports    []Import
}

type Function struct {
	Name         string
	Kind         string
	SourceFile   string
	SourceLine   int
	ReceiverName string
	ReceiverType string
	Params       []Param
	Returns      []Return
	Complexity   int
	BodyLines    int
	Code         string
	Trust        string
	TaintRole    string
	AuthRequired string
	Metas        []Meta
}

type Param struct {
	Name string
	Type string
}

type Return struct {
	Type string
}

type Type struct {
	Name       string
	Kind       string
	SourceFile string
	SourceLine int
	Fields     []Field
	Implements []string
	Summary    string
}

type Field struct {
	Name       string
	Type       string
	Visibility string
}

type Resource struct {
	Kind       string
	Name       string
	SourceFile string
	SourceLine int
	APIGroup   string
	APIVersion string
	Scope      string
	Origin     string
	FieldCount int
	Summary    string
	Children   []ResourceChild
}

type ResourceChild struct {
	XMLContent string
}

type Relationship struct {
	Kind       string
	From       Endpoint
	To         Endpoint
	Confidence float64
}

type Endpoint struct {
	Function  string
	Resource  string
	TypeName  string
	File      string
	Line      int
	TaintRole string
	APIGroup  string
	Resolved  *bool
}

type Finding struct {
	ID          string
	Domain      string
	Severity    string
	Rule        string
	SourceFile  string
	SourceLine  int
	Title       string
	Description string
	Evidence    []Ref
}

type Ref struct {
	Type string
	Name string
	File string
	Line int
}

type Config struct {
	Kind     string
	Path     string
	Children []ConfigChild
}

type ConfigChild struct {
	XMLContent string
}

type Import struct {
	Module  string
	Kind    string
	Path    string
	Version string
}

type Meta struct {
	Domain string
	Key    string
	Value  string
}
