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
    triggers {
        cron "H 9 * * 3"
    }
    stages {
        stage("Empty_Folder") {
            steps {
                dir('/data/automation/checkouts'){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage('Checkout_Bowerbird'){
            steps{
                dir('/data/automation/checkouts/bowerbird'){
                    git url: 'https://github.com/farghul/bowerbird.git' , branch: 'main'
                }
            }
        }
        stage('Build_Bowerbird') {
            steps {
                dir('/data/automation/checkouts/bowerbird'){
                    script {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/bowerbird"
                    }
                }
            }
        }
        stage("Checkout_DAC") {
            steps{
                dir('/data/automation/checkouts/dac'){
                    git credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git', branch: 'main'
                }
            }
        }
        stage('Run_Bowerbird') {
            steps {
                dir('/data/automation/checkouts/dac'){
                    script {
                        sh './scripts/plugin/bowerbird.sh'
                    }
                }
            }
        }
    }
}