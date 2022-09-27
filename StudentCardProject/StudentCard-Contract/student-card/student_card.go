package student_card

import "encoding/json"

type StudentCard struct {
	Sno  string `json:"sno"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//  实现StudentCard的基本数据处理

func (sc *StudentCard) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, sc)
}
func (sc *StudentCard) MarshalJson() (data []byte, err error) {
	data, err = json.Marshal(sc)
	return data, err
}
func (sc *StudentCard) Serialize() ([]byte, error) {
	return json.Marshal(sc)
}
func Deserialize(data []byte, sc *StudentCard) error {
	return sc.UnmarshalJson(data)
}

//实现State接口
//两个函数Serialize()和GetSplitKey()

// GetSplitKey 返回StudentCard的主键字符串集
func (sc *StudentCard) GetSplitKey() []string {
	return []string{sc.Sno}
}
