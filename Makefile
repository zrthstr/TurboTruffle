### test / build / dev

build:
	go build gen_html.go

go_test:
	go run ./gen_html.go ./results/osin-storage.trufflog ./results/osin-storage.trufflog.html highlighter.js


### usage

run:
	bash test_token.sh
	bash get.sh
	bash rename.sh
	bash scan.sh
	bash gen_html.sh
