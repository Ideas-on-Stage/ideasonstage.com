{{/* <!--
	
	f/getimgpathhero
	
	Gets hero picture for a page (if it exists).
	- if a file named hero.webp exists in page directory, select it
	- else if a file named hero.jpg exists in page directory, select it
	- else if a picture is specified in .Params.picture, select it
	- else returns an empty string
	
	Offers the advantage of completing the path if only the image name is specified in the parameter, and also looks for a default hero.jpg image first.
	
	Arguments:
	- page object
	
	Returns:
	- name of hero picture
	- or empty string if not found.

--> */}}

{{ $data := partial "f/getdata" . }}
{{ $imgpath := "" }}
{{ $img := "" }}
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
			{{ $img = "hero.webp" }}
			{{ $found = true }}
		{{ else if (fileExists $thumbtestb) }}
			{{/* <!-- then a file named hero.jpg exists, use it --> */}}
			{{ $img = "hero.jpg" }}
			{{ $found = true }}
		{{ else if .Params.img }}
			{{/* <!-- else if .picture exists, use it --> */}}
			{{ $img = .Params.img }}
			{{ $found = true }}
		{{ else if .Params.picture }}
			{{/* <!-- else if .picture exists, use it --> */}}
			{{ $img = .Params.picture }}
			{{ $found = true }}
		{{ end }}
	{{ else }}
		{{ $found = false }}
	{{ end }}
{{ end }}

{{ if $found }}
	{{/* <!-- if it's a string check if path should be completed --> */}}
	{{/* <!-- if string contains at least one /... --> */}}
	{{ if strings.Contains $img "/" }}
		{{/* <!-- ...then it already contains path to img --> */}}
		{{ $imgpath = $img }}
	{{ else }}
		{{/* <!-- else add path --> */}}
		{{ if (isset $data "url") }}
			{{ $imgpath = path.Join $data.url $img }}
		{{ else }}
			{{ $imgpath = path.Join page.RelPermalink $img }}
		{{ end }}
	{{ end }}
{{ end }}

{{ return $imgpath }}
