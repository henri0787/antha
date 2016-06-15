ASL=antha/AnthaStandardLibrary/Packages

all: gen_comp

gen_comp:
	go run cmd/antha/antha.go -outdir=antha/component/lib antha/component/an
	gofmt -w -s antha/component/lib

test:
	go test -v `go list ./... | grep -v internal | grep -v bvendor`

gen_pb:
	go generate github.com/antha-lang/antha/driver
	find driver/pb -name '*.pb.go' | xargs perl -p -i -e 's|proto "([^"]*)"|proto "github.com/antha-lang/antha/bvendor/\1"|'
	find driver/pb -name '*.pb.go' | xargs perl -p -i -e 's|context "([^"]*)"|context "github.com/antha-lang/antha/bvendor/\1"|'
	find driver/pb -name '*.pb.go' | xargs perl -p -i -e 's|grpc "([^"]*)"|grpc "github.com/antha-lang/antha/bvendor/\1"|'

fmt_json:
	for i in `find antha/examples -name '*.json' -o -name '*.yml'`; do \
	  python -mjson.tool "$$i" > "$$i.bak" && mv "$$i.bak" "$$i"; \
	done

compile:
	go install github.com/antha-lang/antha/cmd/...

test_workflows: compile
	for d in `find antha/examples -type d -o -name '*.yml'`; do \
	  if [[ -f "$$d/workflow.json" && -f "$$d/parameters.yml" ]]; then \
	    (cd "$$d" && antharun --workflow workflow.json --parameters parameters.yml $(ANTHA_ARGS) > /dev/null); \
	    if [[ $$? == 0 ]]; then \
	      echo "OK $$d"; \
	    else \
	      echo "FAIL $$d"; \
	    fi; \
	  fi; \
	done

assets: $(ASL)/asset/asset.go

$(ASL)/asset/asset.go: $(GOPATH)/bin/go-bindata-assetfs $(ASL)/asset_files/rebase/type2.txt
	$(GOPATH)/bin/go-bindata-assetfs -pkg=asset $(ASL)/asset_files && mv bindata_assetfs.go $@
	gofmt -s -w $@
	perl -p -i -e 's|github.com/elazarl/go-bindata-assetfs|github.com/antha-lang/antha/internal/github.com/elazarl/go-bindata-assetfs|' $@

$(ASL)/asset_files/rebase/type2.txt: ALWAYS
	mkdir -p `dirname $@`
	curl -o $@ ftp://ftp.neb.com/pub/rebase/type2.txt

$(GOPATH)/bin/2goarray:
	go get -u github.com/cratonica/2goarray

$(GOPATH)/bin/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

$(GOPATH)/bin/go-bindata-assetfs: $(GOPATH)/bin/go-bindata
	go get -u -f github.com/elazarl/go-bindata-assetfs/...
	touch $@

.PHONY: all gen_comp fmt_json test test_workflows compile assets ALWAYS
