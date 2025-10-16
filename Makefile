### test / build / dev

build:
	go build gen_html.go

go_test:
	go run ./gen_html.go ./results/osin-storage.trufflog ./results/osin-storage.trufflog.html highlighter.js


### usage

run:
	bash test_token.sh
	bash get.sh
	bash scan.sh
	bash fix_repo_urls.sh
	bash rm_empty.sh
	bash gen_html.sh
	bash echo_url.sh


rm:
	rm -rf results repos
	mkdir results repos 
