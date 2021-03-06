package template

import "github.com/deadcheat/goblet"

// Assets struct for assign values to template
type Assets struct {
	ExecutedCommand    string
	PackageName        string
	VarName            string
	GenerateGoGenerate bool
	DirMap             map[string][]string
	FileMap            map[string]*goblet.File
	Paths              []string
	HasFileData        bool
}

// AssetFileTemplate template for asset file
var AssetFileTemplate = `package {{.PackageName}}

import(
	"time"

	"github.com/deadcheat/goblet"
)

{{ if .GenerateGoGenerate }}//go:generate {{ .ExecutedCommand }}{{ end }}

{{ $FileMap := .FileMap}}{{ $DirMap := .DirMap}}{{ $VarName := .VarName }}
// {{ $VarName }} a generated file system
var {{ $VarName }} = goblet.NewFS(
	map[string][]string{
		{{- range $p := .Paths }}{{ with (index $DirMap $p)}}
		"{{ $p }}": []string{
			{{ range $s := . }}"{{ $s }}", {{ end }}
		},{{ end }}
		{{- end }}
	},
	map[string]*goblet.File {
		{{- range $p := .Paths }}{{ with (index $FileMap $p)}}
		"{{$p}}": goblet.NewFile("{{$p}}", {{if .IsDir }}nil{{ else if eq (len .Data) 0 }}[]byte{}{{ else }}_{{ $VarName }}{{ sha1 $p }}{{ end }}, {{ printf "%#v" .FileMode }}, time.Unix({{ .ModifiedAt.Unix }}, {{ .ModifiedAt.UnixNano }})),{{ end }}
		{{- end }}
	},
)
{{ if .HasFileData }}
// binary data
var (
	{{- range $p := .Paths }}{{ with (index $FileMap $p) }}
	{{if and (not .IsDir) ( ne (len .Data) 0) }}_{{ $VarName }}{{ sha1 $p}} = []byte({{ printData .Data }}){{ end }}{{ end }}
	{{- end }}
)
{{ end }}
`
