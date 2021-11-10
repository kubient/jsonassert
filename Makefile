HELP_FUN = \
		%help; \
		while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^(\w+)\s*:.*\#\#(?:@(\w+))?\s(.*)$$/ }; \
		print "usage: make [target]\n\n"; \
	for (keys %help) { \
		print "$$_:\n"; $$sep = " " x (20 - length $$_->[0]); \
		print "  $$_->[0]$$sep$$_->[1]\n" for @{$$help{$$_}}; \
		print "\n"; }

ifeq (release,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "release"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

confirm=sh -c '\
  read -p "type y or n: " -n 1 -r ; \
  if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
    echo "\nok"; \
  else \
    exit; \
  fi' confirm

help:           ##@miscellaneous Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

release: ## Make release and push to Github with tag
	echo "version is $(RUN_ARGS)?"
	${confirm}
	echo "did you add the version $(RUN_ARGS) to CHANGELOG.md?"
	${confirm}
	git commit -m "bumped to version $(RUN_ARGS)" ./CHANGELOG.md
	git tag -a $(RUN_ARGS) -m "version $(RUN_ARGS)"
	git push --atomic origin master  $(RUN_ARGS)

test: ## Run tests
	go clean -testcache && go test -v ./...