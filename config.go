package main

// Config data
type Config struct {
	Reference      string        `json:"validation_reference"`
	Period         string        `json:"validation_period"`
	Survey         string        `json:"validation_survey"`
	Periodicity    string        `json:"periodicity"`
	BpmID          string        `json:"bpmid"`
	Contributors   []Contributor `json:"contributor"`
	Responses      []Response    `json:"response"`
	QuestionSchema []struct {
		Period    string `json:"period"`
		Survey    string `json:"survey"`
		Question  string `json:"question"`
		Datatype  string `json:"datatype"`
		Repeating bool   `json:"repeating"`
	} `json:"question_schema"`
	Validations []struct {
		Template        string `json:"template"`
		Formula         string `json:"formula"`
		QuestionDetails []struct {
			ValidationID    int    `json:"validationid"`
			PrimaryQuestion string `json:"primary_question"`
			Default         string `json:"default"`
			OutputFormula   string
			Parameters      []ValidationParameter `json:"parameters"`
		} `json:"question_details"`
	} `json:"validation_config"`
}

// Contributor ...
type Contributor struct {
	Reference string `json:"reference"`
	Period    string `json:"period"`
	Survey    string `json:"survey"`
	Status    string `json:"status"`
	Frosic    string `json:"frosic"`
	Rusic     string `json:"rusic"`
}

// Response ...
type Response struct {
	Reference string `json:"reference"`
	Period    string `json:"period"`
	Survey    string `json:"survey"`
	Instance  int    `json:"instance"`
	Question  string `json:"question"`
	Response  string `json:"response"`
}

// ValidationParameter ...
type ValidationParameter struct {
	Name             string `json:"name"`
	Value            string `json:"value"`
	Source           string `json:"source"`
	ResponseOffset   int    `json:"response_offset"`
	OffsetPeriod     string
	ReplacementValue string
}
