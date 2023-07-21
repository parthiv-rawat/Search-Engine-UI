package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type SearchResult struct {
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Established int      `json:"established"`
	Website     string   `json:"website"`
	Programs    []string `json:"programs"`
	Tuition     struct {
		InState    string `json:"Cost-in-state"`
		OutOfState string `json:"Cost-out-of-state"`
	} `json:"tuition"`
	Admission struct {
		AcceptanceRate string `json:"acceptance_rate"`
		SATScoreRange  string `json:"SAT_score_range"`
		ACTScoreRange  string `json:"ACT_score_range"`
	} `json:"admission"`
}

func searchResults(query string, results []SearchResult) any {
	var searchResults []SearchResult

	if query != "" {
		for _, result := range results {
			if containsQueryInSearchResult(result, query) {
				searchResults = append(searchResults, result)
			}
		}

		// fmt.Println(searchResults)
		if searchResults != nil {
			return searchResults
		}
	}
	return ""
}

func containsQueryInSearchResult(result SearchResult, query string) bool {
	// Check if the query is present in the Tuition fields (InState and OutOfState).
	if strings.Contains(strings.ToLower(result.Tuition.InState), strings.ToLower(query)) ||
		strings.Contains(strings.ToLower(result.Tuition.OutOfState), strings.ToLower(query)) {
		return true
	}

	// Check if the query is present in the Admission fields (AcceptanceRate, SATScoreRange, and ACTScoreRange).
	if strings.Contains(strings.ToLower(result.Admission.AcceptanceRate), strings.ToLower(query)) ||
		strings.Contains(strings.ToLower(result.Admission.SATScoreRange), strings.ToLower(query)) ||
		strings.Contains(strings.ToLower(result.Admission.ACTScoreRange), strings.ToLower(query)) {
		return true
	}

	// Check if the query is present in the Name, Location, and Programs fields.
	if strings.Contains(strings.ToLower(result.Name), strings.ToLower(query)) ||
		strings.Contains(strings.ToLower(result.Location), strings.ToLower(query)) ||
		containsProgram(result.Programs, query) {
		return true
	}

	return false
}

func containsProgram(programs []string, query string) bool {
	for _, program := range programs {
		if strings.Contains(strings.ToLower(program), strings.ToLower(query)) {
			return true
		}
	}
	return false
}

func main() {

	// results := []SearchResult{
	// 	{
	// 		Title:       "Harward University",
	// 		Description: "This is the official website of Harward University",
	// 		URL:         "https://www.harvard.edu/",
	// 	},
	// 	{
	// 		Title:       "Stanford University",
	// 		Description: "This is the official website of Stanford University",
	// 		URL:         "https://www.stanford.edu/",
	// 	},
	// 	{
	// 		Title:       "University of Oxford",
	// 		Description: "This is the official website of Oxford University",
	// 		URL:         "https://www.ox.ac.uk/",
	// 	},
	// 	{
	// 		Title:       "Massachusetts Institute of Technology",
	// 		Description: "This is the official website of MIT",
	// 		URL:         "https://www.mit.edu/",
	// 	},
	// 	{
	// 		Title:       "University of Cambridge",
	// 		Description: "This is the official website of Cambridge University",
	// 		URL:         "https://www.cam.ac.uk/",
	// 	},
	// 	{
	// 		Title:       "National University of Singapore",
	// 		Description: "This is the official website of NUS",
	// 		URL:         "https://nus.edu.sg/",
	// 	},
	// 	{
	// 		Title:       "Princeton University",
	// 		Description: "This is the official website of Princeton University",
	// 		URL:         "https://www.princeton.edu/",
	// 	},
	// }

	results := []SearchResult{
		{
			Name:        "Harvard University",
			Location:    "Cambridge, MA",
			Established: 1636,
			Website:     "https://www.harvard.edu/",
			Programs: []string{
				"Computer Science",
				"Economics",
				"Political Science",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$51,925",
				OutOfState: "$51,925",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "4.7%",
				SATScoreRange:  "1460-1570",
				ACTScoreRange:  "33-35",
			},
		},
		{
			Name:        "Stanford University",
			Location:    "Stanford, CA",
			Established: 1885,
			Website:     "https://www.stanford.edu/",
			Programs: []string{
				"Engineering",
				"Business",
				"Biology",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$56,169",
				OutOfState: "$56,169",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "4.3%",
				SATScoreRange:  "1440-1570",
				ACTScoreRange:  "32-35",
			},
		},
		{
			Name:        "Massachusetts Institute of Technology (MIT)",
			Location:    "Cambridge, MA",
			Established: 1861,
			Website:     "https://www.mit.edu/",
			Programs: []string{
				"Computer Science",
				"Engineering",
				"Physics",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$53,450",
				OutOfState: "$53,450",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "6.7%",
				SATScoreRange:  "1490-1570",
				ACTScoreRange:  "34-36",
			},
		},
		{
			Name:        "California Institute of Technology (Caltech)",
			Location:    "Pasadena, CA",
			Established: 1891,
			Website:     "https://www.caltech.edu/",
			Programs: []string{
				"Physics",
				"Chemistry",
				"Biology",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$54,600",
				OutOfState: "$54,600",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "6.4%",
				SATScoreRange:  "1530-1580",
				ACTScoreRange:  "35-36",
			},
		},
		{
			Name:        "University of Oxford",
			Location:    "Oxford, England",
			Established: 1096,
			Website:     "https://www.ox.ac.uk/",
			Programs: []string{
				"Mathematics",
				"History",
				"Medicine",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "£9,250",
				OutOfState: "£37,000",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "17.5%",
				SATScoreRange:  "N/A",
				ACTScoreRange:  "N/A",
			},
		},
		{
			Name:        "Cambridge University",
			Location:    "Cambridge, England",
			Established: 1209,
			Website:     "https://www.cam.ac.uk/",
			Programs: []string{
				"Computer Science",
				"Engineering",
				"Psychology",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "£9,250",
				OutOfState: "£37,000",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "21%",
				SATScoreRange:  "N/A",
				ACTScoreRange:  "N/A",
			},
		},
		{
			Name:        "Princeton University",
			Location:    "Princeton, NJ",
			Established: 1746,
			Website:     "https://www.princeton.edu/",
			Programs: []string{
				"Physics",
				"Economics",
				"Psychology",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$53,757",
				OutOfState: "$53,757",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "5.8%",
				SATScoreRange:  "1460-1570",
				ACTScoreRange:  "32-35",
			},
		},
		{
			Name:        "Yale University",
			Location:    "New Haven, CT",
			Established: 1701,
			Website:     "https://www.yale.edu/",
			Programs: []string{
				"English",
				"Political Science",
				"Biology",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$57,700",
				OutOfState: "$57,700",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "4.6%",
				SATScoreRange:  "1460-1570",
				ACTScoreRange:  "33-35",
			},
		},
		{
			Name:        "University of Cambridge",
			Location:    "Cambridge, England",
			Established: 1209,
			Website:     "https://www.cam.ac.uk/",
			Programs: []string{
				"Computer Science",
				"Mathematics",
				"Chemistry",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "£9,250",
				OutOfState: "£37,000",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "21%",
				SATScoreRange:  "N/A",
				ACTScoreRange:  "N/A",
			},
		},
		{
			Name:        "Columbia University",
			Location:    "New York, NY",
			Established: 1754,
			Website:     "https://www.columbia.edu/",
			Programs: []string{
				"Business",
				"Journalism",
				"Architecture",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$61,788",
				OutOfState: "$61,788",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "5.5%",
				SATScoreRange:  "1480-1570",
				ACTScoreRange:  "33-35",
			},
		},
		{
			Name:        "University of Chicago",
			Location:    "Chicago, IL",
			Established: 1890,
			Website:     "https://www.uchicago.edu/",
			Programs: []string{
				"Economics",
				"Law",
				"Political Science",
			},
			Tuition: struct {
				InState    string `json:"Cost-in-state"`
				OutOfState string `json:"Cost-out-of-state"`
			}{
				InState:    "$61,548",
				OutOfState: "$61,548",
			},
			Admission: struct {
				AcceptanceRate string `json:"acceptance_rate"`
				SATScoreRange  string `json:"SAT_score_range"`
				ACTScoreRange  string `json:"ACT_score_range"`
			}{
				AcceptanceRate: "6.2%",
				SATScoreRange:  "1490-1570",
				ACTScoreRange:  "33-35",
			},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println(r.URL.Query())
		query := r.URL.Query().Get("query")

		// fmt.Println(query)
		searchResults := searchResults(query, results)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(searchResults)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
