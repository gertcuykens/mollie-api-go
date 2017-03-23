.PHONY: test dispatch upload html rollback index

test:
	dev_appserver.py --host=0.0.0.0 demohtml/demohtml.yaml demomodule/demomodule.yaml

dispatch:
	appcfg.py update_dispatch ./

upload:
	appcfg.py update demomodule/demomodule.yaml
	appcfg.py update demohtml/demohtml.yaml

html:
	tsc -p demohtml
	appcfg.py update demohtml/demohtml.yaml

# rollback:
# 	appcfg.py rollback demomodule/demomodule.yaml
# 	appcfg.py rollback demohtml/demohtml.yaml

# index:
# 	appcfg.py update_indexes demomodule/index.yaml
# 	appcfg.py vacuum_indexes demomodule/index.yaml

# https://v1-dot-demo-dot-mollie-api-go.appspot.com
# https://v1-dot-default-dot-mollie-api-go.appspot.com
