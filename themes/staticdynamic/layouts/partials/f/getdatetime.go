{{/* <!--

	f/getdatetime

	Get a human readable date and time string localized in the current language.

	Configure additional languages in /data/multilingual.toml
	
	The date must be in a valid ISO 8601 format
	https://en.wikipedia.org/wiki/ISO_8601
	
	The date returned uses the local time zone of the machine running Hugo.
	To set the local time zone on Netlify, use the "TZ" build environment variable.
	Set TZ to a value from the TZ time zones:
	https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	
	Arguments: 
	- date as string or date object
	
	Returns: 
	- localized date and time string with format "Monday 1 March 2000 15:00 CEST"
	
--> */}}

{{ $date := partial "f/getdate.go" . }}
{{ $time := partial "f/gettime.go" . }}
{{ $dateout := add $date " " }}
{{ $dateout = add $dateout $time }}
{{ return $dateout }}
