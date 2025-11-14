import os
import sys

if os.environ.get('AWS_LAMBDA_RUNTIME_API'):
    # Running in AWS Lambda with the Lambda Runtime Interface Client (RIC)
    os.execv(sys.executable, [sys.executable, '-m', 'awslambdaric'] + sys.argv[1:])
else:
    # Running locally with the Lambda Runtime Interface Emulator (RIE)
    os.execv('/usr/local/bin/aws-lambda-rie', ['/usr/local/bin/aws-lambda-rie', sys.executable, '-m', 'awslambdaric'] + sys.argv[1:])
