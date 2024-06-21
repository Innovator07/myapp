# MyApp

This repository contains three applications: `go-app`, `nextjs-app`, and `wordpress-site`. This README provides instructions on how to set up and run these applications locally, as well as how to push changes to GitHub.

## Prerequisites

Ensure you have the following installed on your local machine:

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Node.js](https://nodejs.org/) (for Next.js)
- [Git](https://git-scm.com/)
- [Go](https://golang.org/)

## Repository Structure 

```plaintext
myapp/
├── go-app/
│   ├── Dockerfile
│   ├── main.go
│   └── go.mod
├── nextjs-app/
│   ├── Dockerfile
│   ├── package.json
│   ├── public/
│   ├── pages/
│   └── ...
├── wordpress-site/
│   ├── Dockerfile
│   └── ...
├── docker-compose.yml
└── README.md

## Steps to Setup and Run Applications on Local Machine:

1. Clone the Repository:
git clone https://github.com/innovator07/MyApp.git
cd MyApp

2. Build and Run with Docker Compose:
docker-compose up --build

This command will:
- Build Docker images for go-app, nextjs-app, and wordpress-site.
- Start all services defined in the docker-compose.yml file.

3. Individually Build and Run Applications:

1. Go App
Navigate to the go-app directory and build the Docker image:
cd go-app
docker build -t go-app:latest .
docker run -d -p 5002:5002 go-app:latest

2. Next.js App
Navigate to the nextjs-app directory and build the Docker image:
cd nextjs-app
docker build -t nextjs-app:latest .
docker run -d -p 3000:3000 nextjs-app:latest

3. WordPress Site
Navigate to the wordpress-site directory and build the Docker image:
cd wordpress-site
docker build -t wordpress-site:latest .
docker run -d -p 80:80 wordpress-site:latest

4. Pushing Changes to GitHub:

1. Configure Git
Set your Git configuration (only if not set previously):
git config --global user.email "you@xyz.com"
git config --global user.name "Your Name"

2. Commit and Push Changes
Navigate to the root directory of the repository, add your changes, commit, and push to GitHub:
git add .
git commit -m "Your commit message"
git push origin main

Or, if necessary:
git status
git push origin main --force

5. GitHub Actions for CI/CD:
The repository includes a GitHub Actions workflow (.github/workflows/deploy.yml) to automate the build and deployment process. Ensure you have set up the following secrets in your GitHub repository settings:
DOCKER_HUB_USERNAME
DOCKER_HUB_ACCESS_TOKEN

These secrets should be saved in the repository settings under the Secrets section.



This README includes instructions on setting up and running the applications locally, pushing changes to GitHub, and the GitHub Actions CI/CD workflow configuration. Make sure to adjust any paths, tags, or specific details to match your actual project setup.
