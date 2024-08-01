package main

import (
	"war_ticket/internal/handler"
	ginhandler "war_ticket/internal/handler/gin"
	"war_ticket/internal/repository"
)

type Inject struct {
	eh  handler.EventHandler
	th  handler.TicketHandler
	oh  handler.OrderHandler
	geh ginhandler.EventHandler
	gth ginhandler.TicketHandler
	goh ginhandler.OrderHandler
	ur  repository.UserRepository
}
