def gv
def commitHash

pipeline{
    agent any

    // environment {
    //     commitHash = ''
    // }

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
                echo "Building..."
            }
        }
        stage("Test") {
            steps{
                echo "Testing..."
            }
        }
        stage("Build Image"){
            steps{
                script {
                    commitHash = gv.getCommitHash()
                    docker.build "subratohld/user-service:${commitHash}"
                }
            }
        }
        stage("Deploy"){
            steps{
                echo "Deploying..."
            }
        }
    }
}