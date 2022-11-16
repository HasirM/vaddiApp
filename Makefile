project_name = vaddiapp/vaddi
c = 

# Clean packages
run:
	go mod tidy
	go run cmd/vaddi/main.go

# Generate go.mod & go.sum files
git:
	git add .
	git commit -m "$(c)"
	git push
