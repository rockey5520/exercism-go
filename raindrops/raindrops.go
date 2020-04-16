package raindrops

import (
	"strings"
	"strconv"
)
// Convert function converts integer to rain sounds
func Convert(a int) string {
	var sb strings.Builder
	if ( a%3 == 0 && a%5 == 0 && a%7 == 0)  {
		sb.WriteString("PlingPlangPlong")
	}else if ( a%3 != 0 && a%5 != 0 && a%7 != 0)  {
		v := int64(a)
		sb.WriteString(strconv.FormatInt(v, 10))
	}else{
		if(a%3 == 0){
			sb.WriteString("Pling")
		}
		if(a%5 == 0){
			sb.WriteString("Plang")
		}
		if(a%7 == 0){
			sb.WriteString("Plong")
		}
	}
	return sb.String()
}
