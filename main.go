package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

type h map[string]interface{}

type Bit struct {
	ID  int64 `json:"id"`
	Bit int   `json:"bit"`
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("open file failed %v\n", err)
	}
	logger := log.New(f, "[bitmap] ", 2)
	cfg, err := ini.Load(os.Args[1])
	if err != nil {
		logger.Fatalf("reed config.ini failed %v", err)
	}
	var key = cfg.Section("redis").Key("key").String()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Section("redis").Key("addr").String(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect failed %v\n", err)
	}

	handler := func(w http.ResponseWriter, req *http.Request) {
		b, _ := ioutil.ReadAll(req.Body)
		logger.Printf("%s, %s, %s\n", req.Method, req.URL.String(), string(b))
		if req.Method == http.MethodGet {
			ids := req.URL.Query().Get("id")
			id, _ := strconv.Atoi(ids)
			if id == 0 {
				handleErr(w, errors.New("id out of range"))
				return
			}

			bit, err := redisClient.GetBit(key, int64(id)).Result()
			log.Println(bit)
			if err != nil {
				handleErr(w, err)
				return
			}

			b, _ := json.Marshal(h{
				"message": "sucess",
				"bit":     Bit{int64(id), int(bit)},
			})

			w.WriteHeader(200)
			w.Write(b)

		} else if req.Method == http.MethodPut {
			if len(b) == 0 {
				handleErr(w, errors.New("content is empty"))
				return
			}

			bit := &Bit{}
			err := json.Unmarshal(b, bit)
			if err != nil {
				handleErr(w, err)
				return
			}

			err = redisClient.SetBit(key, bit.ID, bit.Bit).Err()
			if err != nil {
				handleErr(w, err)
				return
			}

			b, _ = json.Marshal(h{
				"message": "sucess",
				"bit":     bit,
			})

			w.WriteHeader(200)
			w.Write(b)

		} else {
			w.WriteHeader(404)
			io.WriteString(w, "not found")
		}
	}

	http.HandleFunc("/bitmap", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleErr(w http.ResponseWriter, err error) {
	if err != nil {
		b, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.WriteHeader(400)
		w.Write(b)
	}
}
