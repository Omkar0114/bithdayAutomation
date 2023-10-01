# Use the official AWS Lambda base image for Go

FROM public.ecr.aws/lambda/go:1

# Set the working directory
WORKDIR /var/task

# Copy your Go application binary and any other necessary files
COPY birthdays birthdaysList.csv ./ 

# Set the AWS Lambda function handler (replace with your handler)
# CMD ["sendBirthdayWish"]

# Optional: Set environment variables if needed
# ENV ENV_VAR_NAME=value
ENV TWILIO_ACCOUNT_SID=""
ENV TWILIO_AUTH_TOKEN=""


# Expose port if needed (AWS Lambda doesn't use ports)
# EXPOSE 8080

# Make sure your binary is executable
RUN chmod +x birthdays

# Ensure that any shared libraries are available
# You may need to add this step if your application depends on shared libraries

# Set the Lambda runtime for Go
ENV AWS_LAMBDA_GO_RUNTIME=go1.x

# Start the Lambda function
CMD ["birthdays"]
