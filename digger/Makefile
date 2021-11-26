pub:
	rm -rf rpc/digger/*
	go generate && go build
	mv digger /home/gopath/bin/
	supervisorctl restart digger_prober
	supervisorctl status digger_prober