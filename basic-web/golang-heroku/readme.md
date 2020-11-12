# Dpeloying Golang in heroku 

1.  Create a Procfile and project name 
```
echo "web: golang-heroku" > Procfile
```
2.  Then run 
```
go mod init <project path>
```
3. Just deploy on heroku