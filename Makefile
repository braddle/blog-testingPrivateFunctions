GOFILES= $$(go list -f '{{join .GoFiles " "}}')

test:
	go test -timeout=5s -cover -race ./number/ -v
