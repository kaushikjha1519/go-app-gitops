pipeline {
    agent any

    environment {
        DOCKERHUB_USER   = 'kaushikkjha'
        IMAGE_NAME       = 'go-app'
        GITOPS_REPO_NAME = 'gomanifest'
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o server .'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build & Push') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                    echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin

                    docker build -t $DOCKERHUB_USER/$IMAGE_NAME:$BUILD_NUMBER .
                    docker tag $DOCKERHUB_USER/$IMAGE_NAME:$BUILD_NUMBER $DOCKERHUB_USER/$IMAGE_NAME:latest

                    docker push $DOCKERHUB_USER/$IMAGE_NAME:$BUILD_NUMBER
                    docker push $DOCKERHUB_USER/$IMAGE_NAME:latest
                    '''
                }
            }
        }

        stage('Update GitOps Repo') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'github-creds',
                    usernameVariable: 'GIT_USER',
                    passwordVariable: 'GIT_PASS'
                )]) {
                    sh '''
                    git clone https://github.com/kaushikjha1519/$GITOPS_REPO_NAME.git
                    cd $GITOPS_REPO_NAME

                    git config user.email "jenkins@ci.local"
                    git config user.name "Jenkins"

                    git remote set-url origin https://$GIT_USER:$GIT_PASS@github.com/kaushikjha1519/$GITOPS_REPO_NAME.git

                    sed -i '' 's|image: '"$DOCKERHUB_USER/$IMAGE_NAME"':.*|image: '"$DOCKERHUB_USER/$IMAGE_NAME:$BUILD_NUMBER"'|' deployment.yaml

                    git add deployment.yaml
                    git commit -m "ci: update image tag to build-$BUILD_NUMBER" || true
                    git push
                    '''
                }
            }
        }
    }

    post {
        always {
            sh 'rm -rf $GITOPS_REPO_NAME'
        }
        success {
            echo "✅ Pipeline succeeded. ArgoCD will sync shortly."
        }
        failure {
            echo "❌ Pipeline failed. Check logs above."
        }
    }
}