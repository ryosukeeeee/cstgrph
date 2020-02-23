PROGRAM=cost_explorer
SLACK_TOKEN=
SLACK_CHANNEL_ID=

.PHONY: debug
debug:
	@go build; \
	export SLACK_TOKEN=$(SLACK_TOKEN); \
	export SLACK_CHANNEL_ID=$(SLACK_CHANNEL_ID); \
	./$(PROGRAM);

.PHONY: deploy
deploy:
	@GOOS=linux go build -o bin/main; \
	export SLACK_TOKEN=$(SLACK_TOKEN); \
	export SLACK_CHANNEL_ID=$(SLACK_CHANNEL_ID); \
	sls deploy --token $(SLACK_TOKEN) --channel $(SLACK_CHANNEL_ID)

.PHONY: invoke
invoke:
	sls invoke -f main

.PHONY: clean
clean:
	rm $(PROGRAM)
