# First stage: Build the binary
FROM golang:latest AS build

# ARG GA_CLIENT_SECRET
# ARG PORT
# ARG GA_CLIENT_ID

# # Set an environment variable with the value of the build argument
# ENV GA_CLIENT_ID=${GA_CLIENT_ID}
# ENV GA_CLIENT_SECRET=${GA_CLIENT_SECRET}
# ENV PORT=${PORT}
# RUN echo ${PORT}
# Set the working directory in the container
WORKDIR /app

# RUN go install github.com/melkeydev/go-blueprint@latest

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
# RUN git config --global user.email "go-blueprint@fake.com"
# RUN git config --global user.name "go-blueprint"


# Copy the source code to the container
COPY . .
# RUN rm -rf ./.git
RUN echo "APP_STAGE= STAGING" > app.env
# Build the binary
RUN CGO_ENABLED=0 go build -o main ./cmd/api

# Second stage: Copy the binary to a minimal image
FROM alpine:latest
RUN apk add --no-cache tzdata
# # Set the working directory
WORKDIR /
# # Copy the binary from the build stage
COPY --from=build /app/main .

COPY --from=build /app/app.env .
# Make the binary executable
RUN chmod +x main

# Expose the port that the application will run on
EXPOSE 8080

# Run the binary
ENTRYPOINT [ "./main" ]
