pub:
	rm -rf service/internal/rpc/*
	go generate && go build
	mv digger /home/gopath/bin/
	cp web/recommend_list.html /data/www/index.html
	supervisorctl restart all
	supervisorctl status

regen:
	rm -rf service/internal/rpc/*.pb.go
	rm -rf service/internal/model/*
	go generate

prober:
	go run main.go prober
