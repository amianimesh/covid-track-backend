package cases

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Covidcases struct {
	Patientnumber              string `json:"patientNumber"`
	Statepatientnumber         string `json:"statePatientNumber"`
	Dateannounced              string `json:"dateAnnounced"`
	Estimatedonsetdate         string `json:"estimatedOnSetdate"`
	Agebracket                 string `json:"ageBracket"`
	Gender                     string `json:"gender"`
	Detectedcity               string `json:"detectedCity"`
	Detecteddistrict           string `json:"detectedDistrict"`
	Detectedstate              string `json:"detectedState"`
	Statecode                  string `json:"stateCode"`
	Currentstatus              string `json:"currentStatus"`
	Notes                      string `json:"notes"`
	Contractedfromwhichpatient string `json:"contractedFromWhichPatient"`
	Nationality                string `json:"nationality"`
	Typeoftransmission         string `json:"typeOfTransmission"`
	Statuschangedate           string `json:"statusChangeDate"`
	Source1                    string `json:"source1"`
	Source2                    string `json:"source2"`
	Source3                    string `json:"source3"`
	Backupnotes                string `json:"backUpnNotes"`
	Numcases                   string `json:"numCases"`
}

var DatabaseConnection *sql.DB

func Getcovidcases(c *gin.Context) {

	operation, e := sql.Open("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if e != nil {
		log.Fatal("Error Opening Database: ", e)
	}

	DatabaseConnection = operation

	query := QueryBuild(c)

	location := c.Query("statecode")
	gender := c.Query("gender")
	// age := c.Query("agebracket")
	status := c.Query("currentstatus")
	var rows *sql.Rows
	var err error

	if location != "" && gender != "" && status != "" {
		rows, err = DatabaseConnection.Query(query, location, gender, status)
	} else if location != "" && gender != "" {
		rows, err = DatabaseConnection.Query(query, location, gender)
	} else if location != "" && status != "" {
		rows, err = DatabaseConnection.Query(query, location, status)
	} else if gender != "" && status != "" {
		rows, err = DatabaseConnection.Query(query, gender, status)

	} else if location != "" {
		rows, err = DatabaseConnection.Query(query, location)

	} else if gender != "" {
		rows, err = DatabaseConnection.Query(query, gender)

	} else if status != "" {
		rows, err = DatabaseConnection.Query(query, status)
	} else {
		rows, err = DatabaseConnection.Query(query)
	}

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error Getting data: %q %s", err, query))
		return
	}

	var cases []Covidcases

	defer rows.Close()
	for rows.Next() {
		var (
			patientnumber              sql.NullString
			statepatientnumber         sql.NullString
			dateannounced              sql.NullString
			estimatedonsetdate         sql.NullString
			agebracket                 sql.NullString
			gender                     sql.NullString
			detectedcity               sql.NullString
			detecteddistrict           sql.NullString
			detectedstate              sql.NullString
			statecode                  sql.NullString
			currentstatus              sql.NullString
			notes                      sql.NullString
			contractedfromwhichpatient sql.NullString
			nationality                sql.NullString
			typeoftransmission         sql.NullString
			statuschangedate           sql.NullString
			source1                    sql.NullString
			source2                    sql.NullString
			source3                    sql.NullString
			backupnotes                sql.NullString
			numcases                   sql.NullString
		)

		if err := rows.Scan(&patientnumber, &statepatientnumber, &dateannounced, &estimatedonsetdate,
			&agebracket, &gender, &detectedcity, &detecteddistrict, &detectedstate,
			&statecode, &currentstatus, &notes, &contractedfromwhichpatient,
			&nationality, &typeoftransmission, &statuschangedate, &source1,
			&source2, &source3, &backupnotes, &numcases); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error Scanning Database: %q", err))
			return
		}

		newcases := Covidcases{
			Patientnumber:              patientnumber.String,
			Statepatientnumber:         statepatientnumber.String,
			Dateannounced:              dateannounced.String,
			Estimatedonsetdate:         estimatedonsetdate.String,
			Agebracket:                 agebracket.String,
			Gender:                     gender.String,
			Detectedcity:               detectedcity.String,
			Detecteddistrict:           detecteddistrict.String,
			Detectedstate:              detectedstate.String,
			Statecode:                  statecode.String,
			Currentstatus:              currentstatus.String,
			Notes:                      notes.String,
			Contractedfromwhichpatient: contractedfromwhichpatient.String,
			Nationality:                nationality.String,
			Typeoftransmission:         typeoftransmission.String,
			Statuschangedate:           statuschangedate.String,
			Source1:                    source1.String,
			Source2:                    source2.String,
			Source3:                    source3.String,
			Backupnotes:                backupnotes.String,
			Numcases:                   numcases.String,
		}

		cases = append(cases, newcases)
	}
	c.JSON(http.StatusOK, cases)
}
