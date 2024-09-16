{{ $currentlangcode := site.LanguageCode | lower }}
{{ $foundlang := slice $currentlangcode }}
{{ $alllang := slice }}
{{ $result := slice }}

{{ range site.Languages }}
	{{ $alllang = $alllang | append .LanguageCode }}
{{ end }}

{{ range .Translations }}
	{{ $entry := dict "name" .Language.LanguageName "code" .Language.LanguageCode "url" .Permalink }}
	{{ $result = $result | append $entry }}
	{{ $foundlang = $foundlang | append .Language.LanguageCode }}
{{ end }}

{{ $missinglang := symdiff $alllang $foundlang }}
{{ range $missinglang }}
	{{ $lang := . }}
	{{ range site.Languages }}
		{{ if eq .LanguageCode $lang }}
		{{ $entry := dict "name" .LanguageName "code" .LanguageCode "url" (printf "%s%s" site.BaseURL .LanguageCode) }}
		{{ $result = $result | append $entry }}
		{{ end }}
	{{ end }}
{{ end }}

{{ return $result }}
 