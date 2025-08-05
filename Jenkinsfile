pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    stages {
        stage('Clean WS') {
            steps {
                cleanWs()
            }
        }
        stage("Checkout Bowerbird") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[url: 'https://github.com/farghul/bowerbird.git']]
                )
            }
        }
        stage("Build Bowerbird") {
            steps {
                script {
                    sh "/data/apps/go/bin/go build -o /data/automation/bin/bowerbird"
                }
            }
        }
        stage("Checkout DAC") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git']]
                )
            }
        }
        stage('Run Bowerbird') {
            steps {
                script {
                    sh './scripts/plugin/bowerbird.sh'
                }
            }
        }
    }
}