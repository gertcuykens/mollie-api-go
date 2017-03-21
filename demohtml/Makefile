ALL := $(wildcard css/*.css)
CSS := $(filter-out %.min.css,$(ALL))
JS := $(wildcard scripts/*.js)

.PHONY: uglify build install clean demo

uglify: build $(CSS) $(JS)
	uglifyjs node_modules/alameda/alameda.js -o scripts/alameda.js -c

%.css:
	uglifycss $@ > $(patsubst %.css,%.min.css,$@)

%.js:
	uglifyjs $@ -o $@ -c

build:
	tsc --pretty

install:
	npm i typescript -g
	npm i uglifyjs -g
	npm i uglifycss -g
	npm i

clean:
	-rm scripts/*.js
	-rm scripts/*.map
	-rm css/*.min.css
