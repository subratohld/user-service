def gv

pipeline{
    agent any

    environment {
        commitHash = ''
    }

    stages{
        stage("Init") {
            steps {
                script {
                    gv = load 'pipeline.groovy'
                }
            }
        }

        stage("Build"){
            steps{
                script {
                    env.commitHash = gv.getCommitHash()
                    docker.build "subratohld/user-service:${env.commitHash}"
                }
            }
        }

        stage("Test") {
            steps{
                echo "Testing..."
            }
        }

        stage("Deploy"){
            steps{
                echo "Deploying..."
            }
        }
    }
}