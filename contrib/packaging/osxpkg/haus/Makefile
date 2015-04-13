VERSION=0.1.3

buildmake:
	mkdir -p Haus/bin/
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
