{{/* <!--
	
	f/getheropicturepath
	
	Gets hero picture for a page (if it exists).
	- if a file named hero.webp exists in page directory, select it
	- else if a file named hero.jpg exists in page directory, select it
	- else if a picture is specified in .Params.picture, select it
	- else returns an empty string
	
	Offers the advantage of completing the path if only the image name is specified in the parameter, and also looks for a default hero.jpg image first.
	
	Arguments:
	- page object or path as string
	
	Returns:
	- name of hero picture
	- or empty string if not found.

--> */}}

{{ $imgpath := "" }}
{{ $picture := "" }}
{{ $found := false }}
{{ $thumbtest := "" }}
{{ $thumbtesta := "" }}
{{ $thumbtestb := "" }}

{{ if partial "f/ispage.go" . }}
	{{/* <!-- then this is a page, check for hero.jpg in same directory --> */}}
	{{ if .File }}
		{{ $thumbtest = .File.Dir }}
		{{ $thumbtesta = add $thumbtest "hero.webp" }}
		{{ $thumbtestb = add $thumbtest "hero.jpg" }}
		{{ if (fileExists $thumbtesta) }}
			{{/* <!-- then a file named hero.webp exists, use it --> */}}
			{{ $picture = "hero.webp" }}
			{{ $found = true }}
		{{ else if (fileExists $thumbtestb) }}
			{{/* <!-- then a file named hero.jpg exists, use it --> */}}
			{{ $picture = "hero.jpg" }}
			{{ $found = true }}
		{{ else if .Params.picture }}
			{{/* <!-- else if .picture exists, use it --> */}}
			{{ $picture = .Params.picture }}
			{{ $found = true }}
		{{ end }}
	{{ else }}
		{{ $thumbtest = "" }}
	{{ end }}
{{ end }}

{{ if $found }}
	{{ $imgpath = $picture }}
{{ end }}

{{ return $imgpath }}
