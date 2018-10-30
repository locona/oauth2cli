.DEFAULT_GOAL := run


.PHONY: run
run:
	@go install
	@oauth2cli token -u 'http://localhost:14444/oauth2/token' -i 'subjects:hydra:clients:oathkeeper-client' -s 'secret'
