# chromedp-AWSlambda
The steps I'm running it:    
1. `docker build -t lambda-crawler .`

2. `docker tag lambda-crawler:latest 123456789012.dkr.ecr.us-east-1.amazonaws.com/lambda-crawler:latest`

3. `docker push 123456789012.dkr.ecr.us-east-1.amazonaws.com/lambda-crawler:latest`    


4. `aws lambda update-function-code --region us-east-1 --function-name lambda-crawler --image-uri 123456789012.dkr.ecr.us-east-1.amazonaws.com/lambda-crawler:latest`  

5. `aws lambda invoke --function-name lambda-crawler --cli-binary-format raw-in-base64-out  --payload '{"name": "success"}' output.txt`  
