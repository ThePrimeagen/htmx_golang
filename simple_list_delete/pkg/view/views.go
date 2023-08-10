package view

var Index = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title></title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <script src="https://unpkg.com/htmx.org@1.9.4" integrity="sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV" crossorigin="anonymous"></script>
        <link href="css/style.css" rel="stylesheet">
    </head>
    <body>
        <div class="container">
            {{template "items" . }}
        </div>
    </body>
</html>`

var Items = `{{define "items"}}
{{range .Items}}
    {{template "item" .}}
{{end}}
{{end}}`

var ItemCount = `{{define "item-count"}}
<div hx-trigger="click"
    hx-post="/count/{{.Id}}"
    hx-swap="outerHTML"
    class="count">{{.Count}}</div>
{{end}}`

var Item = `{{define "item"}}
<div class="item">
    <div class="name">{{.Name}}</div>
    {{template "item-count" .}}
    <form hx-trigger="click" hx-delete="/item/{{.Id}}" hx-target="closest .item">
        Delete
    </form>
</div>
{{end}}`





