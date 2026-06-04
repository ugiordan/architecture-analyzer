package flow

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed embed/flowlens.iife.js
var flowlensBundle []byte

const initScript = `(function(){var d=JSON.parse(document.getElementById('flowlens-data').textContent);var c=document.getElementById('canvas');var fl=new FlowLens.FlowLens(c,d);fl.play();}())`

// GenerateHTML produces a self-contained HTML page that embeds the flowlens
// IIFE bundle and the diagram JSON. The page uses a Content Security Policy
// with SHA-256 hashes so only the exact embedded scripts can execute.
func GenerateHTML(diagram Diagram) (string, error) {
	diagramJSON, err := json.Marshal(diagram)
	if err != nil {
		return "", fmt.Errorf("marshaling diagram: %w", err)
	}
	escapedJSON := escapeForHTML(string(diagramJSON))

	iifeScript := string(flowlensBundle)
	iifeHash := sha256Base64(iifeScript)
	initHash := sha256Base64(initScript)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html><head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<meta http-equiv="Content-Security-Policy" content="default-src 'none'; script-src 'sha256-%s' 'sha256-%s'; style-src 'unsafe-inline'">
<title>flowlens</title>
<style>*{margin:0;padding:0;box-sizing:border-box}html,body{width:100%%;height:100%%;overflow:hidden;background:#0d1117}canvas{display:block;width:100%%;height:100%%}</style>
</head><body>
<canvas id="canvas"></canvas>
<script type="application/json" id="flowlens-data">%s</script>
<script>%s</script>
<script>%s</script>
</body></html>`, iifeHash, initHash, escapedJSON, iifeScript, initScript)

	return html, nil
}

// escapeForHTML prevents embedded JSON from breaking out of a <script> tag.
// json.Marshal already escapes <, >, and & in string values; this function
// handles the remaining dangerous sequences in raw JSON output.
func escapeForHTML(s string) string {
	s = strings.ReplaceAll(s, "</", "<\\/")
	s = strings.ReplaceAll(s, "<!--", "<\\!--")
	return s
}

// sha256Base64 returns the base64-encoded SHA-256 digest of content,
// suitable for use in CSP hash directives.
func sha256Base64(content string) string {
	h := sha256.Sum256([]byte(content))
	return base64.StdEncoding.EncodeToString(h[:])
}
