VERSION=0.1.7

buildmake:
	rm -rf Haus
	mkdir -p Haus/usr/local/bin
	rm -rf pkg1
	mkdir pkg1
	cp ../../../../haus ./Haus/usr/local/bin/
	pkgbuild --root Haus \
    		--identifier org.haus.command \
    		--version $(VERSION) \
    		--ownership recommended \
    		pkg1/haus.pkg
	productbuild --distribution ./distrobution.xml \
              --resources ./resources \
              --package-path ./pkg1 \
              --version $(VERSION) \
              ../Haus-v$(VERSION)-amd64-osx10.11.pkg
