
@echo build web
@node build/build.js

@echo commit to github

@cd ../vue-tally
@rmdir /s/q static

@xcopy /E/I/Y "../tally/dist" "."
@copy "../tally/manifest.json" "vue-tally"

@git pull
@git add .
@git commit -m "code build dispense to service"
@git push origin master 

@cd ../tally

@git pull
@echo build go
@set GOARCH=amd64
@set GOOS=linux
@go build -v -a -o tally

@copy "tally" "../go-tally/"

@cd ../go-tally

@echo commit to github
@git add .
@git commit -m "code build dispense to service"
@git push origin master 

@cd ../tally



