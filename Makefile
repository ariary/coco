before.build:
	go mod tidy && go mod download

build.agent:
	@echo "build in ${PWD}";go build cmd/agent/coco-agent.go