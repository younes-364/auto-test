pipeline {
    agent any
    stages {
        // stage('Install Dependencies') {
        //     steps {
	    //         sh "go install github.com/gruntwork-io/terratest/modules/terraform@latest"
	    //         sh "go install github.com/stretchr/testify/assert@latest"
        //         sh "go install github.com/gruntwork-io/terratest/modules/test-structure"
        //     }
        // }
        stage('git'){
            steps {
                git branch: 'main', credentialsId: '9aac788b-90c4-4013-b24e-d305139e5b42', url: 'https://github.com/younes-364/auto-test'
            }
        }
        stage('Run Tests') {
            steps {
                sh 'go test -v ./'
            }
        }
        stage('Upload Results') {
            steps {
                sh 'curl -X POST -F "file=@test-results.xml" http://reporting-server/upload'
            }
        }
        // stage('Configure Terraform Cloud') {
        //     steps {
        //         env.ATLAS_TOKEN = 'eM6P0JyKaZzJaA.atlasv1.U0oMslxMeZtNlTMM5owqnPg0MomT8GIHa2GhyrTNl0zgE7vzksFfzCqHF2Wjzyuaflw'
        //     }
        // }
        // stage('Provision infrastructure') {
        //     steps {
        //         terraform(
        //             installation: 'Terraform Cloud',
        //              env.ATLAS_TOKEN = 'eM6P0JyKaZzJaA.atlasv1.U0oMslxMeZtNlTMM5owqnPg0MomT8GIHa2GhyrTNl0zgE7vzksFfzCqHF2Wjzyuaflw',
        //             backendType: 'workspaces',
        //             backendConfig: [
        //                 name: 'my-app-prod'
        //             ]
        //             // vars: [
        //             //     myVar: 'myValue'
        //             // ]
        //         )
        //     }
        // }


    }
}
