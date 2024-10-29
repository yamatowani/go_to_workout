package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
)

const (
	notionAPIURL    = "https://api.notion.com/v1/pages"
	notionAPIVersion = "2022-06-28"
)

var apiToken = os.Getenv("API_TOKEN")
var databaseID = os.Getenv("DATABASE_ID")


func addLog(exercise string, weight float64, sets int, reps []int, date string) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println("API Token:", apiToken)
fmt.Println("Database ID:", databaseID)
	payload := map[string]interface{}{
		"parent": map[string]string{
			"database_id": databaseID,
		},
		"properties": map[string]interface{}{
			"Exercise": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]string{"content": exercise},
					},
				},
			},
			"Weight": map[string]interface{}{
				"number": weight,
			},
			"Sets": map[string]interface{}{
				"number": sets,
			},
			"Reps": map[string]interface{}{
				"rich_text": []map[string]interface{}{
					{
						"text": map[string]string{"content": fmt.Sprint(reps)},
					},
				},
			},
			"Date": map[string]interface{}{
				"date": map[string]interface{}{
					"start": parsedDate.Format("2006-01-02"),
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", notionAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", notionAPIVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to Notion:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to add workout log to Notion:", resp.Status)
	} else {
		fmt.Println("Workout logged successfully to Notion")
	}
}
