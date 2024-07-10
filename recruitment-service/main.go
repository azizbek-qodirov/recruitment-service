package main

// import (
// 	"fmt"
// 	"net/http"
// )

import (
	"log"
	"net"
	cfg "recruitment/config"
	er "recruitment/genprotos"
	src "recruitment/services"
	post "recruitment/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := post.DBCon()
	if err != nil {
		log.Fatal(err)
	}

	// config := cfg.Load()
	liss, err := net.Listen("tcp", cfg.Load().HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	er.RegisterCompanyServiceServer(server, src.NewCompanyService(db))
	er.RegisterJobServiceServer(server, src.NewJobService(db))
	er.RegisterVacancyServiceServer(server, src.NewVacancyService(db))

	log.Println("Server is running on port :5050")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}


// func Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!") // write response
// }

// func main() {
// 	http.HandleFunc("/", Handler)
// 	http.ListenAndServe(":7070", nil)
// }
