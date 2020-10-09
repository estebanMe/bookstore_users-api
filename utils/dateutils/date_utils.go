package dateutils

import "time"


const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)


func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString  return date Now with custom format 
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}