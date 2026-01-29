{{/* <!--
	
	f/getimgsrcset

	Returns a slice with three things:
	- the "src" path for the fallback picture
	- the "srcset" paths for the list of responsive pictures
	- the sizes for the pictures
	
	Arguments:
	- image path as string
	
	Returns:
	- a slice with:
		- "src" as src attribute
		- "srcset" with list of image sources for different sizes
		- "sizes" as list of sizes
		
	Example code:
	{{ $srcset := partial "f/getimgsrcset.go" "/img/picture.jgp" }}
	<img src="{{ index $srcset 0 }}" srcset="{{ index $srcset 1 }}" sizes="{{ index $srcset 2 }}">

--> */}}

{{/* <!-- initialize variables --> */}}
{{ $imgSrc := false }}
{{ $imgSrcSet := false }}
{{ $imgSizes := slice }}
{{ $scale := 0 }}
{{ $size := "" }}
{{ $srcUrl := "" }}

{{ $image := page.Resources.Get . }}
{{ if not $image }}
	{{ $image = resources.Get . }}
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
			{{/* <!-- initialize variables --> */}}
			{{ $scale = .scale }}
			{{ $size = .size }}
			{{ $srcUrl = "" }}
			{{/* <!-- if width greater than scale --> */}}
			{{ if gt $image.Width $scale }}
				{{ $srcUrl = (printf "%dx webp" $scale | $image.Resize).RelPermalink }}
			{{ else }}
				{{ $srcUrl = (printf "%dx%d webp" $image.Width $image.Height | $image.Resize).RelPermalink }}
			{{ end }}
			{{ $imgSrc = $srcUrl }}
			{{ $imgSrcSet = $imgSrcSet | append (printf "%s %dw" $srcUrl $scale) }}
			{{ $imgSizes = $imgSizes | append (printf "%s" $size ) }}
		{{ end }}
		{{ $sizes := (delimit $imgSizes ",") }}
		{{ $set := delimit $imgSrcSet "," }}
		{{ $imgSrcSet = slice $imgSrc $set $sizes }}
	{{ else }}
		{{ $imgSrcSet = slice $image false false }}
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
