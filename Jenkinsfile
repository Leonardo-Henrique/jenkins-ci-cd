#!/usr/bin/env groovy

pipeline {

    agent {
        docker {
            image 'golang'
            args '-u root'
        }
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go mod download'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test'
            }
        }
    }
}