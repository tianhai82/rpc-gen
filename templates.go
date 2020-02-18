package rpc_gen

const importTemplate = `import objectHash from 'object-hash';
import {
{{range .Classes}}  {{.}},
{{end}}} from '../apis/models';
`

const functionTemplate = `
export function {{.FunctionName}}({{if .Input}}data: {{.Input.ClassName}}{{if .Input.IsArray}}[]{{- end}}, {{end}}toCache: boolean){{if .Output}}: Promise<{{.Output.ClassName}}{{if .Output.IsArray}}[]{{- end}}> {{- end}} { {{- if .Input}}
  let hash: string | undefined; {{- end}}
  if (toCache) { {{- if .Input}}
    hash = objectHash(data); {{- end}}	{{- if .Input}}
    const value = cache.get(` + "`{{.FunctionName}}|${hash}`" + `);
	{{- else}}
    const value = cache.get('{{.FunctionName}}'); {{- end}}
    if (value) {
      return Promise.resolve(value);
    }
  }
  const req: JsonRequest = { {{- if .Input}}
    args: data,
    hash,{{- end}}
    toCache,
  };
  return fetch('{{.Path}}', {
    method: 'POST',
    mode: 'same-origin',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json',
    },
    redirect: 'follow',
    body: JSON.stringify(req),
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
    .then(response => response.json())
    .then((value) => {
      if (toCache) {
      {{- if .Input}}
        const key = ` + "`{{.FunctionName}}|${hash}`;" + `
	  {{- else}}
        const key = '{{.FunctionName}}';	 
	  {{- end}}
        cache.set(key, value);
      }
      return value;
    }){{end}};
}
`
