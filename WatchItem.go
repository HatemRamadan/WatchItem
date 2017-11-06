package WatchItem
import (
	"encoding/json"
)
type WatchItem struct {
	Name string `json:"name"`
	Overview string `json:"overview"`
	ID uint `json:"id"`
}
	//
func DecodeItem(watchItemJSON map[string]interface{}) *WatchItem{
	var Name = watchItemJSON["name"]
	if Name==nil{
		Name = watchItemJSON["title"]
	}
	return &WatchItem{Name.(string),watchItemJSON["overview"].(string),uint(watchItemJSON["id"].(float64))}
}
func (movie *WatchItem) EncodeItem() ([]byte,error){
	movJS,err := json.Marshal(movie)
	return []byte(movJS),err
}