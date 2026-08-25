package api

import "ticketverify/internal/audit"

func (s *Service) AuditEntries(actor string) []audit.Entry { return s.Audit.Query(actor) }
