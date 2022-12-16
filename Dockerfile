FROM amazonlinux:2
RUN yum install -y sudo jq awscli shadow-utils htop lsof telnet bind-utils yum-utils tar wget && \
    yum install -y https://s3.ap-northeast-1.amazonaws.com/amazon-ssm-ap-northeast-1/latest/linux_amd64/amazon-ssm-agent.rpm && \
    adduser ssm-user && echo "ssm-user ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/ssm-agent-users && \
    mv /etc/amazon/ssm/amazon-ssm-agent.json.template /etc/amazon/ssm/amazon-ssm-agent.json && \
    mv /etc/amazon/ssm/seelog.xml.template /etc/amazon/ssm/seelog.xml
RUN wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz && \
    export GOPATH="/usr/local/go" && \
    export PATH=$PATH:$GOPATH/bin && \
    go version
COPY run.sh /run.sh
COPY main.go /main.go
CMD ["sh", "/run.sh"]