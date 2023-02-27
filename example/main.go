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
	log.SetFlags(log.Ltime | log.Lmsgprefix | log.Llongfile | log.Lshortfile)
	TryCatch(
		func() {
			var hwnd int32

			// 启动线程监控窗口
			go func() {
				for {
					newHwnd := utils.GetHwndByTitle("v2rayN - V3.29 - 2020/12/05")
					if hwnd != newHwnd {
						hwnd = newHwnd
						utils.CaptureInit(hwnd, 0, 0, 640, 640)
					}
					time.Sleep(1 * time.Second)
				}
			}()

			defer utils.CaptureRelease()

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
				// 判断窗口是否存在
				if hwnd <= 0 || utils.IsWindowIconic(hwnd) || !utils.IsWindowVisible(hwnd) {
					log.Println("窗口不存在或最小化或不可见")
					time.Sleep(1 * time.Second)
					continue
				}
				TryCatch(
					func() {
						// 截图
						startTime := time.Now()
						utils.CaptureBmp()
						captureTime := time.Since(startTime)

						// 获取当前鼠标位置
						currentX, currentY := utils.GetCursorPos()

						// 调用算法
						outputX, outputY := utils.Compute("LinearCompute", currentX, currentY, float64(centerX), float64(centerY))

						// 鼠标移动
						// utils.MoveMouse(outputX, outputY)

						log.Printf("当前位置：(%2d,%2d) - 截图耗时：%5v\n", int(outputX), int(outputY), captureTime)
						// os.WriteFile("output.bmp", b, fs.ModePerm)
					},
					func(err any) {
						log.Println(err)
						time.Sleep(1 * time.Second)
					},
				)

			}
		},
		func(err any) {
			log.Println(err)
		},
	)

}
