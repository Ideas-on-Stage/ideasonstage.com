{{/* <!--
	
	f/geturl.go
	
	Gets url parameter for pages and data files.
	- If page then use .Permalink
	- else try to use the .href property
	- if not found then try to use .link property
	- if not found then try to use .url property
	
	Arguments:
	- page, data object, or a structure containing a .link or .url property
	
	Returns:
	- url as a string
	- empty string if not found

--> */}}

{{ $result := "" }}
{{ if partial "f/ispage.go" . }}
	{{ $result = .RelPermalink }}
{{ else if isset . "href" }}
	{{ $result = .href }}
{{ else if isset . "link" }}
	{{ $result = .link }}
{{ else if isset . "url" }}
	{{ $result = .url }}
{{ end }}

{{ return $result }}