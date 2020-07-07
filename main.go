package main

import (
	"fmt"
	"gengine/base"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	Name string
	Age  int
	Male bool
}

func (u *User)GetNum(i int64) int64 {
	return i
}

func (u *User)Print(s string){
	fmt.Println(s)
}

func (u *User)Say(){
	fmt.Println("hello world")
}

const (
	rule2 = `
rule "测试" "测试描述"  salience 0 
begin
		// 重命名函数 测试
		Sout(@name)
		// 普通函数 测试
		Hello()
		//结构提方法 测试
		User.Say()
		// if
		if 7 == User.GetNum(7){
			//自定义变量 和 加法 测试
			variable = "hello" + " world"
			// 加法 与 内建函数 测试
			User.Name = "hhh" + strconv.FormatBool(true)
			//结构体属性、方法调用 和 除法 测试
			//User.Age = User.GetNum(12345) / 1000
			//布尔值设置 测试
			User.Male = false
			//规则内自定义变量调用 测试
			User.Print(variable)
			//float测试	也支持科学计数法		
			f = 9.56			
			PrintReal(f)
			//嵌套if-else测试
			if false	{
				Sout("嵌套if测试")
			}else{
				Sout("嵌套else测试")
			}
		}else{ //else
			//字符串设置 测试
			User.Name = "yyyy"
		}
end`)

func Hello()  {
	fmt.Println("hello")
}

func PrintReal(real float64){
	fmt.Println(real)
}

func exe(user *User){
	/**
	不要注入除函数和结构体指针以外的其他类型(如变量)
	*/
	dataContext := context.NewDataContext()
	//注入结构体指针
	dataContext.Add("User", user)
	//重命名函数,并注入
	dataContext.Add("Sout",fmt.Println)
	//直接注入函数
	dataContext.Add("Hello",Hello)
	dataContext.Add("PrintReal",PrintReal)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(rule2)
	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
		logrus.Infof("user.Age=%d,Name=%s,Male=%t", user.Age, user.Name, user.Male)
	}
}

func main() {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}
	exe(user)
}
