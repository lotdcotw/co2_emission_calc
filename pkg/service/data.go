package service

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/farkow/co2e/pkg/api"
	"github.com/farkow/co2e/pkg/utils"
	"github.com/farkow/etcd/clientv3"
	etcd "github.com/farkow/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/ini.v1"
)

const (
	timeout        = 3 * time.Second
	staticDataFile = "./data/co2e.ini"
)

func getValues() {
	emissions = make(map[string]float64, 0)
	localAddress := "http://localhost:" + utils.Flags.EtcdPort // TODO clusters

	// etcd connection
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	cli, err := clientv3.New(etcd.Config{
		DialTimeout: timeout,
		Endpoints:   []string{localAddress},
	})
	// if connection is failed, get the defaults from ini file
	if err != nil {
		cancel()
		loadDefaults("")
		return
	}
	defer cli.Close()
	kv := etcd.NewKV(cli)

	// get all values one by one from etcd
	// TODO: get all at once
	for k := range transmMap {
		gr, err := kv.Get(ctx, k)
		if err != nil {
			log.Printf("[ETCD ERROR] %v\n", err)
			loadDefaults("")
			break
		}
		f, _ := strconv.ParseFloat(string(gr.Kvs[0].Value), 64)
		emissions[k] = f
	}

	cancel()
}

func validate(req *api.Request) error {
	// TODO validate and secure the request

	transportationMethod := keyConversion(req.TransportationMethod)
	if len(transportationMethod) == 0 {
		return status.Error(codes.InvalidArgument, "given transportation value is invalid")
	}
	(*req).TransportationMethod = transportationMethod

	return nil
}

// load default emission values from ini file
func loadDefaults(override string) {
	log.Println("Loading default emission values...")

	dataFile := staticDataFile
	if len(override) > 0 {
		log.Println("Overriding default path of INI file...")
		dataFile = override
	}

	var err error
	var filename = ""
	// check current and one above directory
	if _, err = os.Stat(dataFile); os.IsNotExist(err) {
		if _, err = os.Stat("." + dataFile); os.IsNotExist(err) {
			log.Println("Emission values are not found; all will be taken as 1")
			return
		}
		filename = "." + dataFile
	} else {
		filename = dataFile
	}

	cfg, err := ini.Load(filename)

	// map static values
	for k := range transmMap {
		kSplit := strings.Split(k, "_")
		section, err := cfg.GetSection(kSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		key, err := section.GetKey(kSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		val, err := strconv.ParseFloat(key.Value(), 64)
		if err != nil {
			log.Fatal(err)
		}
		emissions[k] = val
	}
}
