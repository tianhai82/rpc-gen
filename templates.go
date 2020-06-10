package rpc_gen

const importTemplate = `{{- if .NeedCaching -}}const objectHash = require('object-hash');{{- end -}}{{if .Classes}}
import {
{{range .Classes}}  {{.}},
{{end -}}} from './models';{{- end}}
`

const functionTemplate = `
export function {{.FunctionName}}({{if .Input}}data: {{.Input.ClassName}}{{if .Input.IsArray}}[]{{- end}}{{end}}){{if .Output}}: Promise<{{.Output.ClassName}}{{if .Output.IsArray}}[]{{- end}}> {{- end}} { 
{{- if .Cache -}}{{- if .Output -}}
{{- if .Input}}
  let hash: string | undefined;
  hash = objectHash(data); {{- end}}
{{- if .Input}}
  const value = cache.get(` + "`{{.FunctionName}}|${hash}`" + `);
  {{- else}}
  const value = cache.get('{{.FunctionName}}'); {{- end}}
  if (value) {
    return Promise.resolve(value);
  }
{{- end}}{{- end}}
  return fetch('{{.Path}}', {
    method: 'POST',
    mode: 'same-origin',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json',
    },
    redirect: 'follow',{{if .Input}}
    body: JSON.stringify(data),{{- end}}
  })
    .then((response) => {
      if (response.ok) {
        return response;
      }
      if (response.status === 401) {
        const src = window.location.pathname + window.location.search + window.location.hash;
        window.location.href = ` + "`/sign/in?source=${src}`" + `;
      }
      throw new Error(` + "`${response.statusText} : ${response.status}`" + `);
    })
	{{- if .Output}}
    .then(response => response.json()){{if .Cache}}
    .then((value) => {
    {{- if .Input}}
      const key = ` + "`{{.FunctionName}}|${hash}`;" + `
	{{- else}}
      const key = '{{.FunctionName}}';	 
	{{- end}}
      cache.set(key, value);
      return value;
    }){{end}}{{end}};
}
`
