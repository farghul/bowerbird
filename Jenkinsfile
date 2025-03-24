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
        stage("Pull Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/bowerbird") {
                        sh "git pull"
                    }
                }
            }
        }
        stage("Build Bowerbird") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/bowerbird") {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/bowerbird ."
                    }
                }
            }
        }
        stage("Run Bowerbird") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/automation/scripts/bowerbird.sh"
                        }
                    }
                }
            }
        }
    }
}