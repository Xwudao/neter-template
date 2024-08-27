{{define "pages/index.tpl"}}
  {{- /*gotype: github.com/Xwudao/neter-template/internal/domain/payloads.IndexMap*/ -}}
  <!doctype html>
  <html lang="zh-CN">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.Title}}</title>
  </head>
  <body>
  hello, {{.MainTitle}}
  </body>
  </html>

{{end}}
