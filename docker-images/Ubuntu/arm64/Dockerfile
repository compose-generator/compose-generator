FROM ubuntu

RUN apt-get update && apt-get install -y ca-certificates curl gnupg wget lsb-release

RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt-get update

RUN apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

RUN wget https://github.com/compose-generator/ccom/releases/latest/download/ccom_arm64.deb && dpkg -i ccom_arm64.deb
RUN wget https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_arm64.deb && dpkg -i compose-generator_arm64.deb