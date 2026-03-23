FROM mcr.microsoft.com/dotnet/sdk:6.0.100 AS build
WORKDIR /src

COPY *.sln .
COPY Application/*.csproj Application/
COPY Domain/*.csproj Domain/
COPY Infrastructure/*.csproj Infrastructure/
COPY WebAPI/*.csproj WebAPI/

RUN dotnet restore

COPY . .

RUN dotnet publish "WebAPI/WebAPI.csproj" -c Debug -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS runtime
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "WebAPI.dll"]
