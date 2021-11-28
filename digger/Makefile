pub:
	rm -rf rpc/digger/*
	go generate && go build
	mv digger /home/gopath/bin/
	supervisorctl restart digger_prober
	supervisorctl status digger_prober

regen:
	rm -rf rpc/digger/*.pb.go
	rm -rf service/internal/model/*
	go generate

prober:
	go run main.go prober