package dateutils

import "time"


const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

//GetNow return time now
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString  return date Now with custom format 
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}

//GetNowDBFormat return formate date
func GetNowDBFormat()string {
	return GetNow().Format(apiDbLayout)
}