# Backend Deployment TODO (Fly.io + SQLite Volume)

## [x] 1. Prep: go mod tidy
## [ ] 2. Create persistent volume for SQLite (Trial expired - skipped for basic deploy)
## [ ] 3. Set secure secrets (JWT_SECRET, DB_PATH) (Skipped - trial limits)
## [ ] 4. Deploy (Blocked - trial expired)
## [ ] 5. Verify status, open app, check logs
## [x] 6. Docker build/run (Executing)

## [ ] Docker deploy (Render/DH/EC2)

**Status:** Docker setup complete.
**Run:** docker build -t learning-api .
docker run -p 8080:8080 -v $(pwd)/data:/data learning-api   # with vol

**Fly:** Add CC then deploy.

