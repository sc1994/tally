
@echo build web
@node build/build.js

@echo 提交到 github-->master

@xcopy /E/I/Y "dist" "../vue-tally/dist"

@cd ../vue-tally

git add .
git commit -m "code build dispense to service"
git push origin master 

@cd ../tally

@echo build go
set GOARCH=amd64
set GOOS=linux
go build -v






