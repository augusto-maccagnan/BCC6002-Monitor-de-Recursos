package resource

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetResources()  (map[string][]string, error) {
	var cpuCount int
	var fields []string
	var total []string
	var usage []string
	var free []string
	var name []string
	var resources = make(map[string][]string)

	cpuFrequency, cpuCount := getCpuUsage()
	cpuMaxFreq := getCpuMaxFreq()
	cpuMinFreq := getCpuMinFreq()
	memInfo := getMemInfo()
	diskInfo := getDiskInfo()


	for i := range cpuFrequency {
		fields = append(fields, "CPU " + strconv.Itoa(i) + " frequency: " + cpuFrequency[i])
	}
	resources["CPU Frequency"] = fields
	fields = nil
	fields = append(fields, "CPU Max Frequency: " + cpuMaxFreq)
	resources["CPU Max Frequency"] = fields
	fields = nil
	fields = append(fields, "CPU Min Frequency: " + cpuMinFreq)
	resources["CPU Min Frequency"] = fields
	fields = nil
	fields = append(fields, "Number of Physical Cores: " + strconv.Itoa(cpuCount))
	resources["Number of Physical Cores"] = fields
	
	fields = nil
	fields = append(fields, "Memory " + memInfo[0])
	resources["Total Memory"] = fields
	fields = nil
	fields = append(fields, "Memory " + memInfo[1])
	resources["Free Memory"] = fields
	fields = nil
	fields = append(fields, "Memory " + memInfo[2])
	resources["Available Memory"] = fields

	for i := range diskInfo {
		// fmt.Println(diskInfo[i])
		j := 0
		for diskInfo[i][j] == ' '{
			j++
		}
		diskInfo[i] = diskInfo[i][j:] // Corta até primeiro texto
		diskName := diskInfo[i][:strings.Index(diskInfo[i], " ")]
		diskInfo[i] = diskInfo[i][strings.Index(diskInfo[i], " "):] // Corta até o primeiro espaço
		j = 0
		for diskInfo[i][j] == ' '{
			j++
		}
		diskInfo[i] = diskInfo[i][j:] // Corta até primeiro texto
		diskInfo[i] = diskInfo[i][strings.Index(diskInfo[i], " "):] // Corta até o primeiro espaço
		j = 0
		for diskInfo[i][j] == ' '{
			j++
		}
		diskInfo[i] = diskInfo[i][j:] // Corta espaços iniciais
		diskTotal := diskInfo[i][:strings.Index(diskInfo[i], " ")]
		diskInfo[i] = diskInfo[i][strings.Index(diskInfo[i], " "):] // Corta até o primeiro espaço
		j = 0
		for diskInfo[i][j] == ' '{
			j++
		}
		diskInfo[i] = diskInfo[i][j:] // Corta espaços iniciais
		diskUsage := diskInfo[i][:strings.Index(diskInfo[i], " ")]
		diskInfo[i] = diskInfo[i][strings.Index(diskInfo[i], " "):] // Corta até o primeiro espaço
		j = 0
		for diskInfo[i][j] == ' '{
			j++
		}
		diskInfo[i] = diskInfo[i][j:] // Corta espaços iniciais
		diskFree := diskInfo[i][:strings.Index(diskInfo[i], " ")]
		diskInfo[i] = diskInfo[i][strings.Index(diskInfo[i], " "):] // Corta até o primeiro espaço

		name = append(name, strconv.Itoa(i) + ":" + diskName)
		total = append(total, strconv.Itoa(i) + ":" + diskTotal)
		usage = append(usage, strconv.Itoa(i) + ":" + diskUsage)
		free = append(free, strconv.Itoa(i) + ":" + diskFree)
	}
	resources["Disk Name"] = name
	resources["Disk Total"] = total
	resources["Disk Usage"] = usage
	resources["Disk Free"] = free

	return resources, nil
}

func getCpuMaxFreq() string {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq")
	cpuMaxFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(cpuMaxFreq)
}

func getCpuMinFreq() string {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq")
	cpuMinFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(cpuMinFreq)
}

func getCpuUsage() ([]string, int) {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	osOutput, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	text := string(osOutput)

	var cpuFrequency []string
	var found bool = true
	var i, index int
	var temp string
	for (found) { // Search for next line
		index = strings.Index(text, "cpu MHz")
		if(index == -1){
			break
		}
		text = text[index:]
		temp, text, found = strings.Cut(text, "\n")
		cpuFrequency = append(cpuFrequency, temp)
		i++
	}
	return cpuFrequency, i
}

func getMemInfo() []string {
	var memInfo []string

	cmd := exec.Command("cat", "/proc/meminfo")
	temp, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var found bool = true
	var after string = string(temp)

	for found {
		_, after, found = strings.Cut(after, "Mem")
		if(!found) {
			break
		}
		memInfo = append(memInfo, after[:strings.Index(after, "\n")])
	}

	return memInfo
}

func getDiskInfo() []string {
	var diskInfo []string
	var found bool = true
	var after, before string
	cmd := exec.Command("df", "-T")
	temp, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	after = string(temp)
	for found {
		before, after, found = strings.Cut(after, "\n")
		// if(strings.Contains(before, "ext4")) {
		if(before != "" && !(strings.Contains(before, "Tipo") || strings.Contains(before, "Type"))) {
			diskInfo = append(diskInfo, before)
		}
		// }
	}
	return diskInfo
}