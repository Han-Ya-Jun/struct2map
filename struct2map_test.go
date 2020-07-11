package struct2map

import "testing"

/*
* @Author: yajun.han
* @Date: 2020/7/12 2:00 上午
* @Name：struct2map
* @Description:
 */

func TestStruct2mapCase1(t *testing.T) {
	type User struct {
		UserName string `struct2map:"user_name,omitempty" json:"user_name"`
		Age      int    `struct2map:"-" json:"age"`
	}
	user := User{
		UserName: "",
		Age:      23,
	}
	struct2map, err := Struct2map(&user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", struct2map)
}

func TestStruct2mapCase2(t *testing.T) {
	type User struct {
		UserName string `struct2map:"user_name,omitempty" json:"user_name"`
		Age      *int   `struct2map:"age,omitempty" json:"age"`
	}
	user := User{
		UserName: "qwe",
		Age:      nil,
	}
	struct2map, err := Struct2map(&user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", struct2map)
}

func BenchmarkStruct2map(b *testing.B) {
	type User struct {
		UserName string `struct2map:"user_name,omitempty" json:"user_name"`
		Age      int    `struct2map:"-" json:"age"`
	}
	user := User{
		UserName: "",
		Age:      23,
	}
	struct2map, err := Struct2map(&user)
	if err != nil {
		b.Fatal(err)
	}
	b.Logf("%v", struct2map)
}
