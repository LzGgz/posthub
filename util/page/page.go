package page

import "github.com/spf13/viper"

func Offset(pageNum int) (offset int) {
	if pageNum > 0 {
		offset = (pageNum - 1) * viper.GetInt("app.page_size")
	}
	return
}
