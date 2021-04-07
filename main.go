package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

var DatabaseConnection *sql.DB

//CommentParams type
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

func getcovidcases(c *gin.Context) {
	rows, err := DatabaseConnection.Query("SELECT patientnumber,statepatientnumber,dateannounced,estimatedonsetdate,agebracket,gender,detectedcity,detecteddistrict,detectedstate,statecode,currentstatus,notes,contractedfromwhichpatient,nationality,typeoftransmission,statuschangedate,source1,source2,source3,backupnotes,numcases FROM COVID limit 200")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error Getting data: %q", err))
		return
	}
	/*location := c.Query("statecode")
	gender := c.Query("gender")
	age := c.Query("agebracket")
	status := c.Query("currentstatus")*/

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

func main() {

	operation, err := sql.Open("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error Opening Database: %q", err)
	}

	DatabaseConnection = operation

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/patient", func(c *gin.Context) {
		location := c.Query("statecode")
		gender := c.Query("gender") // shortcut for c.Request.URL.Query().Get("lastname")
		age := c.Query("agebracket")
		status := c.Query("currentstatus")

		c.String(http.StatusOK, "Hello %s %s %s %s", location, gender, age, status)
	})

	router.GET("/cases", getcovidcases)

	router.Run()

}
