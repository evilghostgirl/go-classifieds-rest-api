FROM golang AS build
LABEL stage=build
WORKDIR /dist
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a
CMD ["go", "run", "main.go"]
# FROM alpine AS final
# EXPOSE 8080
# ARG ENV
# WORKDIR /
# COPY --from=build /dist/classifieds-rest-api .
# CMD ["./classifieds-rest-api" ]