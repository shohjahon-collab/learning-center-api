# Render Deploy Fix & GitHub Push TODO

## [ ] 1. Edit go.mod: module app
## [ ] 2. Edit cmd/api/main.go: Update imports from "learning-center-api/internal" to "app/internal"
## [ ] 3. Edit Dockerfile: Change build to -o app, CMD ./app
## [ ] 4. go mod tidy
## [ ] 5. Test build: go build -o app ./cmd/api/main.go
## [ ] 6. git add . & commit "Fix for Render: module app"
## [ ] 7. Create GitHub repo learning-center-api
## [ ] 8. git remote add origin https://github.com/{YOUR_USERNAME}/learning-center-api.git
## [ ] 9. git push -u origin main
## [ ] 10. Re-deploy on Render
