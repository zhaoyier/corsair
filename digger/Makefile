pub:
	rm -rf rpc/digger/*
	go generate && go build
	mv digger /home/gopath/bin/
	cp web/hold_recom.html /data/www/index.html
	supervisorctl restart all
	supervisorctl status

regen:
	rm -rf rpc/digger/*.pb.go
	rm -rf service/internal/model/*
	go generate

prober:
	go run main.go prober