{{/* <!--

	f/gettime
	
	Gets the human readable time with time zone.
	See additional comments in f/getdatetime
	
	Arguments: 
	- date as string or date object
	
	Returns: 
	- localized time string with format "15:00 CEST (Paris Time)"

--> */}}

{{/* <!-- get language --> */}}
{{ $lang := site.LanguageCode }}
{{ if isset site.Data.multilingual.days $lang }}
{{ else }}
	{{/* <!-- default to English if language not set --> */}}
	{{ $lang = "en" }}
{{ end }}

{{/* <!-- convert ISO datetime string to datetime object --> */}}
{{ $datestr := printf "%s" . }}
{{ $date := time $datestr }}

{{/* <!-- convert datetime object to local time --> */}}
{{ $date = $date.Local }}

{{/* <!-- use datetime to create human readable date string --> */}}
{{ $datetxt := $date.Format "Monday 2 January 2006 15:04 MST" }}
{{ $dateelements := split $datetxt " " }}

{{/* <!-- get localized time zone name --> */}}
{{ $timezonename := index $dateelements 5 }}
{{/* <!-- get array of localized time zone names  --> */}}
{{ $timezones := index site.Data.multilingual.timezones $lang }}
{{/* <!-- get matching time zone translation --> */}}
{{ $timezoneint := index $timezones $timezonename }}

{{/* <!-- build localized date string --> */}}
{{ $dateout := index $dateelements 4 }}
{{ $dateout = add $dateout " " }}
{{ $dateout = add $dateout $timezoneint }}
{{ return $dateout }}
