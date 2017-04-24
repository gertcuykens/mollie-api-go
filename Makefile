PROJECT := mollie-api-go
ALL := $(wildcard demohtml/css/*.css)
CSS := $(filter-out %.min.css,$(ALL))
JS := $(wildcard demohtml/scripts/*.js)

.PHONY: server uglify build install dispatch acme module html test install clean index

server: build
	dev_appserver.py --host=0.0.0.0 demohtml/default.yaml demomodule/demo.yaml

uglify: build $(JS) $(CSS)

%.js:
	uglifyjs $@ -o $@ -c

%.css:
	uglifycss $@ > $(patsubst %.css,%.min.css,$@)

build:
	tsc --pretty -p demohtml/scripts
	tsc --pretty -p demohtml/workers

dispatch:
	gcloud app deploy dispatch.yaml --project $(PROJECT)

acme:
	gcloud app deploy acme/acme.yaml --project $(PROJECT)

module:
	gcloud app deploy demomodule/demo.yaml --project $(PROJECT)

html: build
	gcloud app deploy demohtml/default.yaml --project $(PROJECT)

test:
	go test -v ./demomodule
	go test -v ./demomail
# go test -run=Time/12:[0-9] -v

install:
	npm i typescript -g
	npm i uglifyjs -g
	npm i uglifycss -g
	go get -u github.com/gertcuykens/mollie-api-go
	go get -u github.com/gertcuykens/httx
	go get -u google.golang.org/appengine

clean:
	-rm demohtml/scripts/*.js
	-rm demohtml/scripts/*.map
	-rm demohtml/workers/*.js
	-rm demohtml/workers/*.map
	-rm demohtml/css/*.min.css

# index:
# 	gcloud datastore create-indexes demomodule/index.yaml
# 	gcloud datastore cleanup-indexes demomodule/index.yaml

# https://v1-dot-demo-dot-mollie-api-go.appspot.com
# https://v1-dot-default-dot-mollie-api-go.appspot.com
