package main

import "fmt"

//设计说明
//组件：包括 name,type 属性,Draw 方法
//容器是可以放组件的组件
//窗口，frame 是容器
//picture,button,lable.textbox paswordbox checkbox linklabel是组件，不是容器

type Component interface {
	Draw()
}
type ComponentContainer interface {
	Component
	AddComponent(subComponent Component)
}
type BaseComponent struct {
	name          string
	componentType string
}

func (base *BaseComponent) GetName() string {
	return base.name
}
func (base *BaseComponent) GetType() string {
	return base.componentType
}
func (base *BaseComponent) Draw() {
	drawContent := fmt.Sprintf("print %s(%s)", base.GetType(), base.GetName())
	fmt.Println(drawContent)
}
func NewBaseComponent(componentType, name string) *BaseComponent {
	component := new(BaseComponent)
	component.componentType = componentType
	component.name = name
	return component
}

type BaseComponentContainer struct {
	Component
	subComponents []Component
}

func (this *BaseComponentContainer) AddComponent(subComponent Component) {
	this.subComponents = append(this.subComponents, subComponent)
}
func (base *BaseComponentContainer) Draw() {
	base.Component.Draw()
	for _, sub := range base.subComponents {
		sub.Draw()
	}
}
func NewBaseComponentContainer(componentType, name string) *BaseComponentContainer {
	container := new(BaseComponentContainer)
	container.Component = NewBaseComponent(componentType, name)
	return container
}

func NewWinForm(name string) ComponentContainer {
	return NewBaseComponentContainer("WinForm", name)
}
func NewPicture(name string) Component {
	return NewBaseComponent("Picture", name)
}
func NewButton(name string) Component {
	return NewBaseComponent("Button", name)
}
func NewFrame(name string) ComponentContainer {
	return NewBaseComponentContainer("Frame", name)
}
func NewLabel(name string) Component {
	return NewBaseComponent("Lable", name)
}
func NewTextBox(name string) Component {
	return NewBaseComponent("TextBox", name)
}
func NewPasswordBox(name string) Component {
	return NewBaseComponent("PasswordBox", name)
}
func NewCheckBox(name string) Component {
	return NewBaseComponent("CheckBox", name)
}
func NewLinkLable(name string) Component {
	return NewBaseComponent("LinkLable", name)
}

func main() {
	createSampleWindow().Draw()
}
func createSampleWindow() Component {
	var winFrom = NewWinForm("WINDOW窗口")
	winFrom.AddComponent(NewPicture("LOGO图片"))
	winFrom.AddComponent(NewButton("登录"))
	winFrom.AddComponent(NewButton("注册"))
	var frame = NewFrame("FRAME1")
	frame.AddComponent(NewLabel("用户名"))
	frame.AddComponent(NewTextBox("文本框"))
	frame.AddComponent(NewLabel("密码"))
	frame.AddComponent(NewPasswordBox("密码框"))
	frame.AddComponent(NewCheckBox("记住用户名"))
	frame.AddComponent(NewLinkLable("忘记密码"))
	winFrom.AddComponent(frame)
	return winFrom
}
