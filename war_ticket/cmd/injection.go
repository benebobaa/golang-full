package main

import (
	"war_ticket/internal/handler"
	gin_handler "war_ticket/internal/handler/gin"
	"war_ticket/internal/repository"
)

type Inject struct {
	eh  handler.EventHandler
	th  handler.TicketHandler
	oh  handler.OrderHandler
	geh gin_handler.EventHandler
	gth gin_handler.TicketHandler
	goh gin_handler.OrderHandler
	ur  repository.UserRepository
}
