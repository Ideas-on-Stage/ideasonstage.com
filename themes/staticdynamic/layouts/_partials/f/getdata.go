{{/* <!--
	
	f/getdata

	Returns a "harmonized" data structure, removing .Params if necessary.
	e.g. the structure:
	.Params
		.title
		.description
	will become
	.title
	.description

	the structure	
	.title
	.description
	will remeain 
	.title
	.description

	Arguments:
	- .Params.data or .data
	
	Returns:
	- .data

--> */}}

{{ $data := . }}
{{ if .Params }}
	{{ $data = .Params }}
{{ end }}
{{ return $data }}
