package pkg2

import (
	"fmt"
	"pkg1/pkg2"
) //1.在gopath里配置项目src路径 2.引用代码提示自动导入

func Test2(){
	abc:= pkg2.NewAbc("Joe", 25)
	fmt.Println(abc,abc.Name)
}
