def gv

pipeline{
    agent any

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