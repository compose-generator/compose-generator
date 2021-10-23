# Builder
FROM gradle:7.2-jdk17 AS builder
WORKDIR /server
COPY *.gradle *.kts ./
RUN gradle build -q || return 0

COPY ./src ./src
RUN gradle build -w
RUN java -Djarmode=layertools -jar build/libs/demo-0.0.1-SNAPSHOT.jar extract


# Minimalistic image
FROM openjdk:18-alpine
EXPOSE 8080

COPY --from=builder /server/dependencies/ ./
COPY --from=builder /server/snapshot-dependencies/ ./
COPY --from=builder /server/spring-boot-loader/ ./
COPY --from=builder /server/application/ ./
ENTRYPOINT ["java", "org.springframework.boot.loader.JarLauncher"]
