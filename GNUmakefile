# Variable for filename for store running procees id
PID = /tmp/gorm_app.pid
# Variable for the binary
BINARY=$(shell basename "$(PWD)")
# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/gorm_app-stderr.txt
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## start: Start in development mode. Auto-starts when code changes.
start:
	bash -c "trap 'make stop' EXIT; $(MAKE) compile start-server watch run='make compile start-server'"

# Start task run the binary app and writes it's process id to PID.
start-server: stop
	@echo " > Starting server..."
	@echo "  >  $(BINARY) binary is available at `pwd`/bin/"
	@-`pwd`/bin/$(BINARY) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"
	@echo " > You can check the proces at localhost:9000"

# Stop task will kill process by ID stored in PID 
stop:
	@echo " > Stop server"
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)  

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
watch:
	@echo " > You can check any errors on "
	`go env GOPATH`/bin/yolo -i . -e vendor -e bin  -c "$(run)" -a localhost:9000

go-compile: go-clean go-get go-build

go-build:
	@echo " > Building binary..."
	go build -o bin/$(BINARY)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	go get $(get)
	
## clean: Clean build files. Runs `go clean` internally.
clean:
	@(MAKEFILE) go-clean

go-clean:
	@echo "  >  Cleaning build cache"
	go clean


# .PHONY is used for reserving tasks words
.PHONY: start stop


