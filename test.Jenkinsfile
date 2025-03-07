pipeline {
    agent any
    stages {
        stage("Build"){
            steps {
                bat "go mod tidy"
            }
        }
        stage("Test"){
            steps {
                bat "go test ."
            }
        }
        stage("Deploy"){
            steps {
                bat "./jenkins/scripts/deliver.bat"
            }
        }
        stage("Complete"){
            steps {
                echo "job complete"
            }
        }
    }
}