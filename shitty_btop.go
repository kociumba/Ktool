package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/wzshiming/ctc"
)

func sysInfo() {

	for {
		v, _ := mem.VirtualMemory()
		o, _ := host.Info()
		c, _ := cpu.Info()

		memoryInfo := fmt.Sprintf("Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent)
		osInfo := fmt.Sprintf("OS: %v, Uptime: %v, Procs: %v", o.OS, o.Uptime, o.Procs)
		cpuInfo := fmt.Sprintf("Vendor: %v, Cores: %v, Mhz: %v, Model: %v", c[0].VendorID, c[0].Cores, c[0].Mhz, c[0].ModelName)

		fmt.Print("\033[2J")
		fmt.Print("\033[H")

		fmt.Println("<----------SYS INFO---------->")
		fmt.Println(ctc.ForegroundYellow, "System:", osInfo, ctc.Reset)
		fmt.Println(ctc.ForegroundBrightCyan, "Cpu:", cpuInfo, ctc.Reset)
		fmt.Println(ctc.ForegroundBrightGreenBackgroundBlack, "Memory:", memoryInfo, ctc.Reset)

		fmt.Println("<---------------------------->")

		time.Sleep(500 * time.Millisecond)
	}

}
