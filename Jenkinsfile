node('huhaiwen-jnlp') {
  stage('Prepare') {
    echo "1.初始化环境变量"
    checkout scm
    script {
      build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
      if (env.BRANCH_NAME != 'master') {
        build_tag = "${env.BRANCH_NAME}-${build_tag}"
      }
    }
  }
  stage('Test') {
    echo "2. Test"
  }
  stage('Build') {
    echo "3.构建docker镜像"
    sh "docker build -t huhaiwen/jenkins-demo:${build_tag} ."
  }
  stage('Push') {
    echo "推送docker镜像"
    withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
      sh "docker login -u ${dockerHubUser} -p ${dockerHubPassword}"
      sh "docker push huhaiwen/jenkins-demo:${build_tag}"
    }
  }
  stage('Deploy') {
    echo "5. Deploy Stage"
      if (env.BRANCH_NAME == 'master') {
        input "确认要部署线上环境吗？"
    }
    sh "sed -i 's/<BUILD_TAG>/${build_tag}/' jenkins.yaml"
    sh "sed -i 's/<BRANCH_NAME>/${env.BRANCH_NAME}/' jenkins.yaml"
    sh "kubectl apply -f k8s.yaml --record"
  }
}