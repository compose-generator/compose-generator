FROM mcr.microsoft.com/dotnet/sdk:${{DOTNET_VERSION}} AS build-env
WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . .
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/aspnet:${{ASPNET_VERSION}}
WORKDIR /app
COPY --from=build-env /app/out .
EXPOSE 80
ENTRYPOINT ["dotnet", "backend.dll"]