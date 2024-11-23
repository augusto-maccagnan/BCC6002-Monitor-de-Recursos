package main

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/resource"
	"net/http"
	"strconv"
	"strings"
)

func main() {

 // Create a new request multiplexer
 // Take incoming requests and dispatch them to the matching handlers
 mux := http.NewServeMux()

//  r := 

 // Register the routes and handlers
 mux.Handle("/", &homeHandler{})
 mux.Handle("/resources", &resourcesHandler{})

 // Run the server
 http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//  w.Write([]byte(getResources()))
    w.Write([]byte("This is my home page"))
}

type resourcesHandler struct{}

func (h *resourcesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // w.Write([]byte("This is my resources page"))
    switch {
    case r.Method == http.MethodGet:
        h.listResources(w, r)
        return
    }
}

func (h *resourcesHandler) listResources(w http.ResponseWriter, r *http.Request){
	dataText, err := resource.GetResources()
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := convertData(dataText)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 Not Found"))
}

func convertData(dataMap map[string][]string) ([]byte, error) {
	var data []model.Resources
	var CpuParams []model.CpuParams
	var CpuCores []model.CpuCores
	var Memory []model.Memory
	var TotalFreq int

	tmp := dataMap["CPU Max Frequency"]
	before, after, _ := strings.Cut(tmp[0], ":")
	idx := strings.Index(after, ":")
	after = after[idx+2:strings.Index(after, "\n")]
	MaxFreq, _ := strconv.Atoi(after[:7])
	after = after[:4] + "." + after[4:]
	CpuParams = append(CpuParams, model.CpuParams{Name: before, Frequency: after})

	tmp = dataMap["CPU Min Frequency"]
	before, after, _ = strings.Cut(tmp[0], ":")
	idx = strings.Index(after, ":")
	after = after[idx+2:strings.Index(after, "\n")]
	MinFreq, _ := strconv.Atoi(after[:7])
	after = after[:4] + "." + after[4:]
	CpuParams = append(CpuParams, model.CpuParams{Name: before, Frequency: after})

	tmp = dataMap["CPU Frequency"]
	for i := range tmp{
		before, after, _ = strings.Cut(tmp[i], ":")
		idx := strings.Index(after, ":")
		after = after[idx+2:]
		TempFreq, _ := strconv.Atoi(after[:4] + after[5:])
		CpuCores = append(CpuCores, model.CpuCores{Name: before, Frequency: after, Percentage: calculatePercentage(TempFreq, MaxFreq, MinFreq)})
		TotalFreq += TempFreq
	}

	tmp = dataMap["Total Memory"]
	before, after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	after = after[strings.LastIndex(after, " ")+1:]
	Memory = append(Memory, model.Memory{Name: before, Quantity: after})

	tmp = dataMap["Free Memory"]
	before, after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	after = after[strings.LastIndex(after, " ")+1:]
	Memory = append(Memory, model.Memory{Name: before, Quantity: after})

	tmp = dataMap["Available Memory"]
	before, after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	after = after[strings.LastIndex(after, " ")+1:]
	Memory = append(Memory, model.Memory{Name: before, Quantity: after})

	data = append(data, model.Resources{Cpu: CpuParams, CpuCores: CpuCores, Memory: Memory})

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("error marshalling data: %v", err)
		return nil, err
	}
	return jsonBytes, nil
}

func calculatePercentage(frequency int, maxFreq int, minFreq int) string {
	return strconv.FormatFloat((float64(frequency-minFreq)/float64(maxFreq-minFreq)*100), 'f', 2, 64)
}