package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"pg_stats/database"
	"pg_stats/models"
	"strings"
)


/*
  This is a Fiber handler that gets Slowest Queries from postgres database statistics
  IMPORTANT NOTICE:
       - this function tested on postgres version 13.3
       - pg_stat_statements module should be installed and enabled on postgresql
         if it's not, please read this article to fix it:
         https://www.postgresql.org/docs/13/pgstatstatements.html
 */


func StatsHandler(c *fiber.Ctx) error {

	// read all query params from request
	limit := c.Query("limit", "15")
	offset := c.Query("offset", "0")
	filter := strings.Trim(strings.ToLower(c.Query("filter", "")), " ")

	filterable := false
    var err error
	// check filter query param if is a valid input or not.
	// if it's valid then make query filterable
	switch filter {
		case "select":
			filterable = true
			break
		case "update":
			filterable = true
			break
		case "delete":
			filterable = true
			break
		case "insert":
			filterable = true
			break
	    case  "":
		    filterable = true
		    break
		default:
			return fiber.NewError(fiber.StatusBadRequest,"invalid filter value. valid values are: select, insert, update, delete")


	}

	//get database connection
	db := database.Db


	var rows *sql.Rows
	// make query statement end fetch data from database
	Stmt := `SELECT query, calls, total_exec_time
       FROM pg_stat_statements `
	if filterable {
		Stmt += `
       Where LOWER(query) like concat( LOWER($1) ,'%')
       ORDER BY total_exec_time 
	   DESC LIMIT $2 offset $3;`
		rows, err = db.Query(Stmt, filter, limit, offset)
	} else {
		Stmt += ` ORDER BY total_exec_time 
	   DESC LIMIT $1 offset $2;`
		rows, err = db.Query(Stmt, limit, offset)
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// read all rows fetched from database and append  data into result array
	var result []models.SatementStat
	for rows.Next() {
		row := models.SatementStat{}
		err := rows.Scan(&row.Query, &row.Calls, &row.TotalExecTime)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		result = append(result, row)
	}
	err = rows.Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(result)
}
