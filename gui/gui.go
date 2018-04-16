package gui

import (
	"korok.io/korok/gfx"
	"korok.io/korok/gfx/font"
)

//	Awesome GUI System
//
type LayoutType int

const (
	Vertical   LayoutType = iota
	Horizontal
	OverLay
)

// Widgets: Text
func Text(id ID, bb Bound, text string, style *TextStyle) {
	if style == nil {
		style = &gContext.Style.Text
	}
	gContext.Text(id, text, bb, style)
	return
}

func TextSizeColored(text string, color uint32, size float32) {

}

// Widgets: InputEditor
func InputText(hint string, style *InputStyle) {

}

// Widget: Image
func Image(id ID, bb Bound, tex gfx.Tex2D, style *ImageStyle) {
	gContext.Image(id, tex, bb, style)
}

// Widget: Button
func Button(id ID, bb Bound, text string, style *ButtonStyle) (event EventType) {
	return gContext.Button(id, text, &bb, style)
}

func ImageButton(id ID, bb Bound, normal, pressed gfx.Tex2D, style *ImageButtonStyle) EventType{
	return gContext.ImageButton(id, normal, pressed, &bb, style)
}

func CheckBox(text string, style *CheckBoxStyle) bool {
	return false
}

// Widget: ProgressBar, Slider
func ProgressBar(fraction float32, style *ProgressBarStyle) {

}

func Slider(id ID, bb Bound, value *float32, style *SliderStyle) (v EventType){
	return gContext.Slider(id, value, &bb, style)
}

// Frame: Rect
func Rect(w, h float32, style *RectStyle) {
	if style == nil {
		style = &gContext.Style.Rect
	}
	gContext.DrawRect(&Bound{0, 0, w,h}, style.FillColor, style.Rounding)
}

// Widget: ListView TODO
func ListView() {

}

// 基于当前 Group 移动光标
func Offset(dx, dy float32) {
	gContext.Cursor.X += dx
	gContext.Cursor.Y += dy
}

func Move(x, y float32) {
	gContext.Cursor.X = x
	gContext.Cursor.Y = y
}

// Theme:
func UseTheme(style *Style) {
	gContext.UseTheme(style)
}

func SetFont(font font.Font) {
	gContext.Style.Text.Font = font
	gContext.Style.Button.Font = font
	texFont, _ := font.Tex2D()
	gContext.DrawList.PushTextureId(texFont)
}

func SetScreenSize(w, h float32) {
	screen.Width = w
	screen.Height = h
}

// gui init, render and destroy
func Init() {

}

func Frame() {

}

func Destroy() {

}

func DefaultContext() *Context {
	return gContext
}

var ThemeLight *Style
var ThemeDark  *Style

////////// implementation
// 应该设计一种状态管理机制，用这套机制来维护状态
// 比如：PopupWindow/Animation/Toast/
// 这些控件的典型特点是状态是变化的，而且不方便用代码维护
// 如果用一个专门的系统，问题可以减轻很多

var gContext *Context

func init() {
	// default theme
	ThemeLight = newLightTheme()
	ThemeDark  = newDarkTheme()

	// default context
	gContext = NewContext(ThemeLight)
}
