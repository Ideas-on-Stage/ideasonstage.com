{{/* <!--
	
	f/getimgsrcset

	Returns a slice with three things:
	- the "src" path for the fallback picture
	- the "srcset" paths for the list of responsive pictures
	- the sizes for the pictures
	
	Arguments:
	- Page object (with .Params.picture defined) or a string with a path to a global resource in /assets.
	
	Returns:
	- a slice with "src" and "srcset" to be used as in the example below.
	
	Example code:
	{{ $srcset := partial "f/getimgsrcset.go" "/img/picture.jgp" }}
	<img src="{{ collections.Index $srcset 0 }}" srcset="{{ collections.Index $srcset 1 }}" alt="" class="">

--> */}}

{{ $imgSrc := "" }}
{{ $imgSrcSet := false }}
{{ $imgSizes := slice }}
{{ $image := false }}

{{ if . }}
	{{ if ne "" . }}
		{{ if (partial "f/ispage.go" .) }}
			<!-- 1. Try to get image as page resource if argument is a page object -->
			{{ $image = .Resources.Get .Params.picture }}
		{{ else if (partial "f/isstring.go" .) }}
			<!-- 2. Try to get image as page resource, but with current page as reference -->
			{{ $image = page.Resources.Get . }}
			{{ if not $image }}
				<!-- 3. If no result, try to get image from global resources in /assets -->
				{{ $image = resources.Get . }}
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}

{{ if $image }}
	<!-- scale pic if it is not a SVG (SVGs don't need to and cannot be scaled) -->
	{{ if ne "svg" $image.MediaType.SubType }}
		<!-- uses settings from config.toml depending on orientation -->
		{{ $scales := site.Params.images.landscape }}
		{{ if gt $image.Height $image.Width }}
			{{ $scales = site.Params.images.portrait }}
		{{ end }}
		
		<!--
		Add URL for each width to $imgSrcSet variable
		format: "/path/img_1000.jpg 1000w,/path/img_500.jpg 500w"
		Note: the first URL is used as "fallback" src in $imgSrc.
		-->
		{{ $imgSrcSet = slice }}
		{{ range $scales }}
			{{ $scale := .scale }}
			{{ $size := .size }}
			{{ if (strings.Contains $image.RelPermalink "logos") }}
				{{ $scale = div $scale 4 }}
			{{ end }}
			{{ $srcUrl := "" }}
			{{ if gt $image.Width $scale }}
				{{ $srcUrl = (printf "%dx webp" $scale | $image.Resize).RelPermalink }}
			{{ else }}
				{{ $srcUrl = (printf "%dx%d webp" $image.Width $image.Height | $image.Resize).RelPermalink }}
			{{ end }}
			{{ if eq $imgSrc "" }}
				{{ $imgSrc = $srcUrl }}
			{{ end }}
			{{ $imgSrcSet = $imgSrcSet | append (printf "%s %dw" $srcUrl $scale) }}
			{{ $imgSizes = $imgSizes | append (printf "%s" $size ) }}
		{{ end }}
		{{ $sizes := (delimit $imgSizes ",") }}
		{{ $set := delimit $imgSrcSet "," }}
		{{ $imgSrcSet = slice $imgSrc $set $sizes }}
	{{ else }}
		{{ $imgSrcSet = slice . false false }}
	{{ end }}
{{ else }}
	<!-- If no resource found, fall back to using default image -->
	{{ if . }}
		{{ if ne "" . }}
			{{ $imgSrcSet = slice . false false }}
		{{ end }}
	{{ end }}
{{ end }}
{{ return $imgSrcSet }}
