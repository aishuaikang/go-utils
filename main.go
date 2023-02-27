package utils

import (
	"errors"
	"math"
	"syscall"
	"unsafe"
)

var (
	dllPath              = "utils.dll"
	utils                = syscall.NewLazyDLL(dllPath)
	_GetCursorPos        = utils.NewProc("GetCursorPos")
	_MoveMouse           = utils.NewProc("MoveMouse")
	_GetScreenResolution = utils.NewProc("GetScreenResolution")
	_LinearInit          = utils.NewProc("LinearInit")
	_MagneticInit        = utils.NewProc("MagneticInit")
	_PidInit             = utils.NewProc("PidInit")
	_SpeedInit           = utils.NewProc("SpeedInit")

	_CaptureInit    = utils.NewProc("CaptureInit")
	_CaptureRelease = utils.NewProc("CaptureRelease")
	_CaptureBitmap  = utils.NewProc("CaptureBitmap")
	_CaptureBmp     = utils.NewProc("CaptureBmp")

	_GetHwndByTitle = utils.NewProc("GetHwndByTitle")
	_GetWindowRect  = utils.NewProc("GetWindowRect")

	_IsWindowIconic  = utils.NewProc("IsWindowIconic")
	_IsWindowVisible = utils.NewProc("IsWindowVisible")
)

// 获取鼠标当前移动速度
func GetCursorPos() (currentX, currentY float64) {
	if code, _, _ := _GetCursorPos.Call(uintptr(unsafe.Pointer(&currentX)), uintptr(unsafe.Pointer(&currentY))); code != 0 {
		panic(errors.New("GetCursorPos call failed"))
	}
	return
}

// 根据鼠标当前位置移动
func MoveMouse(dx, dy float64) {
	if code, _, _ := _MoveMouse.Call(uintptr(dx), uintptr(dy)); code != 0 {
		panic(errors.New("MoveMouse call failed"))
	}
}

// 获取屏幕分辨率
func GetScreenResolution() (width, height int32) {
	if code, _, _ := _GetScreenResolution.Call(uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height))); code != 0 {
		panic(errors.New("GetScreenResolution call failed"))
	}
	return
}

// 线性算法初始化
func LinearInit(t float64) {
	if code, _, _ := _LinearInit.Call(uintptr(math.Float64bits(t))); code != 0 {
		panic(errors.New("LinearInit call failed"))
	}
}

// 磁吸算法初始化
func MagneticInit(k float64) {
	if code, _, _ := _MagneticInit.Call(uintptr(math.Float64bits(k))); code != 0 {
		panic(errors.New("MagneticInit call failed"))
	}
}

// PID算法初始化
func PidInit(p, i, d float64) {
	if code, _, _ := _PidInit.Call(uintptr(math.Float64bits(p)), uintptr(math.Float64bits(i)), uintptr(math.Float64bits(d))); code != 0 {
		panic(errors.New("MagneticInit call failed"))
	}
}

// 加速度算法初始化
func SpeedInit(speed, deltaTime float64) {
	if code, _, _ := _SpeedInit.Call(uintptr(math.Float64bits(speed)), uintptr(math.Float64bits(deltaTime))); code != 0 {
		panic(errors.New("SpeedInit call failed"))
	}
}

// 算法
func Compute(name string, currentX, currentY, targetX, targetY float64) (outputX, outputY float64) {
	if code, _, _ := utils.NewProc(name).Call(
		uintptr(math.Float64bits(currentX)),
		uintptr(math.Float64bits(currentY)),
		uintptr(math.Float64bits(float64(targetX))),
		uintptr(math.Float64bits(float64(targetY))),
		uintptr(unsafe.Pointer(&outputX)),
		uintptr(unsafe.Pointer(&outputY)),
	); code != 0 {
		panic(errors.New("Compute call failed"))
	}
	return
}

// 捕获初始化
func CaptureInit(hwnd, x, y, width, height int32) {
	if code, _, _ := _CaptureInit.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
	); code != 0 {
		panic(errors.New("CaptureInit call failed"))
	}
}

// 捕获释放
func CaptureRelease() {
	if code, _, _ := _CaptureRelease.Call(); code != 0 {
		panic(errors.New("CaptureRelease call failed"))
	}
}

// 捕获bitmap
func CaptureBitmap() (data []byte) {
	if code, _, _ := _CaptureBitmap.Call(uintptr(unsafe.Pointer(&data))); code != 0 {
		panic(errors.New("CaptureBitmap call failed"))
	}
	return
}

// 捕获bmp
func CaptureBmp() (data []byte) {
	if code, _, _ := _CaptureBmp.Call(uintptr(unsafe.Pointer(&data))); code != 0 {
		panic(errors.New("CaptureBmp call failed"))
	}
	return
}

// 通过窗口标题获取窗口句柄
func GetHwndByTitle(title string) (hwnd int32) {
	if code, _, _ := _GetHwndByTitle.Call(uintptr(unsafe.Pointer(&title)), uintptr(unsafe.Pointer(&hwnd))); code != 0 {
		panic(errors.New("GetHwndByTitle call failed"))
	}
	return
}

type Rect struct {
	Left, Top, Right, Bottom int32
}

// 通过窗口句柄获取窗口RECT
func GetWindowRect(hwnd int32) (rect Rect) {
	if code, _, _ := _GetWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect.Left)), uintptr(unsafe.Pointer(&rect.Top)), uintptr(unsafe.Pointer(&rect.Right)), uintptr(unsafe.Pointer(&rect.Bottom))); code != 0 {
		panic(errors.New("GetWindowRect call failed"))
	}
	return
}

// 通过窗口句柄判断窗口是否最小化
func IsWindowIconic(hwnd int32) bool {
	if code, _, _ := _IsWindowIconic.Call(uintptr(hwnd)); code == 0 {
		return true
	}
	return false
}

// 通过窗口句柄判断窗口是否最小化
func IsWindowVisible(hwnd int32) bool {
	if code, _, _ := _IsWindowVisible.Call(uintptr(hwnd)); code == 0 {
		return true
	}
	return false
}
