package main

import "fmt"

type Job struct {
    Title string
    Salary int
	Company string

}

func (j Job) describe() {
	fmt.Printf("Role: %s at %s - %d LPA\n", j.Title, j.Company, j.Salary)

}

type Opportunity interface {
	describe()
}

func printOpportunity(o Opportunity) {
    o.describe()
}

func main() {
	job := Job{
		Title: "Backend Engineer",
		Salary: 8,
		Company: "Zerodha",
		}
		printOpportunity(job)
	}