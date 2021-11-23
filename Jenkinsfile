pipeline{
    agent any

    // tools {
    //     'org.jenkinsci.plugins.docker.commons.tools.DockerTool' 'docker'
    // }

    // environment {
    //     DOCKER_CERT_PATH = credentials('id-for-a-docker-cred')
    // }

    stages{
        stage("Build"){
            steps{
                sh "chmod +x ./scripts/*.sh"
                
                script {
                    def gitCommit = sh 'git rev-parse --short origin/main'
                    echo "${gitCommit}"
                    // docker.build "subratohld/user-service:1.0"
                }

                // ${env.BUILD_TAG}
                // sh "make build"
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
    // post{
    //     always{
    //         echo "========always========"
    //     }
    //     success{
    //         echo "========pipeline executed successfully ========"
    //     }
    //     failure{
    //         echo "========pipeline execution failed========"
    //     }
    // }
}