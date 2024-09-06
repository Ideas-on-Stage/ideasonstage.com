{{/* <!--
	
	f/getheropicturepath
	
	Gets full relative path to hero picture for pages.
	Offers the advantage of completing the path if only the image name is specified in the parameter, and also looks for a default hero.jpg image first.
	
	Arguments:
	- page object or path as string
	
	Process:
	- If argument is a valid page object:
		- if a file named hero.jpg exists in page directory, select it
		- else if a picture is specified in .Params.picture, select it
	- If argument is a string (implicitly a path)
		- starts with "http" then it's an absolute path, return it
		- starts with "/" then it's a relative path, return it
		- else it's only the picture name, add relative path
		
	Returns:
	- (absolute "https://…" or relative "/…" path) to (hero pic or .Params.picture)
	- or empty string if not found.

 --> */}}

{{ $imgpath := "" }}
{{ $picture := "" }}
{{ $found := false }}
{{ $thumbtest := "filenotfound" }}

{{ if partial "f/ispage.go" . }}
	{{/* <!-- then this is a page, check for hero.jpg in same directory --> */}}
	{{ if .File }}
		{{ $thumbtest = .File.Dir }}		
	{{ else }}
		{{ $thumbtest = "filenotfound" }}
	{{ end }}
	{{ $thumbtest = add $thumbtest "hero.jpg" }}
	{{ if (fileExists $thumbtest) }}
		{{/* <!-- then a file named hero.jpg exists, use it --> */}}
		{{ $picture = "hero.jpg" }}
		{{ $found = true }}
	{{ else if .Params.picture }}
		{{/* <!-- else if .picture exists, use it --> */}}
		{{ $picture = .Params.picture }}
		{{ $found = true }}
	{{ end }}
{{ end }}

{{ if $found }}
	{{ $imgpath = $picture }}
{{ end }}

{{ return $imgpath }}