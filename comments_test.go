package weigo

// import (
// 	"testing"
// )

// var api *APIClient

// func init() {
// 	if api == nil {
// 		api = NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
// 		api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
// 	}
// }

// func Test_GET_comments_show(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"id": "3551749023600582",
// 	}
// 	result := new(Comments)
// 	err := api.GET_comments_show(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(len(*result.Comments))
// }

// func Test_POST_comments_create(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"id":      "3551749023600582",
// 		"comment": "Testing...Testing...",
// 	}
// 	result := new(Comment)
// 	err := api.POST_comments_create(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(*result)
// }
