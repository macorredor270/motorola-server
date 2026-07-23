package services

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemInfo struct {
	Host    string
	OS      string
	CPU     float64
	RAM     float64
	Battery int
}

func GetSystem() SystemInfo {
	c, _ := cpu.Percent(0, false)
	r, _ := mem.VirtualMemory()
	battery := getBattery()

	return SystemInfo{
		Host:    "guam",
		OS:      runtime.GOOS,
		CPU:     c[0],
		RAM:     r.UsedPercent,
		Battery: battery,
	}
}

func getBattery() int {
	batteries, _ := host.BatteriesPercent()
	if len(batteries) > 0 {
		return int(batteries[0])
	}
	return 0
}
