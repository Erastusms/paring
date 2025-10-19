# Gunakan JDK 21 + Maven
FROM maven:3.9.6-eclipse-temurin-21 AS build

WORKDIR /app

# Copy file pom dan source
COPY pom.xml .
COPY src src

# Build aplikasi tanpa test
RUN mvn clean package -DskipTests

# Tahap kedua: buat image runtime ringan
FROM eclipse-temurin:21-jre

WORKDIR /app
COPY --from=build /app/target/*.jar app.jar

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
