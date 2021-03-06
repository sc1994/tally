
@echo build web

@node build/build.js

@echo commit to github

@cd ../vue-tally
@rmdir /s/q static

@xcopy /E/I/Y "../tally/dist" "."
@xcopy /E/I/Y "../tally/static" "."

@echo commit to github...

@git pull origin master
@git add .
@git commit -m "code build dispense to service"
@git push -u origin master -f

@cd ../tally

@git pull
@echo build go
@set GOARCH=amd64
@set GOOS=linux
@go build -v -a -o tally

@copy "tally" "../go-tally/"

@cd ../go-tally

@echo commit to github...

@git pull origin master
@git add .
@git commit -m "code build dispense to service"
@git push -u origin master -f

@cd ../tally

@echo END
@pause


