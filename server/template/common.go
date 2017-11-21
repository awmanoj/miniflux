// Code generated by go generate; DO NOT EDIT.
// 2017-11-20 17:09:36.256513528 -0800 PST m=+0.022883883

package template

var templateCommonMap = map[string]string{
	"entry_pagination": `{{ define "entry_pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .prevEntry }}
            <a href="{{ .prevEntryRoute }}" title="{{ .prevEntry.Title }}" data-page="previous">{{ t "Previous" }}</a>
        {{ else }}
            {{ t "Previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .nextEntry }}
            <a href="{{ .nextEntryRoute }}" title="{{ .nextEntry.Title }}" data-page="next">{{ t "Next" }}</a>
        {{ else }}
            {{ t "Next" }}
        {{ end }}
    </div>
</div>
{{ end }}`,
	"layout": `{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width">
    <meta name="robots" content="noindex,nofollow">
    <meta name="referrer" content="no-referrer">
    {{ if .csrf }}
        <meta name="X-CSRF-Token" value="{{ .csrf }}">
    {{ end }}
    <title>{{template "title" .}} - Miniflux</title>
    {{ if .user }}
        <link rel="stylesheet" type="text/css" href="{{ route "stylesheet" "name" .user.Theme }}">
    {{ else }}
        <link rel="stylesheet" type="text/css" href="{{ route "stylesheet" "name" "white" }}">
    {{ end }}
    <script type="text/javascript" src="{{ route "javascript" }}" defer></script>
</head>
<body data-entries-status-url="{{ route "updateEntriesStatus" }}">
    {{ if .user }}
    <header class="header">
        <nav>
            <div class="logo">
                <a href="{{ route "unread" }}">Mini<span>flux</span></a>
            </div>
            <ul>
                <li {{ if eq .menu "unread" }}class="active"{{ end }}>
                    <a href="{{ route "unread" }}" data-page="unread">{{ t "Unread" }}</a>
                    {{ if gt .countUnread 0 }}
                        <span class="unread-counter" title="Unread articles">({{ .countUnread }})</span>
                    {{ end }}
                </li>
                <li {{ if eq .menu "history" }}class="active"{{ end }}>
                    <a href="{{ route "history" }}" data-page="history">{{ t "History" }}</a>
                </li>
                <li {{ if eq .menu "feeds" }}class="active"{{ end }}>
                    <a href="{{ route "feeds" }}" data-page="feeds">{{ t "Feeds" }}</a>
                </li>
                <li {{ if eq .menu "categories" }}class="active"{{ end }}>
                    <a href="{{ route "categories" }}" data-page="categories">{{ t "Categories" }}</a>
                </li>
                <li {{ if eq .menu "settings" }}class="active"{{ end }}>
                    <a href="{{ route "settings" }}" data-page="settings">{{ t "Settings" }}</a>
                </li>
                <li>
                    <a href="{{ route "logout" }}" title="Logged as {{ .user.Username }}">{{ t "Logout" }}</a>
                </li>
            </ul>
        </nav>
    </header>
    {{ end }}
    <section class="main">
        {{template "content" .}}
    </section>
</body>
</html>
{{ end }}`,
	"pagination": `{{ define "pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .ShowPrev }}
            <a href="{{ .Route }}{{ if gt .PrevOffset 0 }}?offset={{ .PrevOffset }}{{ end }}" data-page="previous">{{ t "Previous" }}</a>
        {{ else }}
            {{ t "Previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .ShowNext }}
            <a href="{{ .Route }}?offset={{ .NextOffset }}" data-page="next">{{ t "Next" }}</a>
        {{ else }}
            {{ t "Next" }}
        {{ end }}
    </div>
</div>
{{ end }}
`,
}

var templateCommonMapChecksums = map[string]string{
	"entry_pagination": "f1465fa70f585ae8043b200ec9de5bf437ffbb0c19fb7aefc015c3555614ee27",
	"layout":           "8be69cc93fdc99eb36841ae645f58488bd675670507dcdb2de0e593602893178",
	"pagination":       "6ff462c2b2a53bc5448b651da017f40a39f1d4f16cef4b2f09784f0797286924",
}
