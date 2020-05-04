pipeline {
    agent any

    environment {
        AWS_ACCESS_KEY_ID     = "${env.AWS_ACCESS_KEY_ID}"
        AWS_SECRET_ACCESS_KEY = "${env.AWS_SECRET_ACCESS_KEY}"
        TF_IN_AUTOMATION      = '1'
        ENVIRONMENT           = "${env.ENVIRONMENT}"
    }

    stages {
        stage('Plan') {
            steps {
                script {
                    currentBuild.displayName = "${version}"
                }
                sh '/usr/local/bin/terraform init -input=false'
                sh '/usr/local/bin/terraform workspace select ${ENVIRONMENT}'
                sh "/usr/local/bin/terraform plan -input=false -out tfplan -var 'version=${version}' --var-file=vars/${ENVIRONMENT}.tfvars"
                sh '/usr/local/bin/terraform show -no-color tfplan > tfplan.txt'
            }
        }

        stage('Approval') {
            when {
                not {
                    equals expected: true, actual: params.autoApprove
                }
            }

            steps {
                script {
                    def plan = readFile 'tfplan.txt'
                    input message: "Do you want to apply the plan?",
                        parameters: [text(name: 'Plan', description: 'Please review the plan', defaultValue: plan)]
                }
            }
        }

        stage('Apply') {
            steps {
                sh "/usr/local/bin/terraform apply -input=false tfplan"
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'tfplan.txt'
        }
    }
}