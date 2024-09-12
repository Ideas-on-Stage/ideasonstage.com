{{/* <!--
	
	f/getstyle
	
	Get style parameter from .Params.style for page objects, from .style for other objects such as data objects.
	
	Arguments:
	- page or data object
	
	Returns:
	- style attribute as string to be added to class="style" in html code

 --> */}}

{{ $style := "" }}
{{ if (partial "f/ispage.go" .) }}
	{{ with .Params.style }}
		{{ $style = . }}
	{{ end }}
{{ else }}
	{{ with .style }}
		{{ $style = . }}
	{{ end }}
{{ end }}
{{ return $style }}