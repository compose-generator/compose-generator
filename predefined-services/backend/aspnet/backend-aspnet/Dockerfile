# Builder
FROM mcr.microsoft.com/dotnet/sdk:${{DOTNET_VERSION}} AS builder
WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . .
RUN dotnet publish -c Release -o out


# Minimalistic image
FROM mcr.microsoft.com/dotnet/aspnet:${{ASPNET_VERSION}}
WORKDIR /app
COPY --from=builder /app/out .
EXPOSE 80
ENTRYPOINT ["dotnet", "backend.dll"]