.PHONY: build build2 demo upload clean

build:
	tsc
# r.js -o build.ts

build2:
	vulcanize --inline-scripts --inline-css --strip-comments demo/index.html > demo/index.v.html

demo:
	tsc
	dev_appserver.py --host=0.0.0.0 demo.yaml

upload: build
	appcfg.py update demo.yaml

clean:
	rm demo/*.js
	rm demo/*.map

# build :
# 	polymer serve&
# 	babel src -w -s -o bundle.js

# lint :
# 	eslint src
# 	polymer lint --input hello-world.html

# server -views=default -root=.
# dev_appserver.py --host=0.0.0.0 --port=8081 default.yaml

# screen -d -m browser-sync start --files "default/*.html, default/*.css, default/*.js" --proxy="localhost:8080"
# screen -ls | grep Detached | cut -d. -f1 | awk "{print $1}" | xargs kill
# ~/appengine/dev_appserver.py default.yaml
# vulcanize -o build/index.html default/index.html --inline --strip
# ~/appengine/dev_appserver.py build.yaml
# ~/appengine/appcfg.py --oauth2 update build.yaml
# ~/appengine/appcfg.py --oauth2 update_dispatch dispatch.yaml build.yaml
# ~/appengine/appcfg.py --oauth2 update_indexes index.yaml build.yaml
# ~/appengine/appcfg.py --oauth2 vacuum_indexes index.yaml build.yaml
# ~/appengine/appcfg.py --oauth2 rollback build.yaml build.yaml

# r.js -o baseUrl=./demo paths.requireLib=../bower_components/requirejs/require name=test include=requireLib out=./demo/built.js optimize=none