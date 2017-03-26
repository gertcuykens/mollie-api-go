#ALL := $(wildcard css/*.css)
#CSS := $(filter-out %.min.css,$(ALL))
JS := $(wildcard scripts/*.js)

.PHONY: uglify build install clean demo

uglify: build $(JS) #$(CSS)
	uglifyjs node_modules/alameda/alameda.js -o scripts/alameda.js -c
	uglifyjs ../demoservice/service.js -o scripts/service.js -c

%.js:
	uglifyjs $@ -o $@ -c

#%.css:
#	uglifycss $@ > $(patsubst %.css,%.min.css,$@)

build:
	tsc --pretty

install:
	npm i typescript -g
	npm i uglifyjs -g
	npm i uglifycss -g

clean:
	-rm scripts/*.js
	-rm scripts/*.map
	-rm css/*.min.css
