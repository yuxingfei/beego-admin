pipeline {
    agent any

    tools {
            go 'Go1.21.5'
    }
    stages {
        stage('[checkout]拉取代码中....') {
            steps {
                // 使用 checkout 步骤拉取代码
                checkout([$class: 'GitSCM', branches: [[name: '*/main']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'b4c9eb45-0344-409f-849a-2cdd091749ee', url: 'http://113.250.161.105:8180/geeker_yuyu/beego-admin.git']],destination: 'beego-admin'])
                echo '拉取代码成功'
            }
        }
        stage('[build]构建中....') {
            steps {
                // 构建 Golang 项目
                script {
                    // 使用 Golang 工具
                    def gopath = tool 'Go1.21.5'
                    env.GOPATH = "${WORKSPACE}/gopath"
                    env.PATH = "${gopath}/bin:${env.PATH}"

                    // 构建命令
                    sh 'go build -o beego-admin .'

                    echo '构建完成'
                }
            }
        }
       stage('[test]发送软件到测试服务器....') {
           steps {
               script {
                   // 测试阶段的构建命令和其他操作

                   // 使用 SSH 将生成的可执行文件和 Dockerfile 发送到测试服务器
                   sshPublisher(
                       failOnError: true,
                       publishers: [
                           sshPublisherDesc(
                               configName: 'b4c9eb45-0344-409f-849a-2cdd091749ee', // 替换为你的 SSH 凭据的 ID
                               transfers: [
                                   sshTransfer(
                                       sourceFiles: 'beego-admin ./Dockerfile', // 本地构建好的二进制文件和 Dockerfile 路径
                                       removePrefix: '',
                                       remoteDirectory: '/opt/deployment' // 远程测试服务器的目录
                                   )
                               ]
                           )
                       ]
                   )

                   echo '发送软件到测试服务器'
               }
           }
       }
        stage('[deploy]部署中....') {
            steps {
                script {
                    def remote = [:]
                    remote.name = 'gitlab-jenkins-beego-admin' // 服务器名字，提前配置好的
                    remote.host = '113.250.161.105' // 远程ip，待部署的服务器ip
                    remote.port = 22
                    remote.allowAnyHosts = true

                    withCredentials([usernamePassword(credentialsId: 'db4cb5f7-08fb-465a-baee-3a868f9d69e2', passwordVariable: 'Password', usernameVariable: 'Username')]) {
                        remote.user = "${Username}"
                        remote.password = "${Password}"
                    }

                    writeFile file: 'abc.sh', text: 'docker stop demo;docker container rm demo;docker rmi demo;chmod +x /opt/deployment/beego-admin;cd /opt/deployment;docker build -t demo .;docker run --name demo -d -p 7654:7654 demo'
                    sshScript remote: remote, script: 'abc.sh'
                }

            }
        }
    }

    post {
        success {
            // 构建成功时执行的操作
            echo 'Build and deployment successful!'
        }
        failure {
            // 构建失败时执行的操作
            echo 'Build or deployment failed!'
        }
    }
}
