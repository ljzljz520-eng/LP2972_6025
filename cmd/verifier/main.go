package main

import (
	"fmt"
	"ticketverify/internal/api"
	"time"
)

func main() {
	s := api.New()
	s.Events.Create("event-1", "Final", "Arena", time.Now().Add(24*time.Hour), 1000)
	t, _ := s.Issue("ticket-1", "event-1", "holder-1", "standard")
	fmt.Println(ticketsCode(t), s.Validate([]string{t.ID}, 10))
}
func ticketsCode(t interface{}) string { return fmt.Sprintf("%v", t) }
