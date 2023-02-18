package main

import (
	"log"
	"time"

	"github.com/aishuaikang/go-utils"
)

func TryCatch(callback func(), catch func(err any)) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	callback()
}

func main() {
	TryCatch(
		func() {
			// 获取窗口句柄
			hwnd := utils.GetHwndByTitle("AJ64插件绑定测试工具VIP专用 插件版本号:22.9.1")

			// 计算屏幕中心点
			width, height := utils.GetScreenResolution()
			centerX := int(width / 2)
			centerY := int(height / 2)

			// 初始化算法（可以重复调用更新参数）
			utils.SpeedInit(100, 0.2)
			utils.PidInit(0.1, 0.002, 0.003)
			utils.LinearInit(0.1)
			utils.MagneticInit(0.3)

			// for循环模拟推理
			for {
				// 截图
				startTime := time.Now().UnixMilli()
				utils.GetWindowCapture(hwnd, 0, 0, 160, 160, "bmp")
				captureTime := time.Now().UnixMilli() - startTime

				// 获取当前鼠标位置
				currentX, currentY := utils.GetCursorPos()

				// 调用算法
				outputX, outputY := utils.Compute("LinearCompute", currentX, currentY, float64(centerX), float64(centerY))

				// 鼠标移动
				utils.MoveMouse(outputX, outputY)

				log.Printf("当前位置：(%d,%d) - 截图耗时：(%dms)\n", int(outputX), int(outputY), captureTime)
			}
		},
		func(err any) {
			log.Println(err)
		},
	)

}
