package cases

import (
	"getCovid/sqlqueries"

	"github.com/gin-gonic/gin"
)

func QueryBuild(c *gin.Context) string {
	location := c.Query("statecode")
	gender := c.Query("gender")
	// age := c.Query("agebracket")
	status := c.Query("currentstatus")

	locationQueryEnd, genderQueryEnd, statusQueryEnd, queryEnd := "", "", "", ""

	if location != "" {

		locationQueryEnd = " statecode = $1 "
		queryEnd = " where " + locationQueryEnd
	}

	if gender != "" {
		genderQueryEnd = " gender "

		if queryEnd != "" {
			queryEnd = queryEnd + " AND " + genderQueryEnd + " = $2 "
		} else {
			queryEnd = " where " + genderQueryEnd + " = $1"
		}
	}

	if status != "" {
		statusQueryEnd = " currentstatus  "

		if queryEnd != "" {
			if locationQueryEnd != "" && genderQueryEnd != "" {
				queryEnd = queryEnd + " AND " + statusQueryEnd + " = $3"
			} else {
				queryEnd = queryEnd + " AND " + statusQueryEnd + " = $2"
			}

		} else {
			queryEnd = " where " + statusQueryEnd + "= $1"
		}

	}

	query := sqlqueries.GetLastTwoHundredCases + queryEnd + " limit 200"

	// if age != "" {
	// 	ageQueryEnd := "where statecode is " + age
	// }

	return query

}
