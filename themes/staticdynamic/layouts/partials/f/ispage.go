{{/* <!--
	
	f/ispage
	
	Test if object is a page or list page object. Useful to determine if the object passed as argument is a page or a data file.
	
	Arguments:
	- something (if it's a page, it will return true, else it will return false)
	
	Returns:
	- true if page or list page (node), 
	- false if not
	
 --> */}}
 
{{ $result := false }}

{{ if (partial "f/isstring.go" .) }}
	{{ $result = false }}
{{ else if or .IsPage .IsNode }}
	{{ $result = true }}
{{ else }}
	{{ $result = false }}
{{ end }}

{{ return $result }}