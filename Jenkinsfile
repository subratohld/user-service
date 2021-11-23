def gv = load 'pipeline.groovy'

pipeline{
    agent any

    stages{
        stage("Build"){
            steps{
                sh "chmod +x ./scripts/*.sh"
                
                script {
                    def gitCommit = gv.getCommitHash()
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