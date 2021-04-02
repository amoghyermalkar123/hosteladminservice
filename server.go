package main

import (
	"admin/grpc/services"
	"admin/proto/api"
	"fmt"
	"net"

	"admin/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &handler.Handler{}
	// Routes
	e.POST("/register/", h.RegisterNewStudent)
	e.POST("/message_student/", h.CreateStudentMessage)
	go grpcServerStarter()

	// Start server
	e.Logger.Fatal(e.Start(":1320"))
}

func grpcServerStarter() {
	// listen
	listener, err := net.Listen("tcp", ":3032")
	if err != nil {
		panic(err)
	}

	serverObject := grpc.NewServer()
	fmt.Println("server")

	api.RegisterAdminMessageServiceServer(serverObject, &services.AdminService{})
	reflection.Register(serverObject)

	// Serve
	if e := serverObject.Serve(listener); e != nil {
		panic(e)
	}
}
