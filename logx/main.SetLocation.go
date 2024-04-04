package logx

// loggingLocationName
var loggingLocationName = "UTC"

// SetLocation
//
// e.g.
//
//	logx.SetLocation("Asia/Shanghai")
func SetLocation(locationName string) {
	loggingLocationName = locationName
}
