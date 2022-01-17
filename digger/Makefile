all:
	clean, regen, pub

pub:
	rm -rf service/internal/rpc/*
	go generate && go build
	mv digger /home/gopath/bin/
	cp web/recommend_list.html /data/www/temp/index.html
	cp -rf ../zwadmin/dist/* /data/www/backend
	supervisorctl restart all
	supervisorctl status

clean:
	git checkout service/*

regen:
	rm -rf service/internal/rpc/*.pb.go
	rm -rf service/internal/model/*
	go generate

prober:
	go run main.go prober
