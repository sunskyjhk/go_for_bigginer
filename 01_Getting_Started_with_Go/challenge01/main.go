package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"os"
	"runtime"
	"strconv"
	"text/tabwriter"

	"github.com/shirou/gopsutil/disk"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
Utilities.
*/

func byteToMB(b uint64) uint64 {
	return b / 1024 / 1024
}

func printUsage(u *disk.UsageStat) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	if len(u.Path) > 18 {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Drive/Path", "Percentage Full", "Total", "Free", "Used")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\n",
			u.Path,
			strconv.FormatFloat(u.UsedPercent, 'f', 2, 64)+"%",
			strconv.FormatUint(u.Total/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Free/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Used/1024/1024/1024, 10)+" GB")
	} else {
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\t", "Drive/Path", "Percentage Full", "Total", "Free", "Used")
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\t",
			u.Path,
			strconv.FormatFloat(u.UsedPercent, 'f', 2, 64)+"%",
			strconv.FormatUint(u.Total/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Free/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Used/1024/1024/1024, 10)+" GB")
	}
}

/*
Functionality.
*/

// PrintCPUUtilization - Print CPU utilization.
func PrintCPUUtilization() {
	fmt.Println("\n\t\t CPU Utilization")

	//v, err := mem.VirtualMemory()
	cpuStat, err := cpu.Info()
	check(err)

	percentage, err := cpu.Percent(0, true)
	check(err)

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	for i := 0; i < len(cpuStat); i++ {
		fmt.Println("CPU index number: ", cpuStat[i].CPU)
		fmt.Println("VendorID: ", cpuStat[i].VendorID)
		fmt.Println("Family: ", cpuStat[i].Family)
		fmt.Println("Number of cores: ", cpuStat[i].Cores)
	}

	fmt.Println(cpuStat[0].CPU)
	fmt.Println(strconv.FormatFloat(percentage[0], 'f', 2, 64))

	fmt.Fprintf(w, "\n %s\t%s\t", "core index", "Used Percent")
	fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
	for idx := 0; idx < len(percentage); idx++ {
		fmt.Fprintf(w, "\n %v\t%f\t", idx, percentage[idx])
	}
}

// PrintRAMUtilization = Print RAM utilization.
func PrintRAMUtilization() {
	fmt.Println("\n\t\t RAM Utilization")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n")
}

func main() {
	PrintCPUUtilization()
}
