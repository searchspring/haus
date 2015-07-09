VERSION=0.1.5

buildmake:
	rm -rf Haus
	mkdir -p Haus/bin/
	rm -rf pkg1
	mkdir pkg1
	cp ../../../../haus ./Haus/bin/
	pkgbuild --root Haus \
    		--identifier org.haus.command \
    		--version $(VERSION) \
    		--ownership recommended \
    		pkg1/haus.pkg
	productbuild --distribution ./distrobution.xml \
              --resources ./resources \
              --package-path ./pkg1 \
              --version $(VERSION) \
              ../Haus-v$(VERSION)-amd64-osx10.8.pkg
