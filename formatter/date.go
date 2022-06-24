package formatter

const TimeZoneSantiago = "America/Santiago"
const TimeZoneArgentinaBuenosAires = "America/Argentina/Buenos_Aires"
const TimeZoneColombia = "America/Bogota"
const TimeZonePeru = "America/Lima"
const TimeZoneUtc = "Etc/UTC"

func GetTimeZone(country string) string {
	switch country {
	case "CL":
		return TimeZoneSantiago

	case "PE":
		return TimeZonePeru

	case "CO":
		return TimeZoneColombia

	case "AR":
		return TimeZoneArgentinaBuenosAires

	default:
		return TimeZoneUtc

	}
}
