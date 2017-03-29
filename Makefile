#ALL := $(wildcard demohtml/css/*.css)
#CSS := $(filter-out %.min.css,$(ALL))
JS := $(wildcard demohtml/scripts/*.js)

.PHONY: default uglify build install clean test dispatch module html rollback index

default: build
	dev_appserver.py --host=0.0.0.0 demohtml/demohtml.yaml demomodule/demomodule.yaml

uglify: build $(JS) #$(CSS)
	uglifyjs node_modules/alameda/alameda.js -o scripts/alameda.js -c
	uglifyjs ../demoservice/service.js -o scripts/service.js -c

%.js:
	uglifyjs $@ -o $@ -c

#%.css:
#	uglifycss $@ > $(patsubst %.css,%.min.css,$@)

build:
	tsc --pretty -p demohtml/scripts
	tsc --pretty -p demohtml/workers
# cp demohtml/workers/service.js demohtml/service.js

dispatch:
	appcfg.py update_dispatch ./

module:
	appcfg.py update demomodule/demomodule.yaml

html: build
	appcfg.py update demohtml/demohtml.yaml

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

# rollback:
# 	appcfg.py rollback demomodule/demomodule.yaml
# 	appcfg.py rollback demohtml/demohtml.yaml

# index:
# 	appcfg.py update_indexes demomodule/index.yaml
# 	appcfg.py vacuum_indexes demomodule/index.yaml

# https://v1-dot-demo-dot-mollie-api-go.appspot.com
# https://v1-dot-default-dot-mollie-api-go.appspot.com