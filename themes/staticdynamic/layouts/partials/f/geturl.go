{{/* <!--
	
	f/geturl.go
	
	Gets url parameter for pages and data files.
	- If page then use .Permalink
	- else try to use .link property
	- if not found, try to use .url property
	- if not found, try to use the .href property
	
	Arguments:
	- page, data object, or a structure containing a .link or .url property
	
	Returns:
	- url as a string
	- empty string if not found

--> */}}

{{ $result := "" }}
{{ if partial "f/ispage.go" . }}
	{{ $result = .RelPermalink }}
{{ else if isset . "link" }}
	{{ $result = .link }}
{{ else if isset . "url" }}
	{{ $result = .url }}
{{ else if isset . "href" }}
	{{ $result = .href }}
{{ end }}

{{ return $result }}