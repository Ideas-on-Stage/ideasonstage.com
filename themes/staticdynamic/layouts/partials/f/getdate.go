{{/* <!--

	f/getdate
	
	Get a human readable date string localized in the current language.
	See additional comments in f/getdatetime

	Arguments: 
	- date as string or date object
	
	Returns: 
	- localized human readable date string with format "Monday 1 March 2000"

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
{{ $datetxt := $date.Format "Monday 2 January 2006 15:00 MST" }}
{{ $dateelements := split $datetxt " " }}

{{/* <!-- get localized day name --> */}}
{{ $dayname := index $dateelements 0 }}
{{/* <!-- get array of localized day names  --> */}}
{{ $days := index site.Data.multilingual.days $lang }}
{{/* <!-- get matching day translation --> */}}
{{ $dayint := index $days $dayname }}

{{/* <!-- get localized month name --> */}}
{{ $monthname := printf "%s" $date.Month }}
{{/* <!-- get array of localized month names  --> */}}
{{ $months := index site.Data.multilingual.months $lang }}
{{/* <!-- get matching month translation --> */}}
{{ $monthint := index $months $monthname }}

{{/* <!-- get localized time zone name --> */}}
{{ $timezonename := index $dateelements 5 }}
{{/* <!-- get array of localized time zone names  --> */}}
{{ $timezones := index site.Data.multilingual.timezones $lang }}
{{/* <!-- get matching time zone translation --> */}}
{{ $timezoneint := index $timezones $timezonename }}

{{/* <!-- build localized date string --> */}}
{{ $dateout := add $dayint " " }}
{{ $dateout = add $dateout (index $dateelements 1) }}
{{ $dateout = add $dateout " " }}
{{ $dateout = add $dateout $monthint }}
{{ $dateout = add $dateout " " }}
{{ $dateout = add $dateout (index $dateelements 3) }}
{{ return $dateout }}
