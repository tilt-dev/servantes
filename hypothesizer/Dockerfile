FROM python:3.6

RUN true tilt-magic start_cloud_sync
ADD requirements.txt /app/requirements.txt
RUN cd /app && pip install -r requirements.txt
ADD . /app
