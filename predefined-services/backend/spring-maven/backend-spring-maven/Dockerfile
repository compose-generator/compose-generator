# Maven build instructions
FROM maven:3.8-jdk-11 AS builder
WORKDIR /server
COPY pom.xml /server/pom.xml
RUN mvn dependency:go-offline -q

COPY src /server/src
RUN mvn install -q
RUN mkdir -p target/depency
WORKDIR /server/target/dependency
RUN jar -xf ../*.jar

# JRE run instructions
FROM openjdk:11-jre-slim
EXPOSE 8080
ARG DEPENDENCY=/server/target/dependency
COPY --from=builder ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=builder ${DEPENDENCY}/META-INF /app/META-INF
COPY --from=builder ${DEPENDENCY}/BOOT-INF/classes /app
#? if var.SPRING_MAVEN_LANGUAGE != "kotlin" {
#ENTRYPOINT ["java","-cp","app:app/lib/*","${{SPRING_MAVEN_PACKAGE_NAME}}.${{PROJECT_NAME}}"]
#? }
#? if var.SPRING_MAVEN_LANGUAGE == "kotlin" {
#ENTRYPOINT ["java","-cp","app:app/lib/*","${{SPRING_MAVEN_PACKAGE_NAME}}.${{PROJECT_NAME}}Kt"]
#? }