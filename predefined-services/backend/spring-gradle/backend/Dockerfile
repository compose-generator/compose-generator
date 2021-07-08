# Gradle build instructions
FROM gradle:7.1-jdk16 AS build
WORKDIR /server
COPY *.gradle *.kts ./
RUN gradle build -q || return 0

COPY src src
RUN gradle build -w
RUN java -Djarmode=layertools -jar build/libs/*.jar extract

# JRE run instructions
FROM openjdk:16-alpine AS prod
EXPOSE 8080
ARG DEPENDENCY=/server/target/dependency

COPY --from=build /server/dependencies/ ./
COPY --from=build /server/snapshot-dependencies/ ./
COPY --from=build /server/spring-boot-loader/ ./
COPY --from=build /server/application/ ./
ENTRYPOINT ["java", "org.springframework.boot.loader.JarLauncher"]
