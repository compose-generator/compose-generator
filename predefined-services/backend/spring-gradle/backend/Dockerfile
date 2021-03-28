FROM gradle:6.8.3-jdk15 AS build
WORKDIR /workdir/server
COPY build.gradle.kts build.gradle.kts
COPY settings.gradle.kts settings.gradle.kts
RUN gradle build || return 0

COPY src src
RUN gradle build
RUN java -Djarmode=layertools -jar build/libs/*.jar extract


FROM adoptopenjdk/openjdk15:jre-15.0.2_7-alpine AS prod
EXPOSE 8080
COPY --from=build /workdir/server/dependencies/ ./
COPY --from=build /workdir/server/snapshot-dependencies/ ./
COPY --from=build /workdir/server/spring-boot-loader/ ./
COPY --from=build /workdir/server/application/ ./
ENTRYPOINT ["java", "org.springframework.boot.loader.JarLauncher"]
