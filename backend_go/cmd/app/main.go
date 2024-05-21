package main

import (
	"backend_go/internal/entity"
	handler "backend_go/internal/handler/http"
	"backend_go/internal/pkg/auth"
	"backend_go/internal/repository"
	repo_sqlite "backend_go/internal/repository/sqlite"
	"backend_go/internal/service"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func checkEntities() {
	fmt.Println("###### Checking entity")
	u := entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Email:    "senatorov.pp@phystech.edu",
				Password: "udsifpisfdb",
			},
			UserName: "PETEPEtrek",
		},
	}

	fmt.Printf("%+v\n", u)

	c := entity.Comment{
		CreatedAt: time.Now(),
		UserID:    1,
		IsGame:    true,
	}

	fmt.Printf("%+v\n", c)

	g := entity.Game{
		Name:       "Mass Effect",
		Synopsis:   "Was created in 2007",
		Image:      "http://mass_efeect.jpeg",
		Date:       time.Date(2007, 01, 01, 0, 0, 0, 0, time.UTC),
		Tags:       "/Action_RPG/Aliens/RPG/Shooter/",
		Score:      8.92,
		VoteNumber: 123,
	}
	fmt.Printf("%+v\n", g)
	fmt.Println("###### Ended checking entity")
	fmt.Println()
}

func checkRepository(repo *repository.Repository) {

	repo.Games.Create(
		&entity.Game{
			Name:       "Spider man (2018)",
			Synopsis:   "new game about spider man",
			Image:      "https://assets1.ignimgs.com/2018/09/04/ps4spider-man-blogroll-01-1536034979782_160w.jpg?width=1280",
			Date:       time.Date(2018, 9, 7, 0, 0, 0, 0, time.UTC),
			Tags:       "/Adventure/Hero/",
			Score:      8.78,
			VoteNumber: 423,
		},
	)

	repo.Games.Create(
		&entity.Game{
			Name:       "Assasin's creed",
			Synopsis:   "first game about assasin",
			Image:      "https://staticctf.ubisoft.com/J3yJr34U2pZ2Ieem48Dwy9uqj5PNUQTn/3T6D3ofZkhKmk9C5vb9IBV/b12d40a3e907d82598d0923ba8389c9f/ac-1.jpg?imwidth=360",
			Date:       time.Date(2007, 5, 2, 0, 0, 0, 0, time.UTC),
			Tags:       "/Adventure/",
			Score:      9.31,
			VoteNumber: 652,
		},
	)

	repo.Games.Create(
		&entity.Game{
			Name:       "Mafia 3",
			Synopsis:   "new game about mafia",
			Image:      "https://upload.wikimedia.org/wikipedia/ru/9/9f/Mafia_III_cover_art.jpg",
			Date:       time.Date(2016, 4, 12, 0, 0, 0, 0, time.UTC),
			Tags:       "/Shooter/",
			Score:      8.37,
			VoteNumber: 87,
		},
	)

	repo.Games.Create(
		&entity.Game{
			Name:       "The last of us: Part 2",
			Synopsis:   "new game about zombies",
			Image:      "https://upload.wikimedia.org/wikipedia/ru/thumb/f/f1/The_last_of_us_2_cover.png/274px-The_last_of_us_2_cover.png",
			Date:       time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC),
			Tags:       "/Shooter/Zombies/",
			Score:      9.45,
			VoteNumber: 1254,
		},
	)

	repo.Characters.Create(
		&entity.Character{
			Name:       "Commander Shepard",
			Story:      "Main Protagonist of ME, ME2 and ME3",
			Image:      "https://upload.wikimedia.org/wikipedia/en/d/df/Commander_Shepard.png",
			BirthDate:  time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC),
			Tags:       "/Hero/Human/",
			Score:      9.20,
			VoteNumber: 1521,
		},
	)

	repo.Characters.Create(
		&entity.Character{
			Name:       "Peter Parker/Spider Man",
			Story:      "Main Protagonist of many Spider Man games",
			Image:      "https://static.wikia.nocookie.net/spidermanps4/images/d/d9/IMG_4563.PNG/revision/latest/thumbnail/width/360/height/360?cb=20230721010243",
			BirthDate:  time.Date(1943, 2, 12, 0, 0, 0, 0, time.UTC),
			Tags:       "/Hero/Human/",
			Score:      8.89,
			VoteNumber: 256,
		},
	)
	repo.Characters.Create(
		&entity.Character{
			Name:       "Doctor Octopus",
			Story:      "Main Atagonist of Spider Man(2018)",
			Image:      "https://static.wikia.nocookie.net/spidermanps4/images/4/4d/Doctor_Octopus_from_MSM_concept_art_2.jpg/revision/latest?cb=20181001153130",
			BirthDate:  time.Date(1946, 5, 13, 0, 0, 0, 0, time.UTC),
			Tags:       "/Enemy/Human/",
			Score:      9.20,
			VoteNumber: 621,
		},
	)
	repo.Characters.Create(
		&entity.Character{
			Name:       "Mordin Solus",
			Story:      "Sidekick in ME, ME2 and ME3",
			Image:      "https://static.wikia.nocookie.net/thegamesrp/images/3/3e/Mordin.png/revision/latest?cb=20141124042447",
			BirthDate:  time.Date(2098, 8, 23, 0, 0, 0, 0, time.UTC),
			Tags:       "/Alien/Sidekick/",
			Score:      9.00,
			VoteNumber: 879,
		},
	)

	repo.People.Create(
		&entity.People{
			Name:       "Phil Spencer",
			Story:      "Head of Microsoft Games",
			Image:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSaRpSwV_XRC4MTT7eT2dvV_H9tq22od5GZmA&usqp=CAU",
			BirthDate:  time.Date(1968, 1, 12, 0, 0, 0, 0, time.UTC),
			Tags:       "/Microsoft/",
			Score:      8.50,
			VoteNumber: 842,
		},
	)

}

func checkService(service *service.Service) {
	fmt.Println("###### Checking service")

	fmt.Println("Try to register an user:")
	userReg := entity.UserRegister{
		UserLogin: entity.UserLogin{Email: "ololo@mail.ru", Password: "weiier"},
		UserName:  "Ivan_Ivanov",
	}
	err := service.User.Register(&userReg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}

	fmt.Println("Try to get User with ID=3:")
	user, err := service.User.Get(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *user)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Getting info about user (ID=3):")
	user, err = service.User.Get(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *user)
	}

	fmt.Println("Try to get Game with ID=1:")
	game, err := service.Game.Get(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *game)
	}

	fmt.Println("Try to change score of Game with ID=1:")
	err = service.Game.ChangeScore(1, 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Score Changed successfully")
	}

	fmt.Println("###### Ended checking service")
	fmt.Println()
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logFile, err := os.OpenFile("backend.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	db_uri, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get DB_URI from ENV")
		db_uri = "test.db"
	}

	db, err := repo_sqlite.NewSQLiteDB(db_uri)
	if err != nil {
		log.Panicf("Failed to initialize database: %s\n", err.Error())
	} else {
		log.Println("Database is initialized")
	}

	repo := repository.NewRepository(db)
	checkRepository(repo)
	newService := service.NewService(repo)

	signingKey, ok := os.LookupEnv("AUTH_SIGNING_KEY")
	if !ok {
		log.Println("cannot get AUTH_SIGNING_KEY from ENV")
		signingKey = "siuefui4nfweu"
	}
	authManager := auth.NewAuthManager([]byte(signingKey))

	h := handler.NewHandler(newService, authManager)

	srv := &http.Server{
		Addr: ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
