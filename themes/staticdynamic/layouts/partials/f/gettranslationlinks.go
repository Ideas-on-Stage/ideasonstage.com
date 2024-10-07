{{/* <!--

	f/gettranslationlinks.go
	
	Returns a list of links for all site languages:
	- gets the translation links for current page
	- completes the list with links to homepage for missing translations
	
	For example:
	- site languages are "en" "fr" "es" "it"
	- page is translated in "en" "es" "it"
	- current page is in "en"
	- gettranslationlinks will return
		- links to page translations for "es" and "it"
		- link to homepage for "fr"
	
	Arguments: 
	- page object
	
	Returns: 
	- links for all site languages

--> */}}

{{ $currentlangcode := site.LanguageCode | lower }}
{{ $foundlang := slice $currentlangcode }}
{{ $alllang := slice }}
{{ $result := slice }}

{{/* <!-- Retrieve the complete list of languages of the website --> */}}
{{ range site.Languages }}
	{{ $alllang = $alllang | append .LanguageCode }}
{{ end }}

{{/* <!-- Retrieve translations for current page --> */}}
{{ range .Translations }}
	{{ $entry := dict "name" .Language.LanguageName "code" .Language.LanguageCode "url" .Permalink }}
	{{ $result = $result | append $entry }}
	{{ $foundlang = $foundlang | append .Language.LanguageCode }}
{{ end }}

{{/* <!-- Identify missing translations, and add links to homepage for missing translations --> */}}
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
